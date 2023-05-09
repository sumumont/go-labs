package main

import (
	"fmt"
	"github.com/google/uuid"
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
	fmt.Println(uuid.New().ID())
	// Generate 10 ids
	//ids := make([]xid.ID, 10)
	//ids := map[string]xid.ID{}
	//tm := time.Now()
	//for i := 0; i < 100; i++ {
	//	id := xid.NewWithTime(tm)
	//	ids[id.String()] = id
	//	fmt.Println(i, id.String())
	//}
	//fmt.Println("===============================")
	//for k, v := range ids {
	//	fmt.Println(k, v.String())
	//}
}
