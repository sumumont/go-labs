package exports

import (
	"fmt"
	"sort"
	"testing"
)

func TestMsg(t *testing.T) {
	keys := []int{}
	for k, _ := range errText {
		//fmt.Println(k, v)
		keys = append(keys, k)
	}
	sort.Ints(keys)
	for _, k := range keys {
		fmt.Println(k, errText[k])
	}
}
