package main

import (
	"io/ioutil"
	"log"
	"os/exec"
	"testing"
)

func TestA(t *testing.T) {
	log.Println("test TestA")
}
func TestCmd(t *testing.T) {
	log.Println("test cmd")
	cmd := exec.Command("ps", "-ef", "|grep git")
	log.Println(cmd.String())
	stdout, err := cmd.StdoutPipe()
	if err != nil { //获取输出对象，可以从该对象中读取输出结果
		panic(err)
	}
	defer stdout.Close()
	if opBytes, err := ioutil.ReadAll(stdout); err != nil { // 读取输出结果
		panic(err)
	} else {
		log.Println(string(opBytes))
	}
}
