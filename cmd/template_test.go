package main

import (
	"fmt"
	"html/template"
	"os"
	"strings"
	"testing"
)

func TestTemplate(t *testing.T) {
	tableName := "people"
	autoInfo := AutoInfo{
		//ModulePrefix:     "github.com/apulis/app/apulis-iqi",
		//LogPack:          "github.com/apulisai/sdk/go-utils/logging",
		ModulePrefix: "github.com/go-labs",
		LogPack:      "github.com/go-labs/internal/logging",
		ModelName:    "People",
	}
	autoInfo.SetPrivateModelName()
	writeGo("dto", tableName, autoInfo)
	writeGo("dao", tableName, autoInfo)
	writeGo("services", tableName, autoInfo)
	writeGo("controller", tableName, autoInfo)
}

func TestTemplateDto(t *testing.T) {
	tableName := "people"
	autoInfo := AutoInfo{
		ModulePrefix: "github.com/apulis/app/apulis-iqi",
		LogPack:      "github.com/apulisai/sdk/go-utils/logging",
		ModelName:    "People",
	}
	autoInfo.SetPrivateModelName()
	writeGo("dto", tableName, autoInfo)
}
func TestTemplateDao(t *testing.T) {
	tableName := "people"
	autoInfo := AutoInfo{
		ModulePrefix: "github.com/apulis/app/apulis-iqi",
		LogPack:      "github.com/apulisai/sdk/go-utils/logging",
		ModelName:    "People",
	}
	autoInfo.SetPrivateModelName()
	writeGo("dao", tableName, autoInfo)
}
func TestTemplateService(t *testing.T) {
	tableName := "people"
	autoInfo := AutoInfo{
		ModulePrefix: "github.com/apulis/app/apulis-iqi",
		LogPack:      "github.com/apulisai/sdk/go-utils/logging",
		ModelName:    "People",
	}
	autoInfo.SetPrivateModelName()
	writeGo("services", tableName, autoInfo)
}
func TestTemplateController(t *testing.T) {
	tableName := "people"
	autoInfo := AutoInfo{
		ModulePrefix: "github.com/apulis/app/apulis-iqi",
		LogPack:      "github.com/apulisai/sdk/go-utils/logging",
		ModelName:    "People",
	}
	autoInfo.SetPrivateModelName()
	writeGo("controller", tableName, autoInfo)
}

type AutoInfo struct {
	ModulePrefix     string //代码模块路径前缀
	LogPack          string //日志模块package
	ModelName        string //模型名称
	PrivateModelName string //模型名称
}

func (rec *AutoInfo) SetPrivateModelName() {
	prefix := rec.ModelName[:1]
	prefix = strings.ToLower(prefix)
	rec.PrivateModelName = prefix + rec.ModelName[1:]
}
func writeGo(templateParent string, tabelName string, autoInfo AutoInfo) {
	bytes, err := os.ReadFile(fmt.Sprintf("../configs/struct_template/%s/template.htm", templateParent))
	if err != nil {
		panic(err)
	}
	templateStr := string(bytes)
	tpl, err := template.New("test").Delims("<<", ">>").Parse(templateStr)
	if err != nil {
		panic(err)
	}

	fileName := fmt.Sprintf("%s.go", tabelName)
	file, err := os.Create(fmt.Sprintf("../internal/%s/%s", templateParent, fileName))
	if err != nil {
		panic(err)
	}
	defer file.Close()
	err = tpl.Execute(file, autoInfo)
	if err != nil {
		panic(err)
	}
}
