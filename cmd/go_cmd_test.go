package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os/exec"
	"strconv"
	"strings"
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
	result := getProcess("git2")
	fmt.Println(result)
	lines := strings.Split(result, "\n")
	for idx, line := range lines {
		if line == "" {
			continue
		}
		fmt.Println("idx:", idx, " line:", line)
		columns := strings.Split(line, "\t")
		for i, column := range columns {
			fmt.Println("i:", i, " column:", column)
		}
	}
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

type Process struct {
	pid int
	cpu float64
}

func getPcs() {
	cmd := exec.Command("ps", "aux")
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
	processes := make([]*Process, 0)
	for {
		line, err := out.ReadString('\n')
		if err != nil {
			break
		}
		tokens := strings.Split(line, " ")
		ft := make([]string, 0)
		for _, t := range tokens {
			if t != "" && t != "\t" {
				ft = append(ft, t)
			}
		}
		log.Println(len(ft), ft)
		pid, err := strconv.Atoi(ft[1])
		if err != nil {
			continue
		}
		cpu, err := strconv.ParseFloat(ft[2], 64)
		if err != nil {
			log.Fatal(err)
		}
		processes = append(processes, &Process{pid, cpu})
	}
	for _, p := range processes {
		log.Println("Process ", p.pid, " takes ", p.cpu, " % of the CPU")
	}
}

func TestPcs(t *testing.T) {
	getPcs()
}
