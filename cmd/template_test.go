package main

import (
	"fmt"
	"html/template"
	"os"
	"testing"
)

var (
	//modulePrefix = "github.com/go-labs"
	//logPack      = "github.com/go-labs/internal/logging"
	modulePrefix = "github.com/apulis/app/apulis-iqi"
	logPack      = "github.com/apulisai/sdk/go-utils/logging"

	tableName        = "people"
	modelName        = "People"
	privateModelName = "people"
)

func TestTemplate(t *testing.T) {
	writeGo("dto", tableName)
	writeGo("dao", tableName)
	writeGo("services", tableName)
	writeGo("controller", tableName)
}

func TestTemplateDto(t *testing.T) {
	writeGo("dto", tableName)
}
func TestTemplateDao(t *testing.T) {
	writeGo("dao", tableName)
}
func TestTemplateService(t *testing.T) {
	writeGo("services", tableName)
}
func TestTemplateController(t *testing.T) {
	writeGo("controller", tableName)
}

func writeGo(templateParent string, tabelName string) {
	bytes, err := os.ReadFile(fmt.Sprintf("../configs/struct_template/%s/template.htm", templateParent))
	if err != nil {
		panic(err)
	}
	templateStr := string(bytes)
	tpl, err := template.New("test").Delims("<<", ">>").Parse(templateStr)
	if err != nil {
		panic(err)
	}

	data := map[string]interface{}{
		"modulePrefix":     modulePrefix,
		"modelName":        modelName,
		"privateModelName": privateModelName,
		"logging":          logPack,
	}
	fileName := fmt.Sprintf("%s.go", tabelName)
	file, err := os.Create(fmt.Sprintf("../internal/%s/%s", templateParent, fileName))
	if err != nil {
		panic(err)
	}
	defer file.Close()
	err = tpl.Execute(file, data)
	if err != nil {
		panic(err)
	}
}
