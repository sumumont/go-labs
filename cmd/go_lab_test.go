package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/go-labs/internal/configs"
	"strconv"
	"strings"
	"testing"
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

func TestInline(t *testing.T) {
	var key = "key"
	var value = "value"
	dataProject := configs.Project{
		Key:   key,
		Value: value,
	}
	dataJiraHttpReqField := &configs.JiraHttpReqField{
		Project:     dataProject,
		Summary:     "Summary",
		Description: "Description",
	}
	data, _ := json.Marshal(dataJiraHttpReqField)
	fmt.Println(string(data))
	var config = configs.JiraHttpReqField{}
	err := json.Unmarshal(data, &config)
	if err != nil {
		panic(err)
	}
	fmt.Println(config)
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
