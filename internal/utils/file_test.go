package utils

import (
	"fmt"
	"regexp"
	"strconv"
	"testing"
	"time"
)

func TestFileObject(t *testing.T) {
	child := FileObject{
		Subpath: "go-labs",
		Name:    "go-labs",
		IsDir:   true,
		Child:   nil,
	}
	err := child.ReadChild("/root/apulis/tmp")
	if err != nil {
		panic(err)
	}
	printFile(child)
}
func printFile(fileObj FileObject) {
	fmt.Println(fileObj.Subpath, " ", fileObj.Name, " ", fileObj.IsDir)
	if len(fileObj.Child) != 0 {
		for _, file := range fileObj.Child {
			printFile(file)
		}
	}
}
func TestSetValue(t *testing.T) {
	cell := 'A'
	cell, _ = setCellValue(cell, 1, "dsad")
	cell, _ = setCellValue(cell, 1, "dsad")
	cell, _ = setCellValue(cell, 1, "dsad")
	cell, _ = setCellValue(cell, 1, "dsad")
	cell, _ = setCellValue(cell, 1, "dsad")
}
func setCellValue(cell int32, line int, value interface{}) (int32, error) {
	fmt.Println(fmt.Sprintf("%c%v", cell, line))
	cell = cell + 1

	return cell, nil
}
func TestImg(t *testing.T) {
	imageName := "LWA51XXB45965788P0725_LWA51XXB45965788P0726_Img11321.jpg"
	_, lastName := SplitKey(imageName, "_")
	fmt.Println(lastName)
	var valid = regexp.MustCompile("[0-9]")
	value := valid.FindAllStringSubmatch(lastName, -1)
	resultStr := ""
	for _, x := range value {
		for _, y := range x {
			resultStr = resultStr + y
		}
	}
	var result int64
	if resultStr != "" {
		result, _ = strconv.ParseInt(resultStr, 10, 64)
	}
	fmt.Println(result)
}

func TestTime(t *testing.T) {
	sheetTime := "8/11/22 18:21"
	newTime, err := time.Parse("1/2/06 15:04", sheetTime)
	if err != nil {
		panic(err)
	}
	sheetTime = newTime.Format("20060102150405")
	fmt.Println(sheetTime)
}
