package main

import (
	"encoding/base64"
	"errors"
	"fmt"
	"github.com/fatih/color"
	"github.com/go-labs/internal/http_client"
	"github.com/go-labs/internal/logging"
	"github.com/gosuri/uiprogress"
	"github.com/robfig/cron/v3"
	"github.com/vbauerster/mpb/v8"
	"github.com/vbauerster/mpb/v8/decor"
	"math/rand"
	"net/http"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
	"sync"
	"testing"
	"time"
)

func TestPageUtil(t *testing.T) {

	files := []FileListItem{}
	for i := 0; i < 31; i++ {
		k := FileListItem{
			Name:      "name" + strconv.Itoa(i),
			UpdatedAt: 0,
			Size:      0,
			IsDir:     false,
		}
		files = append(files, k)
		fmt.Println(k)
	}

	cond := &SearchCond{
		Offset:     0,
		TotalCount: int64(len(files)),
		PageNum:    4,
		PageSize:   10,
	}
	if cond.PageNum >= 1 {
		cond.Offset = (cond.PageNum - 1) * cond.PageSize
	}
	result := PageUtil(cond, files)
	for idx, k := range result {
		v := k.(FileListItem)
		fmt.Println(idx, v.Name)
	}
}

func PageUtil(con *SearchCond, list ...interface{}) []interface{} {
	start := con.Offset
	end := start + con.PageSize
	if end > uint(con.TotalCount) {
		end = uint(con.TotalCount)
	}
	if start > end {
		return nil
	}
	return list[0:1]
}

type FileListItem struct {
	Name string `json:"name"`
	//CreatedAt int64     `json:"createdAt"`
	UpdatedAt int64 `json:"createdAt"`
	Size      int64 `json:"size"`
	IsDir     bool  `json:"isDir"`
}
type SearchCond struct {
	Offset     uint
	TotalCount int64
	Next       string
	//start from 1~N
	PageNum      uint   `form:"pageNum"`
	PageSize     uint   `form:"pageSize"`
	Sort         string `form:"sort"`
	UseModelArts bool   `form:"useModelArts"`
	// list by app group
	Group string `form:"group"`
	// indicate "group" list match recursively !
	MatchAll bool `form:"matchAll"`
	// search by keyword
	SearchWord string `form:"searchWord"`
	//enumeration for need detail return
	Detail int32 `form:"detail"`
	//enumeration for deleted item search
	Show int32 `form:"show"`
	// filters by predefined key=value pairs
	EqualFilters map[string]string
	// filters by advacned operator
	AdvanceOpFilters map[string]interface{}
}

func TestString(t *testing.T) {
	filePath := "/app/ai-labs-data/default/iqi/334/eval-4217b01d-4796-49e4-a82b-a76f8bde36af/ok_images"
	imageProxyPath := strings.TrimPrefix(filePath, "/app")
	fmt.Println(imageProxyPath)
}

func TestDefer(t *testing.T) {
	getString(true)
}

func getString(ok bool) error {
	fmt.Println(ok)
	if !ok {
		return errors.New("dsadsa")
	}
	defer fmt.Print("2132121")

	for i := 0; i < 10; i++ {
		fmt.Println(i)
		defer fmt.Println(i, "defer")
	}
	fmt.Println(21321312)
	return nil
}

func TestUnix(t *testing.T) {
	now1 := time.Now().Unix()
	fmt.Println(now1)
	time.Sleep(time.Second * 5)
	now2 := time.Now().Unix()
	fmt.Println(now2)
	fmt.Println(now2 - now1)
}

type Metadata struct {
	Result *bool `json:"result"`
}

func TestTimeFormat(t *testing.T) {
	now := time.Now().Format("20060102150405001")
	fmt.Println("now:", now)
}

type ProjectModelTemplate struct {
	Field string `json:"field"`
	Task  string `json:"task"`
}

type ProjectModelTemplates []ProjectModelTemplate

func (p ProjectModelTemplates) Len() int { return len(p) }
func (p ProjectModelTemplates) Less(i, j int) bool {
	if p[i].Field < p[j].Field {
		return true
	}
	if p[i].Field > p[j].Field {
		return false
	}
	return p[i].Task < p[j].Task
}
func (p ProjectModelTemplates) Swap(i, j int) { p[i], p[j] = p[j], p[i] }

func TestSort(t *testing.T) {
	var req = ProjectModelTemplates{
		ProjectModelTemplate{
			Field: "a",
			Task:  "a1",
		},
		ProjectModelTemplate{
			Field: "a",
			Task:  "a2",
		},
		ProjectModelTemplate{
			Field: "b",
			Task:  "b2",
		},
		ProjectModelTemplate{
			Field: "b",
			Task:  "b1",
		},
		ProjectModelTemplate{
			Field: "c",
			Task:  "c2",
		},
		ProjectModelTemplate{
			Field: "a",
			Task:  "c1",
		},
	}
	sort.Sort(req)
	for _, v := range req {
		fmt.Println(v)
	}
}

func TestReg(t *testing.T) {
	tmp := "YXB1bGlzLmNvbQ=="
	vstr, _ := base64.StdEncoding.DecodeString(tmp)
	fmt.Println(string(vstr))

	fmt.Println(parseId(tmp + "MQ=="))
}
func parseId(idString string) (int64, error) {
	prefix := base64.StdEncoding.EncodeToString([]byte("apulis.com"))
	if !strings.HasPrefix(idString, prefix) {
		return 0, errors.New("id format not right")
	}
	idStr := strings.TrimPrefix(idString, prefix)
	idBytes, err := base64.StdEncoding.DecodeString(idStr)
	if err != nil {
		return 0, err
	}
	id, err := strconv.ParseInt(string(idBytes), 10, 64)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func TestBase64(t *testing.T) {
	uploadmeta := `location aXFp,moduleName YXB1bGlzLWlxaQ==,objectPrefix dW5kZWZpbmVk,taskType c2Rr,relativePath bnVsbA==,name c2VydmUxLnppcA==,type YXBwbGljYXRpb24veC16aXAtY29tcHJlc3NlZA==,filetype YXBwbGljYXRpb24veC16aXAtY29tcHJlc3NlZA==,filename c2VydmUxLnppcA==`
	fmt.Println(uploadmeta)
	ups := strings.Split(uploadmeta, ",")
	kv := []Kv{}
	for _, up := range ups {
		tmp := strings.Split(up, " ")
		vstr, _ := base64.StdEncoding.DecodeString(tmp[1])
		k := tmp[0]
		v := string(vstr)
		fmt.Println(k, v)
		kv = append(kv, Kv{
			K: k,
			V: v,
		})
	}
	res := printKv(kv)
	fmt.Println(res)

	ds := []Kv{
		{
			K: "location", V: "iqi",
		}, {
			K: "moduleName", V: "apulis-iqi",
		}, {
			K: "objectPrefix", V: "undefined",
		}, {
			K: "taskType", V: "sdk",
		}, {
			K: "relativePath", V: "null",
		}, {
			K: "name", V: "serve2.zip",
		}, {
			K: "type", V: "application/x-zip-compressed",
		}, {
			K: "filetype", V: "application/x-zip-compressed",
		}, {
			K: "filename", V: "serve2.zip",
		},
	}
	res1 := printKv(ds)
	fmt.Println(res1)
	fmt.Println(res == res1)
}
func printKv(kv []Kv) string {
	str := strings.Builder{}
	for idx, up := range kv {
		v := up.V
		v = base64.StdEncoding.EncodeToString([]byte(v))
		if idx == len(kv)-1 {
			str.WriteString(up.K + " " + v)
		} else {
			str.WriteString(up.K + " " + v + ",")
		}
	}
	return str.String()
}

type Kv struct {
	K string
	V string

	Key   string
	Value string
}

func TestReflect(t *testing.T) {

	//"upload.url":"https://internal-proxy.default.svc.cluster.local/file-servers/apulis/file-server/api/v1/files/"}
	domain1 := strings.Split("https://internal-proxy.default.svc.cluster.local/file-servers/apulis/file-server/api/v1/files/", "/file-server/")
	domains := strings.Split(domain1[0], "://")
	domain := domains[1]
	proto := domains[0]
	fmt.Println(domain, proto)

}

func TestArray1(t *testing.T) {
	endStatus := 357 & 0xFF
	fmt.Println(endStatus)
}

func Append1(a []string) {
	s := []string{"153", "234432", "13123"}
	a = append(a, s...)
	fmt.Println(a)
}

func TestBar(t *testing.T) {
	waitTime := time.Millisecond * 100
	uiprogress.Start()

	// start the progress bars in go routines
	var wg sync.WaitGroup

	bar1 := uiprogress.AddBar(20).AppendCompleted().PrependElapsed()
	wg.Add(1)
	go func() {
		defer wg.Done()
		for bar1.Incr() {
			time.Sleep(waitTime)
		}
	}()

	bar2 := uiprogress.AddBar(40).AppendCompleted().PrependElapsed()
	wg.Add(1)
	go func() {
		defer wg.Done()
		for bar2.Incr() {
			time.Sleep(waitTime)
		}
	}()

	time.Sleep(time.Second)
	bar3 := uiprogress.AddBar(20).PrependElapsed().AppendCompleted()
	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 1; i <= bar3.Total; i++ {
			bar3.Set(i)
			time.Sleep(waitTime)
		}
	}()

	// wait for all the go routines to finish
	wg.Wait()
}
func TestBar4(t *testing.T) {
	fmt.Println("zzzzzz")
	for i := 0; i < 100; i++ {
		fmt.Printf("\r%d", i)
		time.Sleep(500 * time.Millisecond)
	}
}

func TestBar3(t *testing.T) {
	p := mpb.New(
		mpb.WithOutput(color.Output),
		mpb.WithAutoRefresh(),
	)
	red, green := color.New(color.FgRed), color.New(color.FgGreen)
	task := fmt.Sprintf("Task#%02d:", 1)
	queue := make([]*mpb.Bar, 2)
	queue[0] = p.AddBar(rand.Int63n(201)+100,
		mpb.PrependDecorators(
			decor.Name(task, decor.WC{C: decor.DindentRight | decor.DextraSpace}),
			decor.Name("downloading", decor.WCSyncSpaceR),
			decor.CountersNoUnit("%d / %d", decor.WCSyncWidth),
		),
		mpb.AppendDecorators(
			decor.OnComplete(decor.Percentage(decor.WC{W: 5}), "done"),
		),
	)
	queue[1] = p.AddBar(rand.Int63n(101)+100,
		mpb.BarQueueAfter(queue[0]), // this bar is queued
		mpb.BarFillerClearOnComplete(),
		mpb.PrependDecorators(
			decor.Name(task, decor.WC{C: decor.DindentRight | decor.DextraSpace}),
			decor.OnCompleteMeta(
				decor.OnComplete(
					decor.Meta(decor.Name("installing", decor.WCSyncSpaceR), toMetaFunc(red)),
					"done!",
				),
				toMetaFunc(green),
			),
			decor.OnComplete(decor.EwmaETA(decor.ET_STYLE_MMSS, 0, decor.WCSyncWidth), ""),
		),
		mpb.AppendDecorators(
			decor.OnComplete(decor.Percentage(decor.WC{W: 5}), ""),
		),
	)

	go func() {
		for _, b := range queue {
			complete(b)
		}
	}()
	p.Wait()
}
func TestBar2(t *testing.T) {
	numBars := 1
	// to support color in Windows following both options are required
	p := mpb.New(
		mpb.WithOutput(color.Output),
		mpb.WithAutoRefresh(),
	)

	red, green := color.New(color.FgRed), color.New(color.FgGreen)

	for i := 0; i < numBars; i++ {
		task := fmt.Sprintf("Task#%02d:", i)
		queue := make([]*mpb.Bar, 2)
		queue[0] = p.AddBar(rand.Int63n(201)+100,
			mpb.PrependDecorators(
				decor.Name(task, decor.WC{C: decor.DindentRight | decor.DextraSpace}),
				decor.Name("downloading", decor.WCSyncSpaceR),
				decor.CountersNoUnit("%d / %d", decor.WCSyncWidth),
			),
			mpb.AppendDecorators(
				decor.OnComplete(decor.Percentage(decor.WC{W: 5}), "done"),
			),
		)
		queue[1] = p.AddBar(rand.Int63n(101)+100,
			mpb.BarQueueAfter(queue[0]), // this bar is queued
			mpb.BarFillerClearOnComplete(),
			mpb.PrependDecorators(
				decor.Name(task, decor.WC{C: decor.DindentRight | decor.DextraSpace}),
				decor.OnCompleteMeta(
					decor.OnComplete(
						decor.Meta(decor.Name("installing", decor.WCSyncSpaceR), toMetaFunc(red)),
						"done!",
					),
					toMetaFunc(green),
				),
				decor.OnComplete(decor.EwmaETA(decor.ET_STYLE_MMSS, 0, decor.WCSyncWidth), ""),
			),
			mpb.AppendDecorators(
				decor.OnComplete(decor.Percentage(decor.WC{W: 5}), ""),
			),
		)

		go func() {
			for _, b := range queue {
				complete(b)
			}
		}()
	}

	p.Wait()
}
func complete(bar *mpb.Bar) {
	max := 100 * time.Millisecond
	for !bar.Completed() {
		// start variable is solely for EWMA calculation
		// EWMA's unit of measure is an iteration's duration
		start := time.Now()
		time.Sleep(time.Duration(rand.Intn(10)+1) * max / 10)
		// we need to call EwmaIncrement to fulfill ewma decorator's contract
		bar.EwmaIncrInt64(rand.Int63n(5)+1, time.Since(start))
	}
}

func toMetaFunc(c *color.Color) func(string) string {
	return func(s string) string {
		return c.Sprint(s)
	}
}

type A struct {
	Name string
}

func (a A) Say() {
	fmt.Printf("hello")
}

type Out1 interface {
	Say()
}

type B struct {
}

func (b B) Say() {
	fmt.Printf("hello2")
}

func SendRequest[Res Out1](a Out1) error {
	a.Say()
	return nil
}

func TestAny(t *testing.T) {
	a := A{}
	_ = SendRequest[Out1](a)
}

func TestHeader(t *testing.T) {
	headers := http.Header{}
	headers.Set("Authorization", "Bearer "+"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MTkwNDQ1NzQsImdyb3VwX2FjY291bnQiOiJvcmdhZG1pbi11c2VyLWdyb3VwIiwiZ3JvdXBfaWQiOjczNDE5NjUwLCJvcmdhbml6YXRpb25fYWNjb3VudCI6ImFwdWxpcyIsIm9yZ2FuaXphdGlvbl9pZCI6MSwib3JnYW5pemF0aW9uX3N0YXR1cyI6ImluVXNlIiwib3JpZ19pYXQiOjE3MTY0NTI1NzQsInJvbGUiOjIsInVzYWdlIjoiIiwidXNlcl9pZCI6MTI1NDg2MzUxLCJ1c2VyX25hbWUiOiJoeXNlbi1vcmcifQ.-8vmQQj57k8nZqSaVW8hZOu4x5gX0ju2p2T1U1-dLeM")
	for i := 0; i < 1000; i++ {
		headers.Add("clusterId", "0")
		headers.Set("Content-Type", "application/json")

	}
	rsp, err := http_client.DoRequest("http://127.0.0.1:8001/test/123", "GET", headers, nil)
	//_, err = AIStudioRequestWithData(apHarborUrl.String(), "GET", headers, nil, data3)
	if err != nil {
		fmt.Printf("err: %v", err.Error())
	}
	if rsp.Body != nil {
		fmt.Print(rsp.Body)
	}
}
func TestAny1(t *testing.T) {
	c := cron.New()
	c.AddFunc("* * * * * ?", func() {
		logging.Debug().Msg("hello 1")
	})
	c.Start()
	signals(c)
}
func TestAny2(t *testing.T) {
	fmt.Println(filepath.Dir("/a/b/"))
}
