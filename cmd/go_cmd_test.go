package main

import (
	"bytes"
	"fmt"
	"io"
	"os/exec"
	"strconv"
	"syscall"
	"testing"
	"unsafe"
)

func TestA(t *testing.T) {
	fmt.Println("test TestA")
}

//func TestCmd(t *testing.T) {
//	log.Println("test cmd")
//	cmd := exec.Command("ps", "-ef", "|grep", "git")
//	log.Println(cmd.String())
//	stdout, err := cmd.StdoutPipe()
//	if err != nil { //获取输出对象，可以从该对象中读取输出结果
//		panic(err)
//	}
//	defer stdout.Close()
//	err = cmd.Run()
//	if err != nil {
//		panic(err)
//	}
//	log.Println("1")
//
//	if opBytes, err := ioutil.ReadAll(stdout); err != nil { // 读取输出结果
//		log.Println("2")
//		panic(err)
//	} else {
//		log.Println("3")
//		log.Println(string(opBytes))
//	}
//	log.Println("end")
//}
func TestCmdPs(t *testing.T) {
	fmt.Println("TestCmdPs")
	status := getProcess("git2")
	fmt.Println(status)
}
func getProcess(processName string) string {
	ps := exec.Command("ps", "-ef")
	grep := exec.Command("grep", "-i", processName)
	r, w := io.Pipe() // 创建一个管道
	defer r.Close()
	defer w.Close()
	ps.Stdout = w  // ps向管道的一端写
	grep.Stdin = r // grep从管道的一端读
	var buffer bytes.Buffer
	grep.Stdout = &buffer // grep的输出为buffer
	_ = ps.Start()
	_ = grep.Start()
	ps.Wait()
	w.Close()
	grep.Wait()
	return buffer.String()
}

type ulong int32
type ulong_ptr uintptr

type PROCESSENTRY32 struct {
	dwSize              ulong
	cntUsage            ulong
	th32ProcessID       ulong
	th32DefaultHeapID   ulong_ptr
	th32ModuleID        ulong
	cntThreads          ulong
	th32ParentProcessID ulong
	pcPriClassBase      ulong
	dwFlags             ulong
	szExeFile           [260]byte
}

func getProcesses() {
	kernel32 := syscall.NewLazyDLL("kernel32.dll")
	CreateToolhelp32Snapshot := kernel32.NewProc("CreateToolhelp32Snapshot")
	pHandle, _, _ := CreateToolhelp32Snapshot.Call(uintptr(0x2), uintptr(0x0))
	if int(pHandle) == -1 {
		return
	}
	Process32Next := kernel32.NewProc("Process32Next")
	for {
		var proc PROCESSENTRY32
		proc.dwSize = ulong(unsafe.Sizeof(proc))
		if rt, _, _ := Process32Next.Call(uintptr(pHandle), uintptr(unsafe.Pointer(&proc))); int(rt) == 1 {
			fmt.Println("ProcessName : " + string(proc.szExeFile[0:]))
			fmt.Println("th32ModuleID : " + strconv.Itoa(int(proc.th32ModuleID)))
			fmt.Println("ProcessID : " + strconv.Itoa(int(proc.th32ProcessID)))
		} else {
			break
		}
	}
	CloseHandle := kernel32.NewProc("CloseHandle")
	_, _, _ = CloseHandle.Call(pHandle)
}

func TestProcesses(t *testing.T) {
	getProcesses()
}
