package dao

import (
	"context"
	"errors"
	"github.com/go-labs/internal/logging"
	"github.com/go-labs/internal/models"
	"gorm.io/gorm/clause"
)

type ProductDao struct {
}

var productDao = ProductDao{}

func GetProductDao() ProductDao {
	return productDao
}
func (svc ProductDao) FindByParam(ctx context.Context) ([]models.Product, int64, error) {
	return nil, 0, nil
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
