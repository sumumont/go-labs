package main

import (
	"encoding/base64"
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"testing"
)

type Coordinate struct {
	X          int //X轴坐标
	Y          int //Y轴坐标
	DefectType int //缺陷类型
}

func GetVrsData(str string) ([]Coordinate, error) {
	var coords []Coordinate
	var lines []string

	bytes, err := base64.StdEncoding.DecodeString(str)
	if err != nil {
		panic(err)
	}
	str = string(bytes)
	bytes = nil
	lines = strings.Split(str, "\r\n")
	if len(lines) < 6 {
		return coords, errors.New("invalid vrs file")
	}

	headerLen, err1 := strconv.ParseInt(lines[1], 10, 64) //文件头行数
	defectLen, err2 := strconv.ParseInt(lines[4], 10, 64) // 缺点数量
	checkLen, err3 := strconv.ParseInt(lines[5], 10, 64)  //对位点数量
	if err1 != nil || err2 != nil || err3 != nil {
		log.Printf("prase file header lines error: %v, %v, %v", err1, err2, err3)
		return coords, errors.New("parse file error")
	}
	start := headerLen + checkLen
	lines = lines[start:]

	coords = make([]Coordinate, defectLen)
	for i := 0; i < int(defectLen); i++ {
		item := strings.Split(lines[i], ",")
		x, _ := strconv.Atoi(item[0])
		y, _ := strconv.Atoi(item[1])
		typ, _ := strconv.Atoi(item[2])
		coords[i] = Coordinate{X: x, Y: y, DefectType: typ}
	}

	return coords, nil
}

func TestGetVrsData(t *testing.T) {
	file, err := os.ReadFile("D:\\OneDrive\\OneDrive - 依瞳科技（深圳）有限公司\\桌面\\aoicar\\aoicar\\20220921\\6081\\A.vrs")
	if err != nil {
		panic(err)
	}
	str := base64.StdEncoding.EncodeToString(file)
	file = nil
	result, err := GetVrsData(str)
	fmt.Println(result, err)

}
