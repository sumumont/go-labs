package dao

import (
	"context"
	"github.com/go-labs/internal/logging"
	"github.com/go-labs/internal/models"
	"gorm.io/gorm/clause"
)

type TheClassDao struct {
}

var studentDao = TheClassDao{}

func GetTheClassDao() TheClassDao {
	return studentDao
}
func (svc TheClassDao) FindById(ctx context.Context, id int64) (*models.TheClass, error) {
	var result models.TheClass
	tx := GetDB(ctx)
	err := tx.Preload("Students").Where("id", id).First(&result).Error
	if err != nil {
		return nil, err
	}
	return &result, nil
}
func (svc TheClassDao) Find(ctx context.Context) ([]models.TheClass, error) {
	var result []models.TheClass
	tx := GetDB(ctx)
	err := tx.Preload("Students").Find(&result).Error
	if err != nil {
		return nil, err
	}
	return result, nil
}
func (svc TheClassDao) Create(ctx context.Context, model *models.TheClass) error {
	tx := GetDB(ctx)
	err := tx.Create(model).Error
	if err != nil {
		logging.Error(err).Send()
		return err
	}
	return nil
}
func (svc TheClassDao) CreateBatch(ctx context.Context, model []models.TheClass) error {
	tx := GetDB(ctx)
	err := tx.Create(model).Error
	if err != nil {
		logging.Error(err).Send()
		return err
	}
	return nil
}
func (svc TheClassDao) SaveOrUpdates(ctx context.Context, model []models.TheClass, UpdateColumns []string) error {
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
