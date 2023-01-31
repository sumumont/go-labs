package main

import (
	"fmt"
	"testing"
	"time"
)

func TestTime(t *testing.T) {
	fmt.Println(time.Now().UnixMilli() / 1000)
}
