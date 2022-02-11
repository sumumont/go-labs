package main

import (
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"testing"
)

func TestCmd(t *testing.T) {
	cmd := exec.Command("ps", "-ef", "|grep git")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
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
