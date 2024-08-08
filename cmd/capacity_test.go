package main

import (
	"archive/tar"
	"archive/zip"
	"bytes"
	"compress/gzip"
	"errors"
	"fmt"
	"github.com/go-labs/internal/logging"
	"github.com/samber/lo"
	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"testing"
)

func TestCapacity(t *testing.T) {

	//capacity := Capacity([]string{"train"})

	//fmt.Println(capacity.AllCap(), capacity.OnlyTrain(), capacity.OnlyInfer())
	//a := []string{}
	//b := []string{"b", "a"}
	//fmt.Println(IsEqual(a, b))
	s := "apulis/orgadmin-user-group/11/14"
	pvPaths := filepath.Join(strings.TrimSuffix(s, filepath.Base(s)))
	fmt.Println(pvPaths)
}

type Array []string

func (rec Array) Equal(array Array) bool {
	o := lo.Uniq[string](rec)
	b := lo.Uniq[string](array)
	if len(o) != len(b) {
		return false
	}
	for _, item := range o {
		if !lo.Contains[string](b, item) {
			return false
		}
	}
	return true
}

func IsEqual(a, b []string) bool {
	x := Array(a)
	return x.Equal(b)
}

func TestFp(t *testing.T) {

	//capacity := Capacity([]string{"train"})

	//fmt.Println(capacity.AllCap(), capacity.OnlyTrain(), capacity.OnlyInfer())

	fmt.Println(filepath.Join("/", "code"))
	fmt.Println(filepath.Join("code"))
	fmt.Println(filepath.Join("/code"))
	fmt.Println(filepath.Join("code/"))
}
func TestZipHandler(t *testing.T) {
	zipHandler("D:\\OneDrive\\OneDrive - 依瞳科技（深圳）有限公司\\桌面\\model_sample (1)\\model_sample_utf8.zip", "", "code")
}
func TestStandard(t *testing.T) {
	standard("D:\\Program Files (x86)\\WXWork\\data\\WXWork\\1688855031533911\\Cache\\File\\2023-11\\code+model+manifest.tar", "code+model+manifest.tar")
}

func zipHandler(src, dest string, zipDir string) error {

	r, err := zip.OpenReader(src)
	if err != nil {
		return err
	}

	defer r.Close()
	if len(zipDir) != 0 {
		zipDir = filepath.Join(zipDir)
		dest = filepath.Join(dest, zipDir)
	}

	for _, f := range r.File {
		fmt.Println(f.Name)
		fname := f.Name
		if f.Flags == 0 {
			//如果标致位是0  则是默认的本地编码   默认为gbk
			i := bytes.NewReader([]byte(f.Name))
			decoder := transform.NewReader(i, simplifiedchinese.GB18030.NewDecoder())
			content, _ := ioutil.ReadAll(decoder)
			fname = string(content)
		} else {
			//如果标志为是 1 << 11也就是 2048  则是utf-8编码
			fname = f.Name
		}
		fmt.Println(fname)
	}

	return nil
}
func tarGzHandler(src, dst string) (err error) {
	// 打开准备解压的 tar 包
	fr, err := os.Open(src)
	if err != nil {
		return
	}
	defer fr.Close()

	// 将打开的文件先解压
	gr, err := gzip.NewReader(fr)
	if err != nil {
		return
	}
	defer gr.Close()

	// 通过 gr 创建 tar.Reader
	tr := tar.NewReader(gr)

	// 现在已经获得了 tar.Reader 结构了，只需要循环里面的数据写入文件就可以了
	for {
		hdr, err := tr.Next()

		switch {
		case err == io.EOF:
			return nil
		case err != nil:
			return err
		case hdr == nil:
			continue
		}

		// 处理下保存路径，将要保存的目录加上 header 中的 Name
		// 这个变量保存的有可能是目录，有可能是文件，所以就叫 FileDir 了……
		dstFileDir := filepath.Join(dst, hdr.Name)

		// 根据 header 的 Typeflag 字段，判断文件的类型
		switch hdr.Typeflag {
		case tar.TypeDir: // 如果是目录时候，创建目录
			// 判断下目录是否存在，不存在就创建
			if b := ExistDir(dstFileDir); !b {
				// 使用 MkdirAll 不使用 Mkdir ，就类似 Linux 终端下的 mkdir -p，
				// 可以递归创建每一级目录
				if err := os.MkdirAll(dstFileDir, 0775); err != nil {
					return err
				}
			}
		case tar.TypeReg: // 如果是文件就写入到磁盘
			// 创建一个可以读写的文件，权限就使用 header 中记录的权限
			// 因为操作系统的 FileMode 是 int32 类型的，hdr 中的是 int64，所以转换下
			file, err := os.OpenFile(dstFileDir, os.O_CREATE|os.O_RDWR, os.FileMode(hdr.Mode))
			if err != nil {
				return err
			}
			defer func() {
				file.Close()
			}()
			_, err = io.Copy(file, tr)
			if err != nil {

				return err
			}
		}
	}
}
func isStandard(filename string) bool {
	//if strings.HasPrefix(filename, "code") || strings.HasPrefix(filename, "/code") {
	//	return true
	//}
	//if strings.HasPrefix(filename, "infer") || strings.HasPrefix(filename, "/infer") {
	//	return true
	//}
	return filename == "manifest_v2.yaml"
}
func standard(filePath, fileName string) (bool, error) {
	// decompress file
	// check file format
	checkZipAndTarRe := `.*\.(zip|tar\.gz|tar)$`
	r := regexp.MustCompile(checkZipAndTarRe)
	logging.Info().Msgf("fileName: %v", fileName)
	if !(r.MatchString(fileName)) {
		logging.Error(errors.New("ErrorModelInvalidPostfix")).Send()
		return false, errors.New("ErrorModelInvalidPostfix")
	}

	isZipRe := `.*\.zip$`
	//isTarRe := `.*\.tar$`
	rZip := regexp.MustCompile(isZipRe)
	//rTar := regexp.MustCompile(isTarRe)
	if rZip.MatchString(fileName) {
		{
			r, err := zip.OpenReader(filePath)
			if err != nil {
				logging.Error(err).Send()
				return false, err
			}
			defer r.Close()
			for _, f := range r.File {
				if !f.FileInfo().IsDir() && isStandard(f.Name) {
					return true, err
				}
			}
		}
	} else {
		// 打开要解包的文件
		fr, err := os.Open(filePath)
		if err != nil {
			return false, err
		}
		defer fr.Close()

		// 创建 tar.Reader，准备执行解包操作
		tr := tar.NewReader(fr)

		// 遍历包中的文件
		for hdr, er := tr.Next(); er != io.EOF; hdr, er = tr.Next() {
			if err != nil {
				return false, err
			}
			if isStandard(hdr.FileInfo().Name()) {
				return true, errors.New("ErrorModelInvalidPostfix")
			}
		}
	}
	return false, nil
}

type StudioModelVersion struct {
	Capacities []string
}

func TestArray(t *testing.T) {
	version := &StudioModelVersion{Capacities: []string{"infer", "train", "eval"}}
	removeCapacity(version, "train")
	fmt.Println(version)
}
func removeCapacity(version *StudioModelVersion, cap string) {
	caps := lo.Uniq[string](version.Capacities)

	//caps = lo.DropWhile[string](caps, func(item string) bool {
	//	return !(item == cap)
	//})
	caps = DropAny[string](caps, func(item string) bool {
		return item == cap
	})
	version.Capacities = caps
}
func DropAny[T any](collection []T, predicate func(item T) bool) []T {
	result := make([]T, 0)
	for i := 0; i < len(collection); i++ {
		if !predicate(collection[i]) {
			result = append(result, collection[i])
		}
	}

	return result
}
