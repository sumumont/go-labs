package services

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/go-labs/internal/dao"
	"github.com/go-labs/internal/dto"
	"github.com/go-labs/internal/logging"
	"github.com/go-labs/internal/models"
	"github.com/go-labs/pkg/exports"
)

type AnnotationDataService struct {
}

var AnnotationDataSvc = &AnnotationDataService{}

func GetAnnotationDataSvc() *AnnotationDataService {
	return AnnotationDataSvc
}

func (svc *AnnotationDataService) Save(c *gin.Context, param *dto.AnnotationData) (interface{}, exports.APIError) {
	errs := dao.ExecDBTx(func(ctx context.Context) error {
		AnnotationData := &models.AnnotationData{
			//todo fill Private
			AnnotationTaskId:   param.AnnotationTaskId,
			AttributeKey:       param.AttributeKey,
			AttributeValue:     param.AttributeValue,
			Id:                 param.Id,
			Status:             param.Status,
			VirtualDatasetId:   param.VirtualDatasetId,
			VirtualDatasetName: param.VirtualDatasetName,
		}
		err := dao.GetAnnotationDataDao().SaveOrUpdate(ctx, AnnotationData)
		if err != nil {
			logging.Error(err).Send()
			return err
		}
		return nil
	})
	if errs != nil {
		logging.Error(errs).Send()
		return nil, exports.WrapperError(errs)
	}
	return nil, nil
}
func (svc *AnnotationDataService) Delete(c *gin.Context, param dto.AnnotationDataId) (interface{}, exports.APIError) {
	errs := dao.ExecDBTx(func(ctx context.Context) error {
		err := dao.GetAnnotationDataDao().Delete(ctx, param.Id)
		if err != nil {
			logging.Error(err).Send()
			return err
		}
		return nil
	})
	if errs != nil {
		logging.Error(errs).Send()
		return nil, exports.WrapperError(errs)
	}
	return nil, nil
}
func (svc *AnnotationDataService) Find(c *gin.Context, param dto.AnnotationDataParam) (interface{}, exports.APIError) {
	AnnotationDatas, total, err := dao.GetAnnotationDataDao().FindByParam(c, param, nil)
	if err != nil {
		logging.Error(err).Send()
		return nil, err
	}
	resp := &dto.BaseListResp{}
	resp.Items = AnnotationDatas
	resp.Total = total
	return resp, nil
}
func (svc *AnnotationDataService) Detail(c *gin.Context, param dto.AnnotationDataId) (interface{}, exports.APIError) {
	item, err := dao.GetAnnotationDataDao().FindById(c, param.Id)
	if err != nil {
		logging.Error(err).Send()
		return nil, err
	}
	return item, nil
}
