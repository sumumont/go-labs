package main

import (
	"fmt"
	"github.com/go-labs/internal/logging"
	"testing"
)

func TestName(t *testing.T) {
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
	i := 2
	for ; i <= 18; i++ {
		label := Label{
			Name:        fmt.Sprintf("%dyinjiao", i),
			Description: fmt.Sprintf("%d引脚", i),
			Color:       "#33ddff",
		}
		template.Labels = append(template.Labels, label)
	}
	logging.Info().Interface("template", template).Send()
	fmt.Println(template)
}
