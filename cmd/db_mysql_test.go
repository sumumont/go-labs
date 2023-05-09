package main

import (
	"context"
	"github.com/go-labs/internal/configs"
	"github.com/go-labs/internal/dao"
	"github.com/go-labs/internal/logging"
	"testing"
)

func TestMysql(t *testing.T) {
	initMysql()
	res, _, err := dao.GetProductDao().FindByParam(context.Background())
	if err != nil {
		panic(err)
	}
	logging.Debug().Interface("result", res).Send()
}
func initMysql() {
	_, err := configs.InitConfig("../configs")
	if err != nil {
		panic(err)
	}
	err = dao.InitMysql()
	if err != nil {
		panic(err)
	}
}
