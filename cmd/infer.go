package main

import (
	"fmt"
	"github.com/go-labs/internal/logging"
	"github.com/go-labs/internal/utils"
	"net/http"
	"os"
	"path/filepath"
)

type InferData struct {
	RequestId     string            `json:"request_id"`
	Tags          map[string]string `json:"tags"`
	DataFormat    string            `json:"data_format"`
	DataEncoding  string            `json:"data_encoding"`
	RequestParams map[string]string `json:"request_params"`
	Requests      []RequestData     `json:"requests"`
}
type RequestData struct {
	Name string `json:"name"`
	Data string `json:"data"`
}

func InferDir(dirPath string) error {
	files, err := os.ReadDir(dirPath)
	if err != nil {
		logging.Error(err).Send()
		return err
	}
	for _, file := range files {
		if file.IsDir() {
			path := filepath.Join(dirPath, file.Name())
			err = InferDir(path)
			if err != nil {
				logging.Error(err).Send()
				return err
			}
		} else {
			err = InferImage(dirPath, file.Name())
			if err != nil {
				logging.Error(err).Send()
				return err
			}
		}
	}
	return nil
}
func InferImage(dirPath string, fileName string) error {
	//fileName := "000000000009.jpg"
	//src := "D:\\OneDrive\\OneDrive - 依瞳科技（深圳）有限公司\\桌面\\模型适配\\长鑫存储\\数据集\\coco128\\images\\train2017\\" + fileName
	src := filepath.Join(dirPath, fileName)
	imageBase := ReadImageBase64(src)
	url := "https://192.168.2.75:443/inference/router-r7q2pwnanqbbusj5g4bpvn/api/v1/scenes/20/infer"
	header := map[string]string{
		//"Content-Type": 'application/json',
		"Traceid": "88888888",
		//"X-Apulis-Infer-Result-Store": "N",
	}
	now := utils.GetNowTime()
	param := InferData{
		RequestId: fmt.Sprintf("%v", now),
		Tags: map[string]string{
			"productSN": "cxcc-p1",
		},
		DataFormat:    "image/jpg",
		DataEncoding:  "base64",
		RequestParams: map[string]string{},
		Requests: []RequestData{
			{
				Name: fileName,
				Data: imageBase,
			},
		},
	}

	var result map[string]interface{}
	err := utils.DoRequest(url, http.MethodPost, header, param, &result)
	if err != nil {
		logging.Error(err).Send()
		return err
	}
	logging.Debug().Str("imageName", fileName).Interface("result", result).Send()
	return nil
}
