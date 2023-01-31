package main

import (
	"context"
	"github.com/go-labs/internal/configs"
	"github.com/go-labs/internal/dao"
	"github.com/go-labs/internal/logging"
	"github.com/go-labs/internal/models"
	"github.com/go-labs/internal/utils"
	"testing"
)

func TestDb(t *testing.T) {
	initDb()
	defer utils.TimeCost("insert")()
	err := dao.ExecDBTx(func(ctx context.Context) error {

		var products []models.Product
		for i := 1; i <= 10; i++ {
			attr := models.JsonB{}
			product := models.Product{
				Num: int64(i),
			}
			attr["name"] = "hysen"
			attr["num"] = product.Num
			product.Attributes = attr
			products = append(products, product)
		}
		err := dao.GetProductDao().SaveOrUpdates(ctx, products, []string{"num"})
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		panic(err)
	}

}
func initDb() {
	_, err := configs.InitConfig("../configs")
	if err != nil {
		panic(err)
	}
	err = dao.InitDb()
	if err != nil {
		panic(err)
	}
}
func TestDelete(t *testing.T) {
	initDb()
	defer utils.TimeCost("delete")()
	err := dao.ExecDBTx(func(ctx context.Context) error {
		err := dao.GetProductDao().DeleteBatch(ctx, 1200)
		if err != nil {
			return err
		}
		//logging.Debug().Interface("ids", ids).Send()
		return nil
	})
	if err != nil {
		panic(err)
	}
}
func TestDeleteByJsonB(t *testing.T) {
	initDb()
	defer utils.TimeCost("delete")()
	err := dao.ExecDBTx(func(ctx context.Context) error {
		attr := models.JsonB{}
		attr["name"] = "hysen"
		attr["num"] = 1
		err := dao.GetProductDao().DeleteByAttr(ctx, attr)
		if err != nil {
			return err
		}
		//logging.Debug().Interface("ids", ids).Send()
		return nil
	})
	if err != nil {
		panic(err)
	}
}
func TestFindByJsonB(t *testing.T) {
	initDb()
	defer utils.TimeCost("delete")()
	err := dao.ExecDBTx(func(ctx context.Context) error {
		attr := models.JsonB{}
		attr["name"] = "hysen"
		attr["num"] = 1
		ids, err := dao.GetProductDao().FindByAttr(ctx, attr)
		if err != nil {
			return err
		}
		logging.Debug().Interface("ids", ids).Send()
		return nil
	})
	if err != nil {
		panic(err)
	}
}
