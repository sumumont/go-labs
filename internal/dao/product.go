package dao

import (
	"context"
	"errors"
	"github.com/go-labs/internal/logging"
	"github.com/go-labs/internal/models"
	"gorm.io/datatypes"
	"gorm.io/gorm/clause"
)

type ProductDao struct {
}

var productDao = ProductDao{}

func GetProductDao() ProductDao {
	return productDao
}
func (svc ProductDao) FindByParam(ctx context.Context) ([]models.Product, int64, error) {
	var result []models.Product
	tx := GetDB(ctx)
	err := tx.Where("num", 1).Find(&result).Error
	if err != nil {
		return nil, 0, err
	}
	return result, int64(len(result)), nil
}
func (svc ProductDao) Create(ctx context.Context, model *models.Product) error {
	tx := GetDB(ctx)
	err := tx.Create(model).Error
	if err != nil {
		logging.Error(err).Send()
		return err
	}
	return nil
}
func (svc ProductDao) CreateBatch(ctx context.Context, model []models.Product) error {
	tx := GetDB(ctx)
	err := tx.Create(model).Error
	if err != nil {
		logging.Error(err).Send()
		return err
	}
	return nil
}
func (svc ProductDao) SaveOrUpdates(ctx context.Context, model []models.Product, UpdateColumns []string) error {
	tx := GetDB(ctx)
	err := tx.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "id"}},
		DoUpdates: clause.AssignmentColumns(UpdateColumns),
	}).Create(model).Error
	if err != nil {
		logging.Error(err).Send()
		return err
	}
	return nil
}
func (svc ProductDao) DeleteBatch(ctx context.Context, num int) error {
	tx := GetDB(ctx)
	var ids []models.Product
	logging.Error(errors.New("test error")).Send()
	err := tx.Where("num <= ?", num).Find(&ids).Delete(&models.Product{}).Error
	//tx.Model(&models.Product{}).Raw("delete from ")
	if err != nil {
		logging.Error(err).Send()
		return err
	}
	logging.Debug().Interface("rows", ids).Send()
	return nil
}
func (svc ProductDao) DeleteByAttr(ctx context.Context, attr map[string]interface{}) error {
	tx := GetDB(ctx)
	var ids []models.Product
	logging.Error(errors.New("test error")).Send()
	var querys []*datatypes.JSONQueryExpression
	for k, v := range attr {
		tx = tx.Where(datatypes.JSONQuery("attributes").Equals(v, k))
	}
	tx = tx.Delete(&ids, querys)
	err := tx.Error
	if err != nil {
		logging.Error(err).Send()
		return err
	}
	logging.Debug().Interface("rows", ids).Send()
	return nil
}
func (svc ProductDao) FindByAttr(ctx context.Context, attr map[string]interface{}) ([]models.Product, error) {
	tx := GetDB(ctx)
	var ids []models.Product
	for k, v := range attr {
		//tx = tx.Where(datatypes.JSONQuery("attributes").Equals(v, k))
		tx = tx.Where(datatypes.JSONQuery("attributes").Equals(v, k))
	}
	tx = tx.Find(&ids)
	err := tx.Error
	if err != nil {
		logging.Error(err).Send()
		return ids, err
	}
	logging.Debug().Interface("rows", ids).Send()
	return ids, nil
}
