package main

import (
	"encoding/json"
	"fmt"
	"github.com/go-labs/internal/logging"
	"path"
	"path/filepath"
	"strconv"
	"sync"
	"testing"
	"time"
)

type Print func() string

func Adsa() string {
	return "A"
}

func Bdsa() string {
	return "B"
}

func Cdsa() string {
	return "C"
}

func TestFunc(t *testing.T) {
	fmt.Println(TempPrint(Adsa))
	fmt.Println(TempPrint(Bdsa))
	fmt.Println(TempPrint(Cdsa))

	fmt.Println(filepath.Base("/home/a"))
	fmt.Println(path.Join("/home", "a"))
}

func TempPrint(option Print) string {
	return option()
}

func Dsad() error {
	return nil
}

func TestFunc1(t *testing.T) {
	tmpCodePath := "/tmp/aiarts_launcher/"
	var launch_config = Train{
		Entry:      "mmdetection/tools/train.sh",
		SysParams:  nil,
		UserParams: nil,
	}
	filesuffix := path.Ext(launch_config.Entry)
	runPath := path.Join(tmpCodePath, launch_config.Entry)
	launcher := TrainLauncher{}
	if v, ok := suffixMap[filesuffix]; ok {
		launcher.preExecPro = v[0]
		for _, s := range v[1:] {
			launcher.execArgs = append(launcher.execArgs, s)
		}
		//launcher.execArgs = append(launcher.execArgs, runPath)
		launcher.execArgs = append(launcher.execArgs, launch_config.Entry)
	} else {
		launcher.preExecPro = runPath
	}
	fmt.Printf("%v", launcher)
}

type Train struct {
	Entry     string `json:"entry"`
	SysParams []struct {
		Arg   string      `json:"arg"`
		Name  string      `json:"name"`
		Type  string      `json:"type"`
		Value interface{} `json:"value"`
	} `json:"sysParams"`
	UserParams []struct {
		Desc     string `json:"desc"`
		Name     string `json:"name"`
		Type     string `json:"type"`
		Value    string `json:"value"`
		Default  string `json:"default"`
		Editable bool   `json:"editable"`
		Required bool   `json:"required"`
	} `json:"userParams"`
}

var suffixMap = map[string][]string{
	".py": {"python3"},
	".sh": {"bash"},
	".go": {"go", "run"},
}

type TrainLauncher struct {
	jobName string
	JobDesc string

	codeDir     string //代码所在路径
	codeId      string
	codeVersion string

	jobType    string   //任务类型 train eval
	jobId      string   //任务id
	clusterId  string   //集群id
	ExecPro    string   //执行程序名字
	preExecPro string   //执行程序启动程序
	execArgs   []string //运行参数

	isNeedSave bool
	//训练结果输出目录
	trainDir string
	//训练jobid
	trainJobId string
	//评估结果输出目录
	evalDir string
	//输出目录
	saveDir string
	//数据集目录
	datasetDir string
	datasetId  string
	//用户参数
	userParams []string

	flags uint64

	tags map[string]interface{}

	//上云要用到
	aomBackendServiceHost string
	aomBackendServicePort string
	aiLabAddr             string

	modelArtsSettingUrl string
	aiLabToken          string
	//internal
	runType string
}

func TestFunc2(t *testing.T) {
	now := time.Now().UnixNano()
	now = now / 1000000
	fmt.Println(now)
}

func TestFor(t *testing.T) {
	l := 101
	step := 20
	//x := l / step
	//left := l % step
	//for i := 0; i < x; i++ {
	//	start := i * step
	//	for j := 0; j < step; j++ {
	//		fmt.Println(start + j)
	//	}
	//}
	//if left > 0 {
	//	start := x * step
	//	for j := start; j <= l; j++ {
	//		fmt.Println(j)
	//	}
	//}
	rows := []int{}
	for i := 0; i <= l; i++ {
		rows = append(rows, i)
	}
	var sw = sync.WaitGroup{}
	for i, _ := range rows {
		if i == 0 {
			continue
		}
		if i != 1 && (i-1)%step == 0 { //每次最多发500条  满500 重置
			fmt.Printf("group start :%v ========================\n", rows[i])
			sw = sync.WaitGroup{}
		}
		sw.Add(1)
		go func(i int, rows []int) {
			defer sw.Done()
			fmt.Printf("row:%v\n", rows[i])
			time.Sleep(time.Second * 2)
		}(i, rows)
		if (i)%step == 0 || i == (len(rows)-1) {
			sw.Wait()
			fmt.Printf("group end :%v ========================\n", rows[i])
		}
	}
}

type AILabRunConfig struct {
	JobName string                  `json:"name"`
	JobType string                  `json:"jobType"`
	Config  map[string]JsonMetaData `json:"config"`
}
type JsonMetaData struct {
	//data map[string]interface{}
	data_str []byte
}

func (d *JsonMetaData) Empty() bool {
	return len(d.data_str) <= 2
}
func (d *JsonMetaData) MarshalJSON() ([]byte, error) {
	if len(d.data_str) == 0 {
		return []byte("null"), nil
	}
	return d.data_str, nil
}
func (d *JsonMetaData) UnmarshalJSON(b []byte) error {
	if len(b) >= 2 && (b[0] == '{' || b[0] == '[') {
		d.data_str = b
	} else {
		d.data_str = nil
	}
	return nil
}

func (d *JsonMetaData) Fetch(v interface{}) error {
	if d == nil {
		return nil
	}
	return json.Unmarshal([]byte(d.data_str), v)
}
func (d *JsonMetaData) Save(v interface{}) {
	d.data_str, _ = json.Marshal(v)
}
func TestBox(t *testing.T) {
	str := `{"runId":"apulis-iqi-analysis-0a9eb237-bf00-41a9-bb07-f8bff30e3bed","labId":1016277,"group":"","name":"apulis-iqi-analysis","num":0,"jobType":"apulis-iqi-analysis","creator":"hs-org","userId":48441424,"createdAt":1663574543264,"description":"","start":1663574543372,"deadline":0,"status":7,"extStatus":0,"msg":"","arch":"","cmd":["/start/aiarts_launcher"],"image":"harbor.apulis.cn:8443/algorithm/apulistech/recheck_analysis:1.2","config":{"param":{"orderBy":"generate_date_time DESC","pageNum":1,"dataPath":"20220919_16.02.22","pageSize":50,"singlePage":false,"level1Query":{"service_id":"router-0fc0b168-e7b5-4b20-98c6-e749338fa6ac","request_time":">1660895276743,<1663573676743"},"productName":"","serialNumber":"","dataChannelId":345},"extInfo":{"cmd":"python /usr/src/recheck_analysis/dist_calculate.py","total":3}},"resource":{"custom":{"id":"http://apulis-iqi.apulis:80/api/v1/inner/analysis/runs/update","context":{"labId":1016277,"owner":"","userInfo":{"orgId":1,"userId":48441424,"groupId":146713505,"orgName":"apulis","userName":"hs-org","groupName":"orgadmin-user-group"},"ownerType":"","dataChannelId":345,"analysisResultFilePath":"20220919_16.02.22"}},"analysis":{"path":"pvc://aiplatform-app-data-pvc/apulis-iqi/analysis","type":"store","rpath":"/shared-files/apulis-iqi/analysis","access":1,"context":null}},"envs":{"ANALYSIS_PATH":"/shared-files/apulis-iqi/analysis/20220919_16.02.22"},"quota":{"arch":"","node":0,"limit":{"cpu":"500m","device":{"series":"","deviceNum":"","deviceType":"","computeType":""},"memory":"500Mi"},"quotaId":0,"request":{"cpu":"500m","device":{"series":"","deviceNum":"","deviceType":"","computeType":""},"memory":"500Mi"},"partition":""},"schedState":{"pods":[{"name":"apulis-iqi-analysis-0a9eb237-bf00-41a9-bb07-f8bff30e3bed-k8m9d","phase":"Running","podIP":"172.20.153.242","hostIP":"192.168.2.153","nodeName":"192.168.2.153"}]},"flags":65538,"changedUid":48441424,"changedBy":"hs-org","groupName":"orgadmin-user-group","app":"apulis-iqi","orgName":"apulis","parentData":null}`
	runConfig := AILabRunConfig{}

	json.Unmarshal([]byte(str), &runConfig)
	param := runConfig.Config["param"]
	var extInfo map[string]interface{}
	var total int
	if bytes, ok := runConfig.Config["extInfo"]; ok {
		err := bytes.Fetch(&extInfo)
		if err != nil {
			logging.Error(err).Send()
			panic(err)
		}
		if value, ok := extInfo["total"]; ok {
			//total = int(math.Ceil(value.(float64)))
			total64, _ := strconv.ParseInt(fmt.Sprintf("%v", value), 10, 64)
			total = int(total64)
			logging.Debug().Interface("total", total).Send()
		}
	}
	logging.Debug().Interface("param", param).Send()

}
func TestFirstMissingPositive(t *testing.T) {
	nums := []int{1, 1}

	// 0111 7
	// 1000 8
	// 1001 9
	// 1010   10
	// 1011 11
	// 1100 12
	fmt.Printf("%d", firstMissingPositive(nums))
}
func firstMissingPositive(nums []int) int {
	var tmp int
	for i, _ := range nums {
		for {
			if nums[i] <= 0 || nums[i] > len(nums) {
				break
			}
			if nums[i] == i+1 {
				break
			}
			j := nums[i] - 1
			if nums[j] == nums[i] {
				break
			}
			tmp = nums[i]
			nums[i] = nums[j]
			nums[j] = tmp
		}
	}
	for i, num := range nums {
		if num-1 != i {
			return i + 1
		}
	}
	return len(nums) + 1
}

func TestList(t *testing.T) {
	var s []int

	s = append(s, 1)
	fmt.Println(s)

}

func TestMp(t *testing.T) {
	mp := map[string]uint8{}
	deepMp(mp)
	fmt.Println(mp)
}
func deepMp(mp map[string]uint8) {
	mp["sdadsa"] = 1
	return
}
