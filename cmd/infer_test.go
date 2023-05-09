package main

import (
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
