package main

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"github.com/samber/lo"
	"io"
	"regexp"
	"strconv"
	"strings"
	"testing"
	"time"
)

func TestLo(t *testing.T) {
	var s = []string{"aa", "aa", "a", "b", "bb", "c", "c", "cc"}
	//var s []string
	b := []string(lo.Uniq(s))

	fmt.Println(b)

}

func splitHostPort(hostPort string) (host, port string) {
	host = hostPort

	colon := strings.LastIndexByte(host, ':')
	host, port = host[:colon], host[colon+1:]

	if strings.HasPrefix(host, "[") && strings.HasSuffix(host, "]") {
		host = host[1 : len(host)-1]
	}

	return
}

type PgSchema []string

func removeKthElement(schema PgSchema, k int) PgSchema {
	schema[k] = schema[len(schema)-1]
	return schema[:len(schema)-1]
}

func TestPgSchema(t *testing.T) {
	p := PgSchema{"a", "b", "c"}
	s := removeKthElement(p, 1)
	fmt.Println(s)
}

type JobInfo struct {
	JobId         string
	JobName       string
	Priority      uint
	JobState      string
	Reason        string
	Partition     string
	NodeList      string
	StartTime     time.Time
	EndTime       time.Time
	LastSchedEval time.Time
	RunTime       string
	TaskNum       int
	TresPerJob    string
	NodeNum       int
	TresPerNode   string
}
type ListNodesReq struct {
	Partition string `json:"partition" form:"partition"`
}

type NodesUri struct {
	Name string `uri:"name"`
}

type Gpu struct {
	Total int    `json:"total"`
	Alloc int    `json:"alloc"`
	Type  string `json:"type"`
}

type Node1 struct {
	NodeName   string
	Arch       string
	Partitions []string
	State      string
	Os         string
	CpuTotal   int
	CpuAlloc   int
	MemTotal   int
	MemAlloc   int
	Gpu        []Gpu
}

type ListNodes struct {
	Nodes []Node1 `json:"nodes"`
}

const (
	SECTION_TYPE_PARTITION = "PartitionName"
	SECTION_TYPE_NODE      = "NodeName"
	SECTION_TYPE_JOB       = "JobId"
)

func TestNodeName(t *testing.T) {
	s := SlurmCommand{}
	jobInfos, err := s.ScontrolShowJobs()
	if err != nil {
		fmt.Errorf("GetGpuInfo err = %s", err)
	} else {
		for _, job := range jobInfos {
			//nodes := strings.Split(job.NodeList, ",")
			nodes := getNodeList(job.NodeList)
			if !IsContain(nodes, "dgx019") ||
				job.JobState != strings.ToUpper("running") {
				continue
			}
		}
	}
}

func getNodeList(nodeListStr string) []string {
	if nodeListStr == "(null)" {
		return nil
	}
	leftTrim := strings.Index(nodeListStr, "[")
	rightTrim := strings.Index(nodeListStr, "]")
	if leftTrim != -1 && rightTrim != -1 {
		nodeListTrim := nodeListStr[leftTrim+1 : rightTrim]
		prefix := nodeListStr[:leftTrim]
		var nodeList []string
		if strings.Contains(nodeListTrim, ",") { //以一个逗号分割
			numList := strings.Split(nodeListTrim, ",")
			for _, item := range numList {
				nodeList = append(nodeList, fmt.Sprintf("%s%s", prefix, item))
			}
			return nodeList
		} else if strings.Contains(nodeListTrim, "-") { //如果是一段例如  dxg[01-09]
			nodeListIndexes := strings.Split(nodeListTrim, "-")
			start, _ := strconv.Atoi(nodeListIndexes[0])

			end, _ := strconv.Atoi(nodeListIndexes[1])
			//l := len(nodeListIndexes[1])//看下多少位，计算需要补充多少个0

			for i := start; i <= end; i++ {
				nodeList = append(nodeList, fmt.Sprintf("%s%s", prefix, fmt.Sprintf("%03d", i)))
			}
			return nodeList
		}
	}
	fmt.Println(nodeListStr)
	return []string{nodeListStr}
}

func IsContain(items []string, item string) bool {
	for _, eachItem := range items {
		if eachItem == item {
			return true
		}
	}
	return false
}

func StringToTime(t string, layout string) (*time.Time, error) {
	cc, err := time.Parse(layout, t)
	if err != nil {
		return nil, err
	}
	return &cc, nil
}

type UnixTime time.Time
type Timespan time.Duration

var Layout = "2006-01-02 15:04:05"
var JupyterTimeLayout = "2006-01-02T15:04:05.999999Z"
var SlurmTimeLayout = "2006-01-02T15:04:05"
var DiffLayout = "15:04:05"

func (s *SlurmCommand) kvMap(section Section) (map[string]string, error) {
	m := make(map[string]string)
	kvs := s.GetKeyValue(section.Buffer.Bytes())
	if len(kvs) == 0 {
		return m, errors.New("dsadas")
	}

	// 提取 map
	for _, v := range kvs {
		m[v.Key] = v.Value
	}
	return m, nil
}

func (s *SlurmCommand) ScontrolShowJobs() ([]JobInfo, error) {
	//fmt.Debugf("exec output = [\n%s]", string(output))
	outputSections, err := s.GetSections([]byte(output), SECTION_TYPE_JOB)
	if err != nil {
		return nil, err
	}

	var jobInfos []JobInfo

	total := 0
	for _, outputSection := range *outputSections {
		m, err := s.kvMap(outputSection)
		if err != nil {
			fmt.Errorf("kvMap err = %s", err)
			continue
		}

		info := &JobInfo{
			JobId:     m["JobId"],
			JobName:   m["JobName"],
			JobState:  m["JobState"],
			Reason:    m["Reason"],
			Partition: m["Partition"],
			NodeList:  m["NodeList"],
		}

		u64, _ := strconv.ParseUint(m["Priority"], 10, 32)
		info.Priority = uint(u64)

		// run time
		info.RunTime = m["RunTime"]

		// start time
		st, err := StringToTime(m["SubmitTime"], SlurmTimeLayout)
		if err != nil {
			info.StartTime = time.Time{}
		} else {
			info.StartTime = *st
		}

		// end time
		et, err := StringToTime(m["EndTime"], SlurmTimeLayout)
		if err != nil {
			info.EndTime = time.Time{}
		} else {
			info.EndTime = *et
		}

		// LastSchedEval time 最后的调度评估时间，LastSchedEval - StartTime可近似作为等待时长
		evalTime, err := StringToTime(m["LastSchedEval"], SlurmTimeLayout)
		if err != nil {
			info.LastSchedEval = time.Time{}
		} else {
			info.LastSchedEval = *evalTime
		}

		info.TaskNum, _ = strconv.Atoi(m["NumTasks"])
		if _, ok := m["TresPerJob"]; ok {
			info.TresPerJob = m["TresPerJob"]
		}

		info.NodeNum, _ = strconv.Atoi(m["NumNodes"])
		if _, ok := m["TresPerNode"]; ok {
			info.TresPerNode = m["TresPerNode"]
		}
		total += s.gpuCount(info.TresPerNode, info.TaskNum)
		jobInfos = append(jobInfos, *info)
	}
	return jobInfos, nil
}

type Section struct {
	Type   string
	Buffer bytes.Buffer
}

func emptyLine(line []byte) bool {
	strLine := string(line)
	str := strings.Trim(strLine, OUTPUT_LINE_TRIM_SEP)
	return len(str) == 0
}

type SlurmCommand struct {
}

func (s *SlurmCommand) gpuCount(tres string, taskNum int) int {
	info := strings.Split(tres, ":")
	if len(info) != 3 {
		return 0
	}

	perJob, _ := strconv.Atoi(info[2])
	return perJob * taskNum
}
func (s *SlurmCommand) GetSections(cmdOutput []byte, t string) (*[]Section, error) {
	var sections []Section
	// 哨兵，index 从-1开始，以便后面统一 counter++ 的逻辑
	counter := -1

	bytesReader := bytes.NewReader(cmdOutput)
	bufReader := bufio.NewReader(bytesReader)
	var currentSection *Section = nil
	for {
		a, _, c := bufReader.ReadLine()
		if c == io.EOF {
			break
		}

		// skip to empty line
		if emptyLine(a) {
			continue
		}

		// 一些特殊处理的用例，比较恶心
		//if t == SECTION_TYPE_NODE && s.isOsLine(string(a)) {
		//	fmt.Printf("line = %s\n", string(a))
		//	a = []byte(s.CompactLine(string(a)))
		//}

		// 判断是否 section 的第1行：第1个 Key 等于 t，这里的 t 为 SECTION_TYPE_PARTITION，等等
		// 开始 1 个新的 section
		kvs := s.GetKeyValue(a)
		if len(kvs) == 0 {
			continue
		}

		if kvs[0].Key == t {
			// 保存 section
			if currentSection != nil {
				currentSection.Type = t
				sections = append(sections, *currentSection)
			}
			// 下一个 section
			counter++
			currentSection = &Section{}
		}

		if currentSection != nil {
			_, _ = currentSection.Buffer.Write(a)
		}
	}

	if currentSection != nil {
		currentSection.Type = t
		sections = append(sections, *currentSection)
	}

	return &sections, nil
}

const (
	OUTPUT_LINE_TRIM_SEP = "\t "
)

func (s *SlurmCommand) GetKeyValue(line []byte) []Kv {
	var kvArray []Kv

	strLine := string(line)
	// 分割每一行的键值对
	parts := strings.Fields(strLine)
	for _, part := range parts {
		kv := strings.SplitN(part, "=", 2)
		if len(kv) == 2 {
			key, value := kv[0], kv[1]
			kvArray = append(kvArray, Kv{
				Key:   key,
				Value: value,
			})
		}
	}
	return kvArray
}

//	func (s *SlurmCommand) GetKeyValue(line []byte) []Kv {
//		var kvArray []Kv
//
//		strLine := string(line)
//		str := strings.Trim(strLine, OUTPUT_LINE_TRIM_SEP)
//		str = s.RemoveExtraSpace(str)
//
//		kvs := strings.Split(str, " ")
//		for _, v := range kvs {
//			kv := strings.Split(v, "=")
//			if len(kv) != 2 {
//				continue
//			}
//			kvArray = append(kvArray, Kv{
//				Key:   kv[0],
//				Value: kv[1],
//			})
//		}
//		return kvArray
//	}
func (s *SlurmCommand) RemoveExtraSpace(line string) string {
	reg := regexp.MustCompile("\\s+")
	return reg.ReplaceAllString(line, " ")
}

var (
	output = `JobId=8477 JobName=train-93778f97-a3b0-4d7d-bdab-b695e51c6806.script
   UserId=aion(1002) GroupId=aion(1002) MCS_label=N/A
   Priority=4294899821 Nice=0 Account=(null) QOS=normal
   JobState=RUNNING Reason=None Dependency=(null)
   Requeue=1 Restarts=0 BatchFlag=1 Reboot=0 ExitCode=0:0
   RunTime=6-06:43:15 TimeLimit=UNLIMITED TimeMin=N/A
   SubmitTime=2023-10-06T09:30:29 EligibleTime=2023-10-06T09:30:29
   AccrueTime=2023-10-06T09:30:29
   StartTime=2023-10-06T09:30:29 EndTime=Unknown Deadline=N/A
   SuspendTime=None SecsPreSuspend=0 LastSchedEval=2023-10-06T09:30:29 Scheduler=Main
   Partition=defq AllocNode:Sid=login-02:8895
   ReqNodeList=(null) ExcNodeList=(null)
   NodeList=dgx019
   BatchHost=dgx019
   NumNodes=1 NumCPUs=4 NumTasks=1 CPUs/Task=4 ReqB:S:C:T=0:0:*:*
   TRES=cpu=4,node=1,billing=4,gres/gpu=1
   Socks/Node=* NtasksPerN:B:S:C=0:0:*:* CoreSpec=*
   MinCPUsNode=4 MinMemoryNode=4G MinTmpDiskNode=0
   Features=(null) DelayBoot=00:00:00
   OverSubscribe=OK Contiguous=0 Licenses=(null) Network=(null)
   Command=/home/aion/another-scheduler/job/train-93778f97-a3b0-4d7d-bdab-b695e51c6806/train-93778f97-a3b0-4d7d-bdab-b695e51c6806.script
   WorkDir=/home/aion/another-scheduler/job/train-93778f97-a3b0-4d7d-bdab-b695e51c6806
   StdErr=/home/aion/another-scheduler/job/train-93778f97-a3b0-4d7d-bdab-b695e51c6806/train-93778f97-a3b0-4d7d-bdab-b695e51c6806.log
   StdIn=/dev/null
   StdOut=/home/aion/another-scheduler/job/train-93778f97-a3b0-4d7d-bdab-b695e51c6806/train-93778f97-a3b0-4d7d-bdab-b695e51c6806.log
   Power=
   TresPerNode=gres:gpu:1


JobId=8551 JobName=train-11bf3ba5-d134-4f35-ad9b-b803efae4838.script
   UserId=aion(1002) GroupId=aion(1002) MCS_label=N/A
   Priority=4294899747 Nice=0 Account=(null) QOS=normal
   JobState=RUNNING Reason=None Dependency=(null)
   Requeue=1 Restarts=0 BatchFlag=1 Reboot=0 ExitCode=0:0
   RunTime=5-03:08:41 TimeLimit=UNLIMITED TimeMin=N/A
   SubmitTime=2023-10-07T13:05:02 EligibleTime=2023-10-07T13:05:02
   AccrueTime=2023-10-07T13:05:02
   StartTime=2023-10-07T13:05:03 EndTime=Unknown Deadline=N/A
   SuspendTime=None SecsPreSuspend=0 LastSchedEval=2023-10-07T13:05:03 Scheduler=Main
   Partition=defq AllocNode:Sid=login-02:8895
   ReqNodeList=(null) ExcNodeList=(null)
   NodeList=dgx007
   BatchHost=dgx007
   NumNodes=1 NumCPUs=4 NumTasks=1 CPUs/Task=4 ReqB:S:C:T=0:0:*:*
   TRES=cpu=4,node=1,billing=4,gres/gpu=1
   Socks/Node=* NtasksPerN:B:S:C=0:0:*:* CoreSpec=*
   MinCPUsNode=4 MinMemoryNode=4G MinTmpDiskNode=0
   Features=(null) DelayBoot=00:00:00
   OverSubscribe=OK Contiguous=0 Licenses=(null) Network=(null)
   Command=/home/aion/another-scheduler/job/train-11bf3ba5-d134-4f35-ad9b-b803efae4838/train-11bf3ba5-d134-4f35-ad9b-b803efae4838.script
   WorkDir=/home/aion/another-scheduler/job/train-11bf3ba5-d134-4f35-ad9b-b803efae4838
   StdErr=/home/aion/another-scheduler/job/train-11bf3ba5-d134-4f35-ad9b-b803efae4838/train-11bf3ba5-d134-4f35-ad9b-b803efae4838.log
   StdIn=/dev/null
   StdOut=/home/aion/another-scheduler/job/train-11bf3ba5-d134-4f35-ad9b-b803efae4838/train-11bf3ba5-d134-4f35-ad9b-b803efae4838.log
   Power=
   TresPerNode=gres:gpu:1


JobId=8682 JobName=train-7d32f42e-3d9a-46ff-a757-f7d06cee3758.script
   UserId=aion(1002) GroupId=aion(1002) MCS_label=N/A
   Priority=4294899616 Nice=0 Account=(null) QOS=normal
   JobState=RUNNING Reason=None Dependency=(null)
   Requeue=1 Restarts=0 BatchFlag=1 Reboot=0 ExitCode=0:0
   RunTime=07:14:30 TimeLimit=UNLIMITED TimeMin=N/A
   SubmitTime=2023-10-12T08:59:13 EligibleTime=2023-10-12T08:59:13
   AccrueTime=2023-10-12T08:59:13
   StartTime=2023-10-12T08:59:14 EndTime=Unknown Deadline=N/A
   SuspendTime=None SecsPreSuspend=0 LastSchedEval=2023-10-12T08:59:14 Scheduler=Main
   Partition=defq AllocNode:Sid=login-01:2393654
   ReqNodeList=(null) ExcNodeList=(null)
   NodeList=dgx019
   BatchHost=dgx019
   NumNodes=1 NumCPUs=4 NumTasks=1 CPUs/Task=4 ReqB:S:C:T=0:0:*:*
   TRES=cpu=4,node=1,billing=4,gres/gpu=1
   Socks/Node=* NtasksPerN:B:S:C=0:0:*:* CoreSpec=*
   MinCPUsNode=4 MinMemoryNode=4G MinTmpDiskNode=0
   Features=(null) DelayBoot=00:00:00
   OverSubscribe=OK Contiguous=0 Licenses=(null) Network=(null)
   Command=/home/aion/another-scheduler/job/train-7d32f42e-3d9a-46ff-a757-f7d06cee3758/train-7d32f42e-3d9a-46ff-a757-f7d06cee3758.script
   WorkDir=/home/aion/another-scheduler/job/train-7d32f42e-3d9a-46ff-a757-f7d06cee3758
   StdErr=/home/aion/another-scheduler/job/train-7d32f42e-3d9a-46ff-a757-f7d06cee3758/train-7d32f42e-3d9a-46ff-a757-f7d06cee3758.log
   StdIn=/dev/null
   StdOut=/home/aion/another-scheduler/job/train-7d32f42e-3d9a-46ff-a757-f7d06cee3758/train-7d32f42e-3d9a-46ff-a757-f7d06cee3758.log
   Power=
   TresPerNode=gres:gpu:1


JobId=8683 JobName=train-c72aab85-2b5e-4226-802f-d5ae41230ef6.script
   UserId=aion(1002) GroupId=aion(1002) MCS_label=N/A
   Priority=4294899615 Nice=0 Account=(null) QOS=normal
   JobState=RUNNING Reason=None Dependency=(null)
   Requeue=1 Restarts=0 BatchFlag=1 Reboot=0 ExitCode=0:0
   RunTime=07:10:12 TimeLimit=UNLIMITED TimeMin=N/A
   SubmitTime=2023-10-12T09:03:32 EligibleTime=2023-10-12T09:03:32
   AccrueTime=2023-10-12T09:03:32
   StartTime=2023-10-12T09:03:32 EndTime=Unknown Deadline=N/A
   SuspendTime=None SecsPreSuspend=0 LastSchedEval=2023-10-12T09:03:32 Scheduler=Main
   Partition=defq AllocNode:Sid=login-01:2393654
   ReqNodeList=(null) ExcNodeList=(null)
   NodeList=dgx019
   BatchHost=dgx019
   NumNodes=1 NumCPUs=4 NumTasks=1 CPUs/Task=4 ReqB:S:C:T=0:0:*:*
   TRES=cpu=4,node=1,billing=4,gres/gpu=1
   Socks/Node=* NtasksPerN:B:S:C=0:0:*:* CoreSpec=*
   MinCPUsNode=4 MinMemoryNode=4G MinTmpDiskNode=0
   Features=(null) DelayBoot=00:00:00
   OverSubscribe=OK Contiguous=0 Licenses=(null) Network=(null)
   Command=/home/aion/another-scheduler/job/train-c72aab85-2b5e-4226-802f-d5ae41230ef6/train-c72aab85-2b5e-4226-802f-d5ae41230ef6.script
   WorkDir=/home/aion/another-scheduler/job/train-c72aab85-2b5e-4226-802f-d5ae41230ef6
   StdErr=/home/aion/another-scheduler/job/train-c72aab85-2b5e-4226-802f-d5ae41230ef6/train-c72aab85-2b5e-4226-802f-d5ae41230ef6.log
   StdIn=/dev/null
   StdOut=/home/aion/another-scheduler/job/train-c72aab85-2b5e-4226-802f-d5ae41230ef6/train-c72aab85-2b5e-4226-802f-d5ae41230ef6.log
   Power=
   TresPerNode=gres:gpu:1


JobId=8684 JobName=train-c72aab85-2b5e-4226-802f-d5ae41230ef6.script
   UserId=aion(1002) GroupId=aion(1002) MCS_label=N/A
   Priority=4294899614 Nice=0 Account=(null) QOS=normal
   JobState=RUNNING Reason=None Dependency=(null)
   Requeue=1 Restarts=0 BatchFlag=1 Reboot=0 ExitCode=0:0
   RunTime=07:10:11 TimeLimit=UNLIMITED TimeMin=N/A
   SubmitTime=2023-10-12T09:03:32 EligibleTime=2023-10-12T09:03:32
   AccrueTime=2023-10-12T09:03:32
   StartTime=2023-10-12T09:03:33 EndTime=Unknown Deadline=N/A
   SuspendTime=None SecsPreSuspend=0 LastSchedEval=2023-10-12T09:03:33 Scheduler=Backfill
   Partition=defq AllocNode:Sid=login-02:8895
   ReqNodeList=(null) ExcNodeList=(null)
   NodeList=dgx019
   BatchHost=dgx019
   NumNodes=1 NumCPUs=4 NumTasks=1 CPUs/Task=4 ReqB:S:C:T=0:0:*:*
   TRES=cpu=4,node=1,billing=4,gres/gpu=1
   Socks/Node=* NtasksPerN:B:S:C=0:0:*:* CoreSpec=*
   MinCPUsNode=4 MinMemoryNode=4G MinTmpDiskNode=0
   Features=(null) DelayBoot=00:00:00
   OverSubscribe=OK Contiguous=0 Licenses=(null) Network=(null)
   Command=/home/aion/another-scheduler/job/train-c72aab85-2b5e-4226-802f-d5ae41230ef6/train-c72aab85-2b5e-4226-802f-d5ae41230ef6.script
   WorkDir=/home/aion/another-scheduler/job/train-c72aab85-2b5e-4226-802f-d5ae41230ef6
   StdErr=/home/aion/another-scheduler/job/train-c72aab85-2b5e-4226-802f-d5ae41230ef6/train-c72aab85-2b5e-4226-802f-d5ae41230ef6.log
   StdIn=/dev/null
   StdOut=/home/aion/another-scheduler/job/train-c72aab85-2b5e-4226-802f-d5ae41230ef6/train-c72aab85-2b5e-4226-802f-d5ae41230ef6.log
   Power=
   TresPerNode=gres:gpu:1


JobId=8687 JobName=train-e2859539-f869-4566-bdf5-c87a44e49206.script
   UserId=aion(1002) GroupId=aion(1002) MCS_label=N/A
   Priority=4294899611 Nice=0 Account=(null) QOS=normal
   JobState=RUNNING Reason=None Dependency=(null)
   Requeue=1 Restarts=0 BatchFlag=1 Reboot=0 ExitCode=0:0
   RunTime=06:52:54 TimeLimit=UNLIMITED TimeMin=N/A
   SubmitTime=2023-10-12T09:20:50 EligibleTime=2023-10-12T09:20:50
   AccrueTime=2023-10-12T09:20:50
   StartTime=2023-10-12T09:20:50 EndTime=Unknown Deadline=N/A
   SuspendTime=None SecsPreSuspend=0 LastSchedEval=2023-10-12T09:20:50 Scheduler=Backfill
   Partition=defq AllocNode:Sid=login-01:2393654
   ReqNodeList=(null) ExcNodeList=(null)
   NodeList=dgx[001-002]
   BatchHost=dgx001
   NumNodes=2 NumCPUs=256 NumTasks=2 CPUs/Task=128 ReqB:S:C:T=0:0:*:*
   TRES=cpu=256,node=2,billing=256,gres/gpu=16
   Socks/Node=* NtasksPerN:B:S:C=0:0:*:* CoreSpec=*
   MinCPUsNode=128 MinMemoryNode=128G MinTmpDiskNode=0
   Features=(null) DelayBoot=00:00:00
   OverSubscribe=OK Contiguous=0 Licenses=(null) Network=(null)
   Command=/home/aion/another-scheduler/job/train-e2859539-f869-4566-bdf5-c87a44e49206/train-e2859539-f869-4566-bdf5-c87a44e49206.script
   WorkDir=/home/aion/another-scheduler/job/train-e2859539-f869-4566-bdf5-c87a44e49206
   StdErr=/home/aion/another-scheduler/job/train-e2859539-f869-4566-bdf5-c87a44e49206/train-e2859539-f869-4566-bdf5-c87a44e49206.log
   StdIn=/dev/null
   StdOut=/home/aion/another-scheduler/job/train-e2859539-f869-4566-bdf5-c87a44e49206/train-e2859539-f869-4566-bdf5-c87a44e49206.log
   Power=
   TresPerNode=gres:gpu:8


JobId=8688 JobName=train-592fca75-9c5e-47fc-b540-4435a5d0dbcb.script
   UserId=aion(1002) GroupId=aion(1002) MCS_label=N/A
   Priority=4294899610 Nice=0 Account=(null) QOS=normal
   JobState=RUNNING Reason=None Dependency=(null)
   Requeue=1 Restarts=0 BatchFlag=1 Reboot=0 ExitCode=0:0
   RunTime=06:52:39 TimeLimit=UNLIMITED TimeMin=N/A
   SubmitTime=2023-10-12T09:21:05 EligibleTime=2023-10-12T09:21:05
   AccrueTime=2023-10-12T09:21:05
   StartTime=2023-10-12T09:21:05 EndTime=Unknown Deadline=N/A
   SuspendTime=None SecsPreSuspend=0 LastSchedEval=2023-10-12T09:21:05 Scheduler=Main
   Partition=defq AllocNode:Sid=login-01:2393654
   ReqNodeList=(null) ExcNodeList=(null)
   NodeList=dgx[004-005]
   BatchHost=dgx004
   NumNodes=2 NumCPUs=256 NumTasks=2 CPUs/Task=128 ReqB:S:C:T=0:0:*:*
   TRES=cpu=256,node=2,billing=256,gres/gpu=16
   Socks/Node=* NtasksPerN:B:S:C=0:0:*:* CoreSpec=*
   MinCPUsNode=128 MinMemoryNode=128G MinTmpDiskNode=0
   Features=(null) DelayBoot=00:00:00
   OverSubscribe=OK Contiguous=0 Licenses=(null) Network=(null)
   Command=/home/aion/another-scheduler/job/train-592fca75-9c5e-47fc-b540-4435a5d0dbcb/train-592fca75-9c5e-47fc-b540-4435a5d0dbcb.script
   WorkDir=/home/aion/another-scheduler/job/train-592fca75-9c5e-47fc-b540-4435a5d0dbcb
   StdErr=/home/aion/another-scheduler/job/train-592fca75-9c5e-47fc-b540-4435a5d0dbcb/train-592fca75-9c5e-47fc-b540-4435a5d0dbcb.log
   StdIn=/dev/null
   StdOut=/home/aion/another-scheduler/job/train-592fca75-9c5e-47fc-b540-4435a5d0dbcb/train-592fca75-9c5e-47fc-b540-4435a5d0dbcb.log
   Power=
   TresPerNode=gres:gpu:8


JobId=8689 JobName=train-a643b348-dcd9-4527-aae5-c873abb35147.script
   UserId=aion(1002) GroupId=aion(1002) MCS_label=N/A
   Priority=4294899609 Nice=0 Account=(null) QOS=normal
   JobState=RUNNING Reason=None Dependency=(null)
   Requeue=1 Restarts=0 BatchFlag=1 Reboot=0 ExitCode=0:0
   RunTime=06:52:29 TimeLimit=UNLIMITED TimeMin=N/A
   SubmitTime=2023-10-12T09:21:15 EligibleTime=2023-10-12T09:21:15
   AccrueTime=2023-10-12T09:21:15
   StartTime=2023-10-12T09:21:15 EndTime=Unknown Deadline=N/A
   SuspendTime=None SecsPreSuspend=0 LastSchedEval=2023-10-12T09:21:15 Scheduler=Main
   Partition=defq AllocNode:Sid=login-01:2393654
   ReqNodeList=(null) ExcNodeList=(null)
   NodeList=dgx[008-009]
   BatchHost=dgx008
   NumNodes=2 NumCPUs=256 NumTasks=2 CPUs/Task=128 ReqB:S:C:T=0:0:*:*
   TRES=cpu=256,node=2,billing=256,gres/gpu=16
   Socks/Node=* NtasksPerN:B:S:C=0:0:*:* CoreSpec=*
   MinCPUsNode=128 MinMemoryNode=128G MinTmpDiskNode=0
   Features=(null) DelayBoot=00:00:00
   OverSubscribe=OK Contiguous=0 Licenses=(null) Network=(null)
   Command=/home/aion/another-scheduler/job/train-a643b348-dcd9-4527-aae5-c873abb35147/train-a643b348-dcd9-4527-aae5-c873abb35147.script
   WorkDir=/home/aion/another-scheduler/job/train-a643b348-dcd9-4527-aae5-c873abb35147
   StdErr=/home/aion/another-scheduler/job/train-a643b348-dcd9-4527-aae5-c873abb35147/train-a643b348-dcd9-4527-aae5-c873abb35147.log
   StdIn=/dev/null
   StdOut=/home/aion/another-scheduler/job/train-a643b348-dcd9-4527-aae5-c873abb35147/train-a643b348-dcd9-4527-aae5-c873abb35147.log
   Power=
   TresPerNode=gres:gpu:8


JobId=8690 JobName=train-cb486791-ddcf-4559-bb76-8476474a535a.script
   UserId=aion(1002) GroupId=aion(1002) MCS_label=N/A
   Priority=4294899608 Nice=0 Account=(null) QOS=normal
   JobState=RUNNING Reason=None Dependency=(null)
   Requeue=1 Restarts=0 BatchFlag=1 Reboot=0 ExitCode=0:0
   RunTime=06:52:20 TimeLimit=UNLIMITED TimeMin=N/A
   SubmitTime=2023-10-12T09:21:24 EligibleTime=2023-10-12T09:21:24
   AccrueTime=2023-10-12T09:21:24
   StartTime=2023-10-12T09:21:24 EndTime=Unknown Deadline=N/A
   SuspendTime=None SecsPreSuspend=0 LastSchedEval=2023-10-12T09:21:24 Scheduler=Main
   Partition=defq AllocNode:Sid=login-01:2393654
   ReqNodeList=(null) ExcNodeList=(null)
   NodeList=dgx[010-011]
   BatchHost=dgx010
   NumNodes=2 NumCPUs=256 NumTasks=2 CPUs/Task=128 ReqB:S:C:T=0:0:*:*
   TRES=cpu=256,node=2,billing=256,gres/gpu=16
   Socks/Node=* NtasksPerN:B:S:C=0:0:*:* CoreSpec=*
   MinCPUsNode=128 MinMemoryNode=128G MinTmpDiskNode=0
   Features=(null) DelayBoot=00:00:00
   OverSubscribe=OK Contiguous=0 Licenses=(null) Network=(null)
   Command=/home/aion/another-scheduler/job/train-cb486791-ddcf-4559-bb76-8476474a535a/train-cb486791-ddcf-4559-bb76-8476474a535a.script
   WorkDir=/home/aion/another-scheduler/job/train-cb486791-ddcf-4559-bb76-8476474a535a
   StdErr=/home/aion/another-scheduler/job/train-cb486791-ddcf-4559-bb76-8476474a535a/train-cb486791-ddcf-4559-bb76-8476474a535a.log
   StdIn=/dev/null
   StdOut=/home/aion/another-scheduler/job/train-cb486791-ddcf-4559-bb76-8476474a535a/train-cb486791-ddcf-4559-bb76-8476474a535a.log
   Power=
   TresPerNode=gres:gpu:8


JobId=8693 JobName=train-53421f4e-c965-4ded-b96f-ff15ab5a2df6.script
   UserId=aion(1002) GroupId=aion(1002) MCS_label=N/A
   Priority=4294899605 Nice=0 Account=(null) QOS=normal
   JobState=RUNNING Reason=None Dependency=(null)
   Requeue=1 Restarts=0 BatchFlag=1 Reboot=0 ExitCode=0:0
   RunTime=06:16:24 TimeLimit=UNLIMITED TimeMin=N/A
   SubmitTime=2023-10-12T09:57:17 EligibleTime=2023-10-12T09:57:17
   AccrueTime=2023-10-12T09:57:17
   StartTime=2023-10-12T09:57:20 EndTime=Unknown Deadline=N/A
   SuspendTime=None SecsPreSuspend=0 LastSchedEval=2023-10-12T09:57:20 Scheduler=Main
   Partition=defq AllocNode:Sid=login-01:2393654
   ReqNodeList=(null) ExcNodeList=(null)
   NodeList=dgx[014-015]
   BatchHost=dgx014
   NumNodes=2 NumCPUs=256 NumTasks=2 CPUs/Task=128 ReqB:S:C:T=0:0:*:*
   TRES=cpu=256,node=2,billing=256,gres/gpu=16
   Socks/Node=* NtasksPerN:B:S:C=0:0:*:* CoreSpec=*
   MinCPUsNode=128 MinMemoryNode=128G MinTmpDiskNode=0
   Features=(null) DelayBoot=00:00:00
   OverSubscribe=OK Contiguous=0 Licenses=(null) Network=(null)
   Command=/home/aion/another-scheduler/job/train-53421f4e-c965-4ded-b96f-ff15ab5a2df6/train-53421f4e-c965-4ded-b96f-ff15ab5a2df6.script
   WorkDir=/home/aion/another-scheduler/job/train-53421f4e-c965-4ded-b96f-ff15ab5a2df6
   StdErr=/home/aion/another-scheduler/job/train-53421f4e-c965-4ded-b96f-ff15ab5a2df6/train-53421f4e-c965-4ded-b96f-ff15ab5a2df6.log
   StdIn=/dev/null
   StdOut=/home/aion/another-scheduler/job/train-53421f4e-c965-4ded-b96f-ff15ab5a2df6/train-53421f4e-c965-4ded-b96f-ff15ab5a2df6.log
   Power=
   TresPerNode=gres:gpu:8


JobId=8696 JobName=train-bfa62015-9254-41b7-8593-70207443cf01.script
   UserId=aion(1002) GroupId=aion(1002) MCS_label=N/A
   Priority=4294899602 Nice=0 Account=(null) QOS=normal
   JobState=RUNNING Reason=None Dependency=(null)
   Requeue=1 Restarts=0 BatchFlag=1 Reboot=0 ExitCode=0:0
   RunTime=05:36:06 TimeLimit=UNLIMITED TimeMin=N/A
   SubmitTime=2023-10-12T10:37:38 EligibleTime=2023-10-12T10:37:38
   AccrueTime=2023-10-12T10:37:38
   StartTime=2023-10-12T10:37:38 EndTime=Unknown Deadline=N/A
   SuspendTime=None SecsPreSuspend=0 LastSchedEval=2023-10-12T10:37:38 Scheduler=Main
   Partition=defq AllocNode:Sid=login-01:2393654
   ReqNodeList=(null) ExcNodeList=(null)
   NodeList=dgx012
   BatchHost=dgx012
   NumNodes=1 NumCPUs=256 NumTasks=1 CPUs/Task=256 ReqB:S:C:T=0:0:*:*
   TRES=cpu=256,node=1,billing=256,gres/gpu=8
   Socks/Node=* NtasksPerN:B:S:C=0:0:*:* CoreSpec=*
   MinCPUsNode=256 MinMemoryNode=1987G MinTmpDiskNode=0
   Features=(null) DelayBoot=00:00:00
   OverSubscribe=OK Contiguous=0 Licenses=(null) Network=(null)
   Command=/home/aion/another-scheduler/job/train-bfa62015-9254-41b7-8593-70207443cf01/train-bfa62015-9254-41b7-8593-70207443cf01.script
   WorkDir=/home/aion/another-scheduler/job/train-bfa62015-9254-41b7-8593-70207443cf01
   StdErr=/home/aion/another-scheduler/job/train-bfa62015-9254-41b7-8593-70207443cf01/train-bfa62015-9254-41b7-8593-70207443cf01.log
   StdIn=/dev/null
   StdOut=/home/aion/another-scheduler/job/train-bfa62015-9254-41b7-8593-70207443cf01/train-bfa62015-9254-41b7-8593-70207443cf01.log
   Power=
   TresPerNode=gres:gpu:8


JobId=8702 JobName=train-d9d73df3-8b90-46fa-8920-56bef85ebc4f.script
   UserId=aion(1002) GroupId=aion(1002) MCS_label=N/A
   Priority=4294899596 Nice=0 Account=(null) QOS=normal
   JobState=RUNNING Reason=None Dependency=(null)
   Requeue=1 Restarts=0 BatchFlag=1 Reboot=0 ExitCode=0:0
   RunTime=02:42:52 TimeLimit=UNLIMITED TimeMin=N/A
   SubmitTime=2023-10-12T13:30:52 EligibleTime=2023-10-12T13:30:52
   AccrueTime=2023-10-12T13:30:52
   StartTime=2023-10-12T13:30:52 EndTime=Unknown Deadline=N/A
   SuspendTime=None SecsPreSuspend=0 LastSchedEval=2023-10-12T13:30:52 Scheduler=Main
   Partition=defq AllocNode:Sid=login-01:2393654
   ReqNodeList=(null) ExcNodeList=(null)
   NodeList=dgx007
   BatchHost=dgx007
   NumNodes=1 NumCPUs=4 NumTasks=1 CPUs/Task=4 ReqB:S:C:T=0:0:*:*
   TRES=cpu=4,node=1,billing=4,gres/gpu=1
   Socks/Node=* NtasksPerN:B:S:C=0:0:*:* CoreSpec=*
   MinCPUsNode=4 MinMemoryNode=4G MinTmpDiskNode=0
   Features=(null) DelayBoot=00:00:00
   OverSubscribe=OK Contiguous=0 Licenses=(null) Network=(null)
   Command=/home/aion/another-scheduler/job/train-d9d73df3-8b90-46fa-8920-56bef85ebc4f/train-d9d73df3-8b90-46fa-8920-56bef85ebc4f.script
   WorkDir=/home/aion/another-scheduler/job/train-d9d73df3-8b90-46fa-8920-56bef85ebc4f
   StdErr=/home/aion/another-scheduler/job/train-d9d73df3-8b90-46fa-8920-56bef85ebc4f/train-d9d73df3-8b90-46fa-8920-56bef85ebc4f.log
   StdIn=/dev/null
   StdOut=/home/aion/another-scheduler/job/train-d9d73df3-8b90-46fa-8920-56bef85ebc4f/train-d9d73df3-8b90-46fa-8920-56bef85ebc4f.log
   Power=
   TresPerNode=gres:gpu:1


JobId=8703 JobName=train-54533922-e80a-4121-a985-f793e53e3ffd.script
   UserId=aion(1002) GroupId=aion(1002) MCS_label=N/A
   Priority=4294899595 Nice=0 Account=(null) QOS=normal
   JobState=RUNNING Reason=None Dependency=(null)
   Requeue=1 Restarts=0 BatchFlag=1 Reboot=0 ExitCode=0:0
   RunTime=00:33:35 TimeLimit=UNLIMITED TimeMin=N/A
   SubmitTime=2023-10-12T15:40:09 EligibleTime=2023-10-12T15:40:09
   AccrueTime=2023-10-12T15:40:09
   StartTime=2023-10-12T15:40:09 EndTime=Unknown Deadline=N/A
   SuspendTime=None SecsPreSuspend=0 LastSchedEval=2023-10-12T15:40:09 Scheduler=Backfill
   Partition=defq AllocNode:Sid=login-01:2393654
   ReqNodeList=(null) ExcNodeList=(null)
   NodeList=dgx[016-017]
   BatchHost=dgx016
   NumNodes=2 NumCPUs=256 NumTasks=2 CPUs/Task=128 ReqB:S:C:T=0:0:*:*
   TRES=cpu=256,node=2,billing=256,gres/gpu=16
   Socks/Node=* NtasksPerN:B:S:C=0:0:*:* CoreSpec=*
   MinCPUsNode=128 MinMemoryNode=128G MinTmpDiskNode=0
   Features=(null) DelayBoot=00:00:00
   OverSubscribe=OK Contiguous=0 Licenses=(null) Network=(null)
   Command=/home/aion/another-scheduler/job/train-54533922-e80a-4121-a985-f793e53e3ffd/train-54533922-e80a-4121-a985-f793e53e3ffd.script
   WorkDir=/home/aion/another-scheduler/job/train-54533922-e80a-4121-a985-f793e53e3ffd
   StdErr=/home/aion/another-scheduler/job/train-54533922-e80a-4121-a985-f793e53e3ffd/train-54533922-e80a-4121-a985-f793e53e3ffd.log
   StdIn=/dev/null
   StdOut=/home/aion/another-scheduler/job/train-54533922-e80a-4121-a985-f793e53e3ffd/train-54533922-e80a-4121-a985-f793e53e3ffd.log
   Power=
   TresPerNode=gres:gpu:8
`
)

func TestMap12(t *testing.T) {
	data := `NodeName=dgx001 Arch=x86_64 CoresPerSocket=1
   CPUAlloc=128 CPUTot=256 CPULoad=9.12
   AvailableFeatures=location=local
   ActiveFeatures=location=local
   Gres=gpu:A800:8(S:16-31,48-63,80-95,112-127,144-159,176-191,210-223,240-255)
   NodeAddr=dgx001 NodeHostName=dgx001 Version=21.08.8-2
   OS=Linux 5.4.0-156-generic #173-Ubuntu SMP Tue Jul 11 07:25:22 UTC 2023
   RealMemory=2034931 AllocMem=0 FreeMem=1644035 Sockets=256 Boards=1
   State=MIXED ThreadsPerCore=1 TmpDisk=0 Weight=1 Owner=N/A MCS_label=N/A
   Partitions=defq
   BootTime=2023-09-17T19:49:58 SlurmdStartTime=2023-09-17T19:53:31
   LastBusyTime=2023-10-12T15:42:00
   CfgTRES=cpu=256,mem=2034931M,billing=256,gres/gpu=8
   AllocTRES=cpu=128,gres/gpu=8
   CapWatts=n/a
   CurrentWatts=0 AveWatts=0
   ExtSensorsJoules=n/s ExtSensorsWatts=0 ExtSensorsTemp=n/s`

	lines := strings.Split(data, "\n")
	dataMap := make(map[string]string)

	for _, line := range lines {
		// 分割每一行的键值对
		parts := strings.Fields(line)
		for _, part := range parts {
			kv := strings.SplitN(part, "=", 2)
			if len(kv) == 2 {
				key, value := kv[0], kv[1]
				dataMap[key] = value
			}
		}
	}

	// 打印结果
	for key, value := range dataMap {
		fmt.Printf("%s: %s\n", key, value)
	}
}
