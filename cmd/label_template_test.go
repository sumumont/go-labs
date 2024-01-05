package main

import (
	"fmt"
	"github.com/go-labs/internal/logging"
	"testing"
)

type LabelTemplate struct {
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Labels      []Label `json:"labels"`
	Fields      []Field `json:"fields"`
}

type Label struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Color       string `json:"color"`
}
type Field struct {
	Field string `json:"field"`
	Task  string `json:"task"`
}

func TestName1(t *testing.T) {

	template := LabelTemplate{
		Name:        "冠捷推理模板-贴片",
		Description: "",
		Labels: []Label{
			{
				Name:        "konghanpan",
				Description: "空焊盘",
				Color:       "#fa3253",
			},
		},
		Fields: nil,
	}
	for i := 2; i <= 18; i++ {
		label := Label{
			Name:        fmt.Sprintf("%dyinjiao", i),
			Description: fmt.Sprintf("%d引脚", i),
			Color:       "#33ddff",
		}
		template.Labels = append(template.Labels, label)
	}
	logging.Info().Interface("template", template).Send()
	fmt.Println("dsadsadas")
}
