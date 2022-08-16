package main

import (
	"fmt"
	"path"
	"path/filepath"
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

func TestBox(t *testing.T) {

}
