package main

import (
	"fmt"
	"image"
	"k8s.io/apimachinery/pkg/api/resource"
	"os"
	"strconv"
	"strings"
	"testing"
	"time"
)

func TestTime(t *testing.T) {
	fmt.Println(time.Now().UnixMilli() / 1000)
}

func TestImageDecode(t *testing.T) {
	path := "D:\\Img12.jpg"
	width, height, err := decodeImage2(path)
	if err != nil {
		panic(err)
	}

	fmt.Println(width)
	fmt.Println(height)
}

func decodeImage2(absolutePath string) (width, height int, err error) {
	file, _ := os.Open(absolutePath)
	defer file.Close()

	img, _, err := image.Decode(file)
	if err != nil {
		fmt.Println("err = ", err)
		return
	}

	b := img.Bounds()
	return b.Max.X, b.Max.Y, err
}

func TestConvertUnit(t *testing.T) {
	ss()
}
func ss() {
	cpu, _ := resource.ParseQuantity("2")
	mem, _ := resource.ParseQuantity("4Gi")

	// 打印字符串表示
	fmt.Printf("CPU: %s, Memory: %s", cpu.String(), mem.String())
	fmt.Println()
	// 转换为整数值
	cpuInt64 := cpu.MilliValue()
	fmt.Println(cpuInt64)
	memInt64 := mem.Value()
	fmt.Println(memInt64)
	// 转换为不同格式
	// 转换为不同格式
	memInMi := mem.ScaledValue(resource.Mega)
	memInKi := mem.ScaledValue(resource.Kilo)
	fmt.Printf("CPU: %v, Memory: %v", memInMi, memInKi)

}
func convertUnit(quantity string) (float64, error) {
	if quantity == "" {
		return 0, nil
	}

	quantityParts := strings.Split(quantity, " ")
	value, err := strconv.ParseFloat(quantityParts[0], 64)
	if err != nil {
		return 0, err
	}

	unit := quantityParts[1]

	switch unit {
	case "Ki", "k":
		value *= 1024
	case "Mi", "M":
		value *= 1024 * 1024
	case "Gi", "G":
		value *= 1024 * 1024 * 1024
	case "Ti", "T":
		value *= 1024 * 1024 * 1024 * 1024
	case "Pi", "P":
		value *= 1024 * 1024 * 1024 * 1024 * 1024
	case "Ei", "E":
		value *= 1024 * 1024 * 1024 * 1024 * 1024 * 1024
	case "n":
		value /= 1000000000
	case "u":
		value /= 1000000
	case "m":
		value /= 1000
	}

	return value, nil
}
