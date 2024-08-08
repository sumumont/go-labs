package main

import (
	"fmt"
	"testing"
)

const (
	CodeParam = 1
	CodeIo    = iota
	CodeJsonUnmarshal
)

const (
	Code11 = 1
	Code12 = iota
	Code13
)

func TestIota(t *testing.T) {
	fmt.Println(Code13, CodeJsonUnmarshal)
}
