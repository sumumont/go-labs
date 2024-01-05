package main

import (
	"encoding/base64"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"regexp"
	"sort"
	"strconv"
	"strings"
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
	now := time.Now().Format("20060102150405")
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
	//regex := regexp.MustCompile("^[a-zA-Z0-9_\u4e00-\u9fa5]+$")
	regex := regexp.MustCompile("^[\u4e00-\u9fa5_a-zA-Z0-9-]+$")
	str := "--大河坎打-."
	var result = regex.MatchString(str)
	fmt.Println(result)

	uid := uuid.New().String()
	fmt.Println(uid)
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
