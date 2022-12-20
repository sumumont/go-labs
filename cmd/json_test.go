package main

import (
	"archive/tar"
	"compress/gzip"
	"encoding/json"
	"fmt"
	"github.com/go-labs/internal/logging"
	"github.com/go-labs/internal/utils"
	"io"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestJson(t *testing.T) {
	result := map[string]interface{}{}
	str := "{\"error\":{\"ErrorCode\":\"1\",\"ErrorMsg\":\"Invalid input of patterns or images files\",\"InferenceTime\":{\"TotalTime\":0.0008442401885986328},\"IsSuccess\":false,\"Objects\":[],\"Result\":\"NG\"}}\n"
	err := json.Unmarshal([]byte(str), &result)
	if err != nil {
		panic(err)
	}
	logging.Debug().Interface("result", result).Send()
}
func TestZip(t *testing.T) {
	//var zip = ".zip"
	//var tar = ".tar"
	//var targz = ".tar.gz"
	//var supportFileTpe = []string{zip, tar, targz}
	zipFileName := GetFileName("apulis-iqi/hysen_pic.zip")
	decompressName, _ := utils.SplitKey(zipFileName, ".")
	rootPath := filepath.Join("/tmp", decompressName)
	fmt.Println(rootPath)

	fmt.Println(strings.TrimSuffix("apulis-iqi/hysen_pic.zip", ".zip"))
}

func GetFileName(name string) string {
	_, filename := utils.SplitKey(name, "/")
	return filename
}
func RmSuffix(name string, suffix string) string {
	idx := strings.LastIndex(name, suffix)
	key1 := name[:idx]
	//key2 := name[idx+1:]
	return key1
}
func TestUnTargz(t *testing.T) {
	//fr, err := os.Open("C:\\Users\\haisen\\Downloads\\85300504-029f-45de-b94c-632cd6f9df40.tar.gz")
	//if err != nil {
	//	panic(err)
	//}
	//defer fr.Close()
	//
	//// gzip read
	//gr, err := gzip.NewReader(fr)
	//if err != nil {
	//	panic(err)
	//}
	//defer gr.Close()
	//
	//// tar read
	//tr := tar.NewReader(gr)
	//
	//// 读取文件
	//dest := "file2/"
	//_ = os.MkdirAll(dest, os.ModePerm)
	//for {
	//	h, err := tr.Next()
	//	if err == io.EOF {
	//		break
	//	}
	//	if err != nil {
	//		panic(err)
	//	}
	//
	//	// 显示文件
	//	fmt.Println(h.Name)
	//
	//	// 打开文件
	//	fw, err := os.OpenFile(dest+h.Name, os.O_CREATE|os.O_WRONLY, 0644 /*os.FileMode(h.Mode)*/)
	//	if err != nil {
	//		panic(err)
	//	}
	//	defer fw.Close()
	//
	//	// 写文件
	//	_, err = io.Copy(fw, tr)
	//	if err != nil {
	//		panic(err)
	//	}
	//
	//}
	err := UnTar("D:\\scenes.tar", "dsadsada")
	if err != nil {
		panic(err)
	}
}

func DeCompress(tarFile, dest string) error {
	srcFile, err := os.Open(tarFile)
	if err != nil {
		return err
	}
	defer srcFile.Close()
	gr, err := gzip.NewReader(srcFile)
	if err != nil {
		return err
	}
	defer gr.Close()
	tr := tar.NewReader(gr)
	for {
		hdr, err := tr.Next()
		if err != nil {
			if err == io.EOF {
				break
			} else {
				return err
			}
		}
		filename := filepath.Join(dest, hdr.Name)
		switch hdr.Typeflag {
		case tar.TypeDir:
			if !ExistDir(filename) {
				_ = os.MkdirAll(filename, os.ModePerm)
			}
		case tar.TypeReg:
			file, err := os.OpenFile(filename, os.O_CREATE|os.O_RDWR, os.FileMode(hdr.Mode))
			if err != nil {
				return err
			}
			defer func() {
				_ = file.Close()
			}()
			_, err = io.Copy(file, tr)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
func ExistDir(dirname string) bool {
	fi, err := os.Stat(dirname)
	return (err == nil || os.IsExist(err)) && fi.IsDir()
}

func createFile(name string) (*os.File, error) {
	dir := string([]rune(name)[0:strings.LastIndex(name, "/")])
	err := os.MkdirAll(dir, 0755)
	if err != nil {
		return nil, err
	}
	return os.Create(name)
}
func UnTar(src, dst string) (err error) {
	if !ExistDir(dst) {
		_ = os.MkdirAll(dst, os.ModePerm)
	}
	srcFile, err := os.Open(src)
	if err != nil {
		return err
	}
	defer srcFile.Close()
	tr := tar.NewReader(srcFile)
	for {
		hdr, err := tr.Next()
		if err != nil {
			if err == io.EOF {
				break
			} else {
				return err
			}
		}
		filename := filepath.Join(dst, hdr.Name)
		switch hdr.Typeflag {
		case tar.TypeDir:
			if !ExistDir(filename) {
				_ = os.MkdirAll(filename, os.ModePerm)
			}
		case tar.TypeReg:
			file, err := os.OpenFile(filename, os.O_CREATE|os.O_RDWR, os.FileMode(hdr.Mode))
			if err != nil {
				return err
			}
			defer func() {
				_ = file.Close()
			}()
			_, err = io.Copy(file, tr)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
