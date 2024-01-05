package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"strings"
	"testing"
)

func TestInfer(t *testing.T) {
	//fileName := "000000000009.jpg"
	dir := "D:\\OneDrive\\OneDrive - 依瞳科技（深圳）有限公司\\桌面\\模型适配\\长鑫存储\\数据集\\coco128\\images\\train2017"
	err := InferDir(dir)
	if err != nil {
		panic(err)
	}
}

func TestInferImage(t *testing.T) {
	fileName := "000000000009.jpg"
	dir := "D:\\OneDrive\\OneDrive - 依瞳科技（深圳）有限公司\\桌面\\模型适配\\长鑫存储\\数据集\\coco128\\images\\train2017"
	err := InferImage(dir, fileName)
	if err != nil {
		panic(err)
	}
}

func SplitBy(key string, split string) (string, string) {
	idx := strings.LastIndex(key, split)
	if idx == -1 {
		return "", key
	}
	key1 := key[:idx]
	if key1 == "." {
		key1 = ""
	}
	key2 := key[idx+1:]
	return key1, key2
}
func TestSplitBy(t *testing.T) {
	a, b := SplitBy("./b.json", "/")
	fmt.Println(a, b)
}

func md5Encrypt(inputString string) string {
	hash := md5.Sum([]byte(inputString))
	encryptedString := hex.EncodeToString(hash[:])
	return encryptedString
}
func TestMd5(t *testing.T) {
	fmt.Println(md5Encrypt("apulis@123"))
}
