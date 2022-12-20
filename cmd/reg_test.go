package main

import (
	"fmt"
	"regexp"
	"strings"
	"testing"
)

func TestRegNum(t *testing.T) {
	buf := "A1.JPG"

	reg := regexp.MustCompile(`\d+`)
	if reg != nil {
		s := reg.FindAllString(buf, -1)
		result := strings.Join(s, "")
		fmt.Println(result) //[3.14 345.12 7.8]
		//ss := reg.FindAllStringSubmatch(buf, -1)
		//fmt.Println(ss) //[[3.14] [345.12] [7.8]]
	}
}
