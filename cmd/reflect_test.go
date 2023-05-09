package main

import (
	"errors"
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
func TestRefc(t *testing.T) {
	//主键关键
	//todo 真实的子表

	//dest := dto.People{}//Struct
	//var dest []dto.People //Slice
	//var dest dto.People //Pointer
	//var dest []dto.People
	//dest := dto.People{} //Slice
	//ref(&dest)
	//fmt.Println(dest)
}
func ref(dest interface{}) error {
	t := reflect.TypeOf(dest)
	if t.Kind() != reflect.Ptr {
		return errors.New("It must Pointer!")
	}
	v := reflect.ValueOf(dest)
	fmt.Println(v.Elem().Kind())
	if v.Elem().Kind() == reflect.Slice {
		// 取得数组中元素的类型
		tEE := t.Elem().Elem()
		// 数组的值
		vE := v.Elem()
		//fmt.Println(tEE)
		//fmt.Println(vE)

		// new一个数组中的元素对象
		build := reflect.New(tEE)
		// 对象的值
		buildE := build.Elem()
		// 给对象复制
		sONEId := buildE.FieldByName("Id")
		sONEName := buildE.FieldByName("Name")
		sONEId.SetInt(10)
		sONEName.SetString("李四")

		// 创建一个新数组并把元素的值追加进去
		newArr := make([]reflect.Value, 0)
		newArr = append(newArr, build.Elem())

		// 把原数组的值和新的数组合并
		resArr := reflect.Append(vE, newArr...)

		// 最终结果给原数组
		vE.Set(resArr)
	} else {
		return refl(v.Type(), v)
	}
	return nil
}
func ref1(dest interface{}) {
	t := reflect.TypeOf(dest)
	v := reflect.ValueOf(dest)
	kind := v.Kind()
	fmt.Println(kind)
	if kind == reflect.Slice {

		fmt.Println(t.Elem())
		vl := reflect.New(t.Elem().Elem())
		el := vl.Elem()
		el.Field(0).SetInt(1)
		el.Field(1).SetString("hysen1")
		//es := make([]reflect.Value, 0)
		//es = append(es, el)
		//res := el.Interface()
		//fmt.Println(res)
		//newSlice1 := reflect.MakeSlice(t, 0, 0)
		//
		//newSlice1 = reflect.Append(newSlice1, el)
		newS := make([]reflect.Value, 0)
		newS = append(newS, el.Elem())

		vElem := v.Elem()
		vEAddr := reflect.Append(vElem, newS...)
		//res := newSlice1.Interface()
		//fmt.Println(res)
		//v.Elem().Set(newSlice)
		//v.SetPointer(newSlice1.UnsafePointer())
		//v := reflect.ValueOf(dest)
		//dest = res
		vElem.Set(vEAddr)
		for i := 0; i < t.NumField(); i++ {
			refl(v.Index(i).Type(), v.Index(i))
		}
	} else if t.Kind() == reflect.Pointer {
		v := reflect.ValueOf(dest)
		refl(v.Type(), v)
	} else {
		fmt.Println(5)
	}
	fmt.Println(dest)
}

func refl(t reflect.Type, v reflect.Value) error {
	if t.Kind() != reflect.Pointer {
		return errors.New("It must Pointer!")
	}
	e := t.Elem()
	filedNum := e.NumField()
	for i := 0; i < filedNum; i++ {
		tag := e.Field(i).Tag.Get("json")
		if tag == "name" {
			v.Elem().Field(i).SetString("hysen")
		}
	}
	return nil
}
func refAppend(t reflect.Type, v reflect.Value) error {
	if t.Kind() != reflect.Pointer {
		return errors.New("It must Pointer!")
	}
	e := t.Elem()
	filedNum := e.NumField()
	for i := 0; i < filedNum; i++ {
		tag := e.Field(i).Tag.Get("json")
		if tag == "name" {
			v.Elem().Field(i).SetString("hysen")
		}
	}
	return nil
}
