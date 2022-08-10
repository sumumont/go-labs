package main

import (
	"fmt"
	"github.com/go-labs/internal/configs"
	"reflect"
	"testing"
)

func TestAAA(t *testing.T) {

	var str []configs.AppConfig

	var strValue = reflect.ValueOf(&str)

	indirectStr := reflect.Indirect(strValue)
	fmt.Println("indirectStr", indirectStr, indirectStr.Type())
	valueSlice := reflect.MakeSlice(indirectStr.Type(), 100, 1024)

	kind := valueSlice.Kind()

	cap := valueSlice.Cap()

	length := valueSlice.Len()

	fmt.Printf("Type is [%v] with capacity of %v bytes"+
		" and length of %v .", kind, cap, length)
}
