package main

import (
	"fmt"
	"github.com/go-labs/internal/models"
	"html/template"
	"os"
	"reflect"
	"strings"
	"testing"
)

func TestTemplate(t *testing.T) {
	tableName := "recheck_data"
	autoInfo := FillAutoInfo{
		//ModulePrefix:     "github.com/apulis/app/apulis-iqi",
		//LogPack:          "github.com/apulisai/sdk/go-utils/logging",
		ModulePrefix: "github.com/go-labs",
		LogPack:      "github.com/go-labs/internal/logging",
		//ModelName:    "People",
		Model: models.RecheckData{},
	}
	autoInfo.SetModel()
	writeGo("dto", tableName, autoInfo)
	writeGo("dao", tableName, autoInfo)
	writeGo("services", tableName, autoInfo)
	writeGo("controller", tableName, autoInfo)
}

type FillAutoInfo struct {
	ModulePrefix     string //代码模块路径前缀
	LogPack          string //日志模块package
	Model            interface{}
	ModelName        string //模型名称
	PrivateModelName string //模型名称
	Attr             map[string]interface{}
}

func (rec *FillAutoInfo) SetModel() {
	value := reflect.ValueOf(rec.Model)
	tp := value.Type()
	modelName := tp.Name()
	numField := tp.NumField()
	if numField > 0 {
		rec.Attr = map[string]interface{}{}
	}
	for i := 0; i < numField; i++ {
		field := tp.Field(i)
		if field.Name == "BaseModelId" || field.Name == "UserInfo" || field.Name == "BaseModelTime" {
			continue
		}
		rec.Attr[field.Name] = field.Type.String()
	}
	rec.ModelName = modelName
	rec.PrivateModelName = FirstToLow(modelName)

}
func FirstToLow(str string) string {
	if len(str) < 1 {
		return ""
	}
	prefix := str[:1]
	prefix = strings.ToLower(prefix)
	prefix = prefix + str[1:]
	return prefix
}

var (
	funcMap = template.FuncMap{
		"FirstToLow": FirstToLow,
	}
)

func writeGo(templateParent string, tabelName string, autoInfo FillAutoInfo) {
	bytes, err := os.ReadFile(fmt.Sprintf("../configs/struct_template/%s/template.htm", templateParent))
	if err != nil {
		panic(err)
	}
	templateStr := string(bytes)
	tpl, err := template.New("test").Delims("[[", "]]").Funcs(funcMap).Parse(templateStr)
	if err != nil {
		panic(err)
	}

	fileName := fmt.Sprintf("%s.go", tabelName)
	file, err := os.Create(fmt.Sprintf("../internal/%s/%s", templateParent, fileName))
	if err != nil {
		panic(err)
	}
	defer func() {
		_ = file.Close()
	}()
	err = tpl.Execute(file, autoInfo)
	if err != nil {
		panic(err)
	}
}
