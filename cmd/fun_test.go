package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/go-labs/internal/logging"
	"image"
	"os"
	"path"
	"path/filepath"
	"strconv"
	"strings"
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
func TestDecodePic(t *testing.T) {
	// label_mask
	pics := `iVBORw0KGgoAAAANSUhEUgAAAQAAAADACAAAAADOhuK6AAABoElEQVR4AeXBgW3DQBDAMGn/oa+Jvyu8cYFIiZM4iZM4iZM4iZM4iZM4iZM4iZM4iZM4iZM4iZM4iZM4iZM4iZM4iZM4iZM4iZM4iZM4iZM4iZM4iZM4iZM4iZM4iZM4iZM4uUOGnyBXyMfwA+QGeQz7yQVyDPvJBXIM+8kFcgz7yQVyDPvJDfIY9pMr5GP4ARIncRIncRIncRIncRIncRIncRInmwgMr5JF5Gt4k+whx/Ai2UOO4UWyhxzDi2QPOYYXyR5yDC+SReRreJNsIjC8SuIkTuIkTuIkTuIkTuIkTuIkTuIkTuIkTuIkTuLkHmHYTq6Rr2E3uUUew25yixzDanKLHMNqcon8G1aTW+QYVpNb5DHsJtfIx7Cc3CMM20mcxEmcxEmcxEmcxEmcxEmcxEmcxEmcxEmcxEmcxEmcxEmcxEmcxEmcxEmcxEmcxEmcxEmcxEmcxEmcxEmcxEmcxEmcxEmcxEmcxEmcxEmcxEmcxEmcxEmcxEmcxEmcxEmcxEmcxEmcxEmcxEmcxEmcxEmcxEmcxEmcxEmcxP0BhKscwa3upNMAAAAASUVORK5CYII=`
	//blob
	//pics := `iVBORw0KGgoAAAANSUhEUgAAAQAAAADACAAAAADOhuK6AAACGElEQVR4AeXByXGEMAAAwZkAyIwfcfIjMxKwfK0TWFFyTbfESZzESZzESZzESZzESZzESZzESZzESZzESZzESZzESZzESZzESZzESZzESZzESZzESZzESZzESZzESZzESZzESZzESZzMcbHzL8gUF8POPyAzXHzZWZ9McPFjZ3nyfhcvO6uTt7sYdi6GndXJ210MOxfDzurk7S6GnYthZ3Xydhd/dlYn73fxsrM6meDix87yZIaLLzvrkykuhp1/QOIkTuIkTuIkTuIkTuJkJTew8ShZyM2njSfJOm6+bTxIlnHza+M5soqbPxuPkVXcDBs3w8ZjZBU3w8bNsPEYWcXNsHEzbDxGVnHzZ+MxsoybXxvPkXXcfNt4kCzk5tPGk2QlN7DxKImTOImTOImTOImTOImTOImTOImTOImTeU44WJ1Mc/LpYG0yy8m3g6XJJCe/DlYmc5y8HKxM5jgZDk6Gg4XJHCfDwclwsDCZ4mQ44GQ4WJjMcTIcnAwHC5M5Tl4OViaTnPw4WJrMcvLlYG0yzclwsDiZ54SD1UmcxEmcxEmcxEmcxEmcxEmcxEmcxEmcxEmcxEmcxEmcxEmcxEmcxEmcxEmcxEmcxEmcxEmcxEmcxEmcxEmcxEmcxEmcxEmcxEmcxEmcxEmcxEmcxEmcxEmcxEmcxEmcxEmcxEmcxEmcxEmcxEmcxEmcxEmcxEmcxEmcxH0AWJAyweJU8FcAAAAASUVORK5CYII=`
	bytes, err := base64.StdEncoding.DecodeString(pics)
	if err != nil {
		panic(err)
	}
	err = os.WriteFile("hysen_pic_label_mask.png", bytes, os.ModePerm)

	if err != nil {
		panic(err)
	}

	//rb := new(io.Buffer)
	reader := strings.NewReader(string(bytes))

	canvas, str, err := image.Decode(reader)
	if err != nil {
		panic(err)
	}
	fmt.Println(str)
	fmt.Println(canvas.ColorModel())
	bounds := canvas.Bounds()
	fmt.Println(bounds.Min, bounds.Max)
	switch canvas.(type) {
	case *image.NRGBA:
		//img := canvas.(*image.NRGBA)
		fmt.Println("NRGBA")
	case *image.RGBA:
		//img := canvas.(*image.RGBA)
		fmt.Println("NRGBA")
	case *image.Gray:
		fmt.Println("Gray")
		img := canvas.(*image.Gray)

		var points [][]uint8
		var line []uint8
		first := -1
		for i, pix := range img.Pix {
			if pix != 0 {
				if first == -1 {
					first = i
				}
			}
			//fmt.Print(pix)

			//x := (i + 1) / img.Stride
			//y := (i+1)%img.Stride - 1
			line = append(line, pix)
			if (i+1)%img.Stride == 0 {
				//fmt.Println()
				points = append(points, line)
				line = nil
			}
		}
		fmt.Println("==========================================")
		fmt.Println("first", first)
		allMap := map[string]point{}
		var mpPoints [][]point
		allPics := []map[string]point{}
		for y := 0; y < len(points); y++ {
			for x := 0; x < len(points[y]); x++ {
				if points[y][x] == 0 {
					continue
				}
				start := point{
					X: x,
					Y: y,
				}
				kk := start.getK()
				if _, ok := allMap[kk]; ok {
					continue
				}
				mp := map[string]point{}
				dfs(start, points, mp, img.Bounds().Dx(), img.Bounds().Dy())
				if len(mp) > 0 {
					for k, v := range mp {
						allMap[k] = v
					}
					allPics = append(allPics, mp)
				}
				//fmt.Println("========================start", y, x)
			}

			//fmt.Println()
		}
		fmt.Println("pics.len", len(mpPoints))
		{
			maxX := -1
			maxY := -1
			for _, v := range allMap {
				if maxX < v.X {
					maxX = v.X
				}
				if maxY < v.Y {
					maxY = v.Y
				}
			}

			for y := 0; y <= maxY; y++ {
				for x := 0; x <= maxX; x++ {
					p := point{
						X: x,
						Y: y,
					}
					k := p.getK()
					if _, ok := allMap[k]; ok {
						fmt.Print(1)
					} else {
						fmt.Print(0)
					}
				}
				fmt.Println()
			}
		}
		//todo 深度搜索找出属于同一张图片的点位

		//for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		//	for x := bounds.Min.X; x < bounds.Max.X; x++ {
		//		location := (y-bounds.Min.Y)*img.Stride + (x-bounds.Min.X)*1
		//		fmt.Print(img.Pix[location])
		//	}
		//	fmt.Println("")
		//}
	}

}

var left, right, top, boot = -1, 1, -1, 1

type point struct {
	X int
	Y int
}

func (p point) getK() string {
	return fmt.Sprintf("%d,%d", p.Y, p.X)
}

func (p point) Top() *point {
	return &point{
		X: p.X,
		Y: p.Y - 1,
	}
}

func (p point) TopLeft() *point {
	return &point{
		X: p.X - 1,
		Y: p.Y - 1,
	}
}
func (p point) Left() *point {
	return &point{
		X: p.X - 1,
		Y: p.Y,
	}
}

func (p point) LeftBoot() *point {
	return &point{
		X: p.X - 1,
		Y: p.Y + 1,
	}
}
func (p point) Boot() *point {
	return &point{
		X: p.X,
		Y: p.Y + 1,
	}
}

func (p point) BootRight() *point {
	return &point{
		X: p.X + 1,
		Y: p.Y + 1,
	}
}
func (p point) Right() *point {
	return &point{
		X: p.X + 1,
		Y: p.Y,
	}
}

func (p point) RightTop() *point {
	return &point{
		X: p.X + 1,
		Y: p.Y + 1,
	}
}
func (p point) Out(xMax, yMax int) bool {
	if p.X < 0 || p.X >= xMax {
		return true
	}
	if p.Y < 0 || p.Y >= yMax {
		return true
	}
	return false
}

// mp 表示一个图的集合 key= "x,y"   value=像素值
// pics 也表示一个图的集合
func dfs(point point, points [][]uint8, mp map[string]point, xMax, yMax int) {
	//节点越界
	if point.Out(xMax, yMax) {
		return
	}
	k := point.getK()
	x := point.X
	y := point.Y
	if points[y][x] == 0 {
		return
	}

	if _, ok := mp[k]; !ok {
		mp[k] = point
		Top := *point.Top()
		dfs(Top, points, mp, xMax, yMax)
		TopLeft := *point.TopLeft()
		dfs(TopLeft, points, mp, xMax, yMax)
		Left := *point.Left()
		dfs(Left, points, mp, xMax, yMax)
		LeftBoot := *point.LeftBoot()
		dfs(LeftBoot, points, mp, xMax, yMax)
		Boot := *point.Boot()
		dfs(Boot, points, mp, xMax, yMax)
		BootRight := *point.BootRight()
		dfs(BootRight, points, mp, xMax, yMax)
		Right := *point.Right()
		dfs(Right, points, mp, xMax, yMax)
		RightTop := *point.RightTop()
		dfs(RightTop, points, mp, xMax, yMax)
	}
	return
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
