package main

import (
	"fmt"
	"strings"
	"testing"
)

type ArrayField struct {
	FieldNameInDB string
	ElementType   string
	Value         *[]interface{}
}

func TestStr1(t *testing.T) {
	ls := []interface{}{"dsadas", "dsadas"}
	elem := ArrayField{
		FieldNameInDB: "",
		ElementType:   "",
		Value:         &ls,
	}
	var sqlStr = ""
	sqlStr += fmt.Sprintf(`INSERT INTO public.%s `, "tbl_122")
	//columnsStr := "("
	columnValuesStr := "("
	strSlice := make([]string, len(*elem.Value))
	for idx, v := range *elem.Value {
		strSlice[idx] = fmt.Sprintf("'%s'", v.(string))
	}

	columnValuesStr += fmt.Sprintf(`ARRAY [%s], `, strings.Join(strSlice, ", "))
	columnValuesStr = strings.TrimSuffix(columnValuesStr, ", ") + ")\n"
	fmt.Println(columnValuesStr)
}

func TestNew(t *testing.T) {
	a := fmt.Sprintf("%.2f", 11.11)
	fmt.Println(a)
	a = fmt.Sprintf("%.2f", 11.00)
	fmt.Println(a)
}
