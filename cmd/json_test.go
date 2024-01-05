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

type Devices []Device
type Device struct {
	DeviceMeta     DeviceMeta   `json:"DeviceMeta"`
	DeviceStatic   DeviceStatic `json:"DeviceStatic"`
	DeviceID       string       `json:"DeviceID"`
	Index          int64        `json:"Index"`
	Memory         int64        `json:"Memory"`         //总显存
	FreeMemory     int64        `json:"FreeMemory"`     //剩余显存
	LeftVDeviceNum int64        `json:"LeftVDeviceNum"` //剩余的虚拟化卡数
	MaxVDeviceNum  int64        `json:"MaxVDeviceNum"`  // 总的虚拟化卡数
	LeftRatio      int64        `json:"LeftRatio"`      // 剩余算力 LeftRatio : 90
	Ratio          int64        `json:"Ratio"`          //总算力100%
	InUsed         bool         `json:"InUsed"`         //正在占用
	NativeDevice   bool         `json:"NativeDevice"`   //是否是本地设备(物理卡) false 表示已经被虚拟化
	LeftVGPUNumber int64        `json:"LeftVGPUNumber"`
	MaxVGPUNumber  int64        `json:"MaxVGPUNumber"`
	Source         string       `json:"Source"`
	ProcessState   string       `json:"ProcessState"`
	ProcessSource  string       `json:"ProcessSource"`
	IsMasked       bool         `json:"IsMasked"`
	Usage          Usage        `json:"Usage"`
	ServerInfo     struct {
		Enable bool `json:"Enable"`
	} `json:"ServerInfo"`
	DefaultRatio    int64       `json:"DefaultRatio"`
	DefaultMemory   int64       `json:"DefaultMemory"`
	CanBeAllocAgain bool        `json:"CanBeAllocAgain"`
	VGPUs           []VGpu      `json:"VGPUs"`
	RatioType       interface{} `json:"RatioType"`
}
type DeviceMeta struct {
	DeviceIp              string      `json:"device_ip"` //节点ip
	DevicePort            int64       `json:"device_port"`
	DeviceIndex           int64       `json:"device_index"`
	DeviceId              string      `json:"device_id"`
	DeviceName            string      `json:"device_name"`   //device_name : "MLU370-X8"
	DeviceVendor          string      `json:"device_vendor"` //device_vendor :"Cambricon"
	PlatformName          string      `json:"platform_name"`
	PlatformVendor        string      `json:"platform_vendor"`
	DeviceType            string      `json:"device_type"` //device_type MLU||GPU
	Protocols             []string    `json:"protocols"`
	Sharable              bool        `json:"sharable"`
	VdeviceNum            int64       `json:"vdevice_num"` //虚拟化的卡数 vdevice_num : 4
	FixedVdeviceFlavor    int64       `json:"fixed_vdevice_flavor"`
	Labels                interface{} `json:"labels"`
	InitForceNativeDevice bool        `json:"init_force_native_device"`
	OcEnable              bool        `json:"oc_enable"`
	OcCoreRatio           int64       `json:"oc_core_ratio"`
	OcMemRatio            int64       `json:"oc_mem_ratio"`
	Protocol              string      `json:"protocol"`
}
type DeviceStatic struct {
	DeviceId      string `json:"device_id"`
	DriverVersion string `json:"driver_version"`
	Memory        int64  `json:"memory"`
	CpuAffinity   int64  `json:"cpu_affinity"`
	Bandwidth     int64  `json:"bandwidth"`
}
type Usage struct { //已使用数量
	VDevices []VDevice `json:"VDevices"`
}
type VDevice struct {
	Index  int64 `json:"Index"`
	Memory int64 `json:"Memory"` //已分配显存
	Ratio  int64 `json:"Ratio"`  // 已分配算力
	//BookingAllocationIDs   []interface{} `json:"BookingAllocationIDs"`
	ConfirmedAllocationIDs []string `json:"ConfirmedAllocationIDs"`
}
type VGpu struct {
	Index  int64 `json:"Index"`
	Memory int64 `json:"Memory"`
	Ratio  int64 `json:"Ratio"`
	//BookingAllocationIDs   []interface{} `json:"BookingAllocationIDs"`
	ConfirmedAllocationIDs []string `json:"ConfirmedAllocationIDs"`
}

type DevicesData struct {
	Total int64   `json:"total"`
	Items Devices `json:"items"`
}

func TestDevicesData(t *testing.T) {
	jsonStr := `{
        "total": 1,
        "items": [
            {
                "DeviceMeta": {
                    "device_ip": "10.133.34.214",
                    "device_port": 9960,
                    "device_index": 0,
                    "device_id": "GPU-e9a1fcbc-3172-9aaf-f403-0f8b2e1d453f",
                    "device_name": "NVIDIA GeForce RTX 3090",
                    "device_vendor": "NVidia",
                    "platform_name": "NVidia",
                    "platform_vendor": "NVidia",
                    "device_type": "GPU",
                    "protocols": [
                        "CUDA"
                    ],
                    "sharable": true,
                    "vdevice_num": 4,
                    "fixed_vdevice_flavor": 0,
                    "labels": null,
                    "init_force_native_device": false,
                    "oc_enable": true,
                    "oc_core_ratio": 0,
                    "oc_mem_ratio": 0,
                    "protocol": "CUDA"
                },
                "DeviceStatic": {
                    "device_id": "GPU-e9a1fcbc-3172-9aaf-f403-0f8b2e1d453f",
                    "driver_version": "525.105.17",
                    "memory": 24576,
                    "cpu_affinity": 0,
                    "bandwidth": 2500
                },
                "DeviceID": "GPU-e9a1fcbc-3172-9aaf-f403-0f8b2e1d453f",
                "Index": 0,
                "Memory": 24576,
                "Ratio": 100,
                "Source": "open_virtualization",
                "ProcessState": "",
                "ProcessSource": "",
                "MaxVDeviceNum": 4,
                "IsMasked": false,
                "Usage": {
                    "VDevices": [
                        {
                            "Index": 0,
                            "Memory": 1024,
                            "Ratio": 15,
                            "BookingAllocationIDs": [],
                            "ConfirmedAllocationIDs": [
                                "f21d0e44-4e83-448e-8499-8ae04979ddb4"
                            ]
                        }
                    ]
                },
                "ServerInfo": {
                    "Enable": true
                },
                "NativeDevice": false,
                "InUsed": true,
                "DefaultRatio": 25,
                "DefaultMemory": 6144,
                "FreeMemory": 23552,
                "LeftRatio": 85,
                "CanBeAllocAgain": true,
                "LeftVDeviceNum": 3,
                "MaxVGPUNumber": 4,
                "LeftVGPUNumber": 3,
                "VGPUs": [
                    {
                        "Index": 0,
                        "Memory": 1024,
                        "Ratio": 15,
                        "BookingAllocationIDs": [],
                        "ConfirmedAllocationIDs": [
                            "f21d0e44-4e83-448e-8499-8ae04979ddb4"
                        ]
                    }
                ],
                "RatioType": null
            }
        ]
    }`

	rsp := DevicesData{}
	if err := json.Unmarshal([]byte(jsonStr), &rsp); err != nil {
		panic(err)
	}
	logging.Debug().Interface("rsp", rsp).Send()
}
