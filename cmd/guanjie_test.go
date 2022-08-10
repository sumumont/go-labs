package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"testing"
)

type AllInferResult struct {
	DontHadAiResult []string              `json:"dontHadAiResult"`
	HadAiResult     map[string][]AiResult `json:"hadAiResult"` //key 是sheetname_x_x_partName
}
type AiResult struct {
	AiImageName string   `json:"aiImageName"`
	Objects     []Object `json:"Objects"`
}
type Object struct {
	Box            string      `json:"box"`
	Classification interface{} `json:"classification"`
	Label          string      `json:"label"`
	Ocr            interface{} `json:"ocr"`
	Result         string      `json:"result"`
	Score          float64     `json:"score"`
	Segmentation   interface{} `json:"segmentation"`
	SubObjects     interface{} `json:"sub_objects"`
}
type Box struct {
	X           int    `json:"X"`
	Y           int    `json:"Y"`
	Angle       int    `json:"Angle"`
	Result      string `json:"result"`
	Width       int    `json:"Width"`
	Height      int    `json:"Height"`
	DefectType  string `json:"DefectType"`
	DetailLabel string `json:"DetailLabel"`
}

func TestReadJson(t *testing.T) {
	//path := filepath.Join("tmp","test08041754_box.json")
	path := "D:\\code_private\\go-labs\\tmp\\test08041754_box.json"
	bytes, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}
	allInferResult := AllInferResult{}
	_ = json.Unmarshal(bytes, &allInferResult)
	fmt.Println(len(allInferResult.DontHadAiResult))
	fmt.Println(len(allInferResult.HadAiResult))
	fmt.Println("====================")
}
