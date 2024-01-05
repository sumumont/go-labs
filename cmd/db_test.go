package main

import (
	"context"
	"github.com/go-labs/internal/configs"
	"github.com/go-labs/internal/dao"
	"github.com/go-labs/internal/logging"
	"github.com/go-labs/internal/models"
	"github.com/go-labs/internal/utils"
	"gorm.io/gorm"
	"testing"
	"time"
)

func TestDb(t *testing.T) {
	initDb()
	defer utils.TimeCost("insert")()
	err := dao.ExecDBTx(func(ctx context.Context) error {

		var products []models.Product
		for i := 1; i <= 20; i++ {
			attr := models.JsonB{}
			now := utils.GetNowTime()
			product := models.Product{
				Num: int64(i),
				BaseModelTime: models.BaseModelTime{
					CreatedAt: now,
					UpdatedAt: now,
				},
			}
			attr["name"] = "hysen"
			attr["num"] = product.Num
			product.Attributes = attr
			products = append(products, product)
			//time.Sleep(time.Second)
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

func TestQuery(t *testing.T) {
	initDb()
	result := []map[string]interface{}{}
	tx := dao.GetDB(context.Background()).Model(&models.Student{}).Select("students.*,b.*").Joins("left join schools as b on students.school_id = b.id")

	rows, err := tx.Rows()
	if err != nil {
		panic(err)
	}

	cl, err := rows.Columns()
	logging.Debug().Interface("cl", cl).Send()
	tx = tx.Find(&result)
	err = tx.Error
	if err != nil {
		panic(err)
	}
	logging.Debug().Interface("result", cl).Send()

}
func TestQuery1(t *testing.T) {
	// 设置 GORM 的连接参数

	initDb()

	tx := dao.GetDB(context.Background())
	// 连续占用连接
	for i := 1; i <= 20; i++ {
		go occupyConnection(tx, i)
	}

	// 阻止程序退出
	select {}
}

func occupyConnection(db *gorm.DB, id int) {
	for {
		attr := models.JsonB{}
		attr["name"] = "hysen"
		attr["num"] = 1
		ids, err := dao.GetProductDao().FindByAttr(context.Background(), attr)
		if err != nil {
			panic(err)
		}
		logging.Debug().Interface("num", id).Interface("ids", ids).Send()

		time.Sleep(120 * time.Second)
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
func TestDbReflect(t *testing.T) {
	initDb()
	err := dao.ExecDBTx(func(ctx context.Context) error {
		//model := &models.People{}
		dao.GetDB(ctx).Table("peoples").Find("")
		return nil
	})
	if err != nil {
		panic(err)
	}
}

type PredictParam struct {
	Tags          Tags                   `json:"tags"`
	DataFormat    string                 `json:"data_format"`
	RequestParams map[string]interface{} `json:"request_params"`
	Requests      []Request              `json:"requests"`
}
type Tags map[string]interface{}
type Request struct {
	Data         string                 `json:"data"`
	Name         string                 `json:"name"`
	RequestParam map[string]interface{} `json:"request_param"`
}

func TestA1(t *testing.T) {
	var param = PredictParam{
		Tags:       nil,
		DataFormat: "dsadsa",
		RequestParams: map[string]interface{}{
			"name": "hysen",
		},
		Requests: []Request{
			{
				Data:         "1",
				Name:         "1",
				RequestParam: map[string]interface{}{},
			}, {
				Data:         "2",
				Name:         "2",
				RequestParam: map[string]interface{}{},
			},
		},
	}
	DispatcherParam := param
	logging.Debug().Interface("dsad", DispatcherParam).Send()
	for i, _ := range param.Requests {
		param.Requests[i].RequestParam["coordinate"] = "realCoordinate"
	}
	logging.Debug().Interface("realCoordinates", DispatcherParam.Requests).Send()
}
