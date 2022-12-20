package main

import (
	"fmt"
	"github.com/go-labs/internal/logging"
	"testing"
)

func TestMap(t *testing.T) {
	next := NextRecheckQuery{Level1Query: map[string]string{
		"product_name": "product_name",
	}}
	Level1QueryConvert(next.Level1Query)
	logging.Debug().Interface("next", next).Send()

	var nexts []NextRecheckQuery
	nexts = append(nexts, NextRecheckQuery{Level1Query: map[string]string{
		"product_name": "product_name",
	}})
	nexts = append(nexts, NextRecheckQuery{Level1Query: map[string]string{
		"product_name": "product_name",
	}})
	nexts = append(nexts, NextRecheckQuery{Level1Query: map[string]string{
		"product_name": "product_name",
	}})
	nexts = append(nexts, NextRecheckQuery{Level1Query: map[string]string{
		"product_name": "product_name",
	}})

	for _, filename := range nexts {
		fmt.Printf("%p\n", &filename)
	}
}
func Level1QueryConvert(level1Query map[string]string) {
	if v, ok := level1Query["product_name"]; ok {
		level1Query["tags_product_name"] = v
		delete(level1Query, "product_name")
	}
	if v, ok := level1Query["serial_number"]; ok {
		level1Query["tags_serial_number"] = v
		delete(level1Query, "serial_number")
	}
}

type NextRecheckQuery struct {
	Level1Query map[string]string `json:"level1Query"`
}
