package main

import (
	"fmt"
	"golang.org/x/net/html/charset"
	"io/ioutil"
	"strings"
	"testing"
)

func TestUtf8(t *testing.T) {
	fileName := "api - \ufffd\ufffd\ufffd\ufffd.md"
	decodedFileName, err := decodeFileName(fileName)
	if err != nil {
		panic(err)
	}

	fmt.Println(decodedFileName)
}

func decodeFileName(fileName string) (string, error) {
	// 使用 charset.DetermineEncoding 函数检测字符集
	encoding, _, _ := charset.DetermineEncoding([]byte(fileName), "")
	utf8Reader := encoding.NewDecoder().Reader(strings.NewReader(fileName))
	decodedFileName, err := ioutil.ReadAll(utf8Reader)
	if err != nil {
		return "", err
	}
	return string(decodedFileName), nil
}
