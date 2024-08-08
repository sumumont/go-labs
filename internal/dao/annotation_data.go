package dao

import (
	"context"
	"github.com/go-labs/internal/dto"
	"github.com/go-labs/internal/utils"

	"github.com/go-labs/internal/logging"
	"github.com/go-labs/internal/models"
	"github.com/go-labs/pkg/exports"
	"gorm.io/gorm/clause"
)

type AnnotationDataDao struct {
}

var annotationDataDao = AnnotationDataDao{}

func GetAnnotationDataDao() AnnotationDataDao {
	return annotationDataDao
}

func (svc AnnotationDataDao) FindByParam(ctx context.Context, param dto.AnnotationDataParam, userInfo *models.UserInfo) ([]models.AnnotationData, int64, exports.APIError) {
	baseList := param.Conversion()
	db := GetDB(ctx).Model(&models.AnnotationData{})
	var p []models.AnnotationData
	var total int64
	db.Count(&total)
	db, errs := models.PageList(db, baseList)
	if errs != nil {
		logging.Error(errs).Send()
		return nil, 0, exports.RaiseAPIErrorAuto(exports.APULIS_IQI_DB_QUERY_FAILED, errs.Error())
	}
	if total == 0 {
		return nil, 0, nil
	}
	err := db.Find(&p).Error
	if err != nil {
		logging.Error(err).Send()
		return nil, 0, exports.RaiseAPIErrorAuto(exports.APULIS_IQI_DB_QUERY_FAILED, err.Error())
	}
	return p, total, nil
}

func (svc AnnotationDataDao) Create(ctx context.Context, model *models.AnnotationData) exports.APIError {
	tx := GetDB(ctx)
	err := tx.Create(model).Error
	if err != nil {
		logging.Error(err).Send()
		return exports.RaiseAPIErrorAuto(exports.APULIS_IQI_DB_UPDATE_UNEXPECT, err.Error())
	}
	return nil
}

func (svc AnnotationDataDao) FindById(ctx context.Context, id int64) (*models.AnnotationData, exports.APIError) {
	tx := GetDB(ctx)
	p := models.AnnotationData{}
	tx = tx.Where("id = ?", id)
	err := tx.First(&p).Error
	if err != nil {
		logging.Error(err).Send()
		return nil, exports.RaiseAPIErrorAuto(exports.APULIS_IQI_DB_UPDATE_UNEXPECT, err.Error())
	}
	return &p, nil
}
func (svc AnnotationDataDao) Delete(ctx context.Context, id int64) exports.APIError {
	tx := GetDB(ctx)
	tx = tx.Where("id = ?", id)
	err := tx.Delete(models.AnnotationData{}).Error
	if err != nil {
		logging.Error(err).Send()
		return exports.RaiseAPIErrorAuto(exports.APULIS_IQI_DB_UPDATE_UNEXPECT, err.Error())
	}
	return nil
}

func (svc AnnotationDataDao) SaveOrUpdate(c context.Context, model *models.AnnotationData) exports.APIError {
	tx := GetDB(c)
	model.BaseModelTime.UpdatedAt = utils.GetNowTime()
	err := tx.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "type"}, {Name: "name"}, {Name: "user_id"}},
		DoUpdates: clause.AssignmentColumns([]string{"config", "updated_at"}),
	}).Create(model).Error
	if err != nil {
		logging.Error(err).Send()
		return exports.RaiseAPIErrorAuto(exports.APULIS_IQI_DB_EXEC_FAILED, err.Error())
	}
	return nil
}
func (svc AnnotationDataDao) BatchSaveOrUpdate(c context.Context, model []models.AnnotationData) exports.APIError {
	tx := GetDB(c)
	for _, m := range model {
		m.UpdatedAt = utils.GetNowTime()
	}
	err := tx.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "type"}, {Name: "name"}, {Name: "user_id"}},
		DoUpdates: clause.AssignmentColumns([]string{"config", "updated_at"}),
	}).Create(model).Error
	if err != nil {
		logging.Error(err).Send()
		return exports.RaiseAPIErrorAuto(exports.APULIS_IQI_DB_EXEC_FAILED, err.Error())
	}
	return nil
}
