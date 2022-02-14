package main

import (
	"bytes"
	"fmt"
	"io"
	"os/exec"
	"testing"
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
	status := getProcess("git")
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
