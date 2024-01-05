package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/disintegration/imaging"
	"github.com/go-labs/internal/logging"
	"image"
	"os"
	"path"
	"reflect"
	"strconv"
	"strings"
	"testing"
)

func TestSample(t *testing.T) {
	// 打开原始图片文件
	file, err := os.Open("D:\\OneDrive\\OneDrive - 依瞳科技（深圳）有限公司\\桌面\\新建文件夹 (2)\\20230320180437\\j2数据\\test.png")
	if err != nil {
		fmt.Println("无法打开图片文件:", err)
		return
	}
	defer file.Close()

	// 解码图片文件
	img, fm, err := image.Decode(file)
	if err != nil {
		fmt.Println("无法解码图片:", err)
		return
	}
	fmt.Println("format", fm)
	rate := 0.5
	// 设置目标宽度和高度
	targetWidth := float64(img.Bounds().Dx()) * rate
	targetHeight := float64(img.Bounds().Dy()) * rate

	// 调用Resize函数进行降采样
	resized := imaging.Resize(img, int(targetWidth), int(targetHeight), imaging.Lanczos)
	//imaging.Sharpen(img, targetWidth, targetHeight, imaging.Lanczos)
	// 创建输出文件
	outFile, err := os.Create("output.png")
	if err != nil {
		fmt.Println("无法创建输出文件:", err)
		return
	}
	defer outFile.Close()

	// 将降采样后的图片保存到输出文件
	format, err := getFormat(fm)
	if err != nil {
		fmt.Println(err)
		return
	}
	err = imaging.Encode(outFile, resized, format)
	if err != nil {
		fmt.Println("无法保存图片:", err)
		return
	}

	fmt.Println("图片降采样完成，已保存为output.jpg")
}

func getFormat(fm string) (format imaging.Format, err error) {
	switch fm {
	case "jpg":
		format = imaging.JPEG
	case "jpeg":
		format = imaging.JPEG
	case "png":
		format = imaging.PNG
	case "gif":
		format = imaging.GIF
	case "bmp":
		format = imaging.BMP
	default:
		err = errors.New("ERROR FORMAT")
	}
	return
}

func TestParseFilenameTimestamp(t *testing.T) {
	res, err := parseFilenameTimestamp("1560159026665.jpg")
	fmt.Println(res, err)
}

func parseFilenameTimestamp(filePath string) (res uint64, err error) {
	base := path.Base(filePath)
	suffix := path.Ext(base)
	nameOnly := filePath[:len(filePath)-len(suffix)]
	res, err = strconv.ParseUint(nameOnly, 10, 64)
	if err != nil {
		return
	}
	return
}

func TestParse1(t *testing.T) {
	inferResults := []map[string]interface{}{}
	path1 := "D:\\OneDrive\\OneDrive - 依瞳科技（深圳）有限公司\\桌面\\新建文件夹 (2)\\20230320180438\\J2camera\\20190610173026_1560159026665.json"
	imagePath := path.Join(path1)
	imageSuffix := path.Ext(imagePath)
	jsonPath := strings.ReplaceAll(imagePath, imageSuffix, ".json")
	err := ReadFromJson(jsonPath, &inferResults)
	if err != nil {
		panic(err)
	}
	//var supportTypes = []string{"freespace", "obs_raw"}
	//for _, item := range inferResults {
	//	for _, supportType := range supportTypes {
	//		if infer, ok := item[supportType]; ok {
	//			prepareAndWriteData(supportType, infer)
	//		}
	//	}
	//}
	var supportTypes = []string{"j2camera"}
	for _, supportType := range supportTypes {
		prepareAndWriteData(supportType, &inferResults)
	}
}
func prepareAndWriteData(supportType string, infer interface{}) {
	tp := reflect.TypeOf(infer)
	fmt.Println(tp, tp.Kind())

	if tp.Kind() == reflect.Ptr {
		tp = tp.Elem()
	}
	switch tp.Kind() {
	case reflect.Slice, reflect.Array:
		fmt.Println("array", supportType)
	default:

		fmt.Println("single", supportType)
	}

	//switch infer.(type) {
	//case []interface{}, []map[string]interface{}:
	//
	//case interface{}:
	//	//newInfer := []interface{}{infer}
	//}
}

// output 需要是一个地址
func ReadFromJson(filename string, output interface{}) error {
	bytes, err := os.ReadFile(filename)
	if err != nil {
		logging.Error(err).Send()
		return err
	}
	err = json.Unmarshal(bytes, output)
	if err != nil {
		logging.Error(err).Send()
		return nil
	}
	return nil
}

func errorgroup() {

}
