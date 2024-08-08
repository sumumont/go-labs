package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/go-labs/internal/dto"
	"github.com/go-labs/internal/logging"
	"github.com/go-labs/internal/services"
	"github.com/go-labs/pkg/exports"
)

type AnnotationDataCtl struct {
	BaseController
}

func NewAnnotationDataCtl() *AnnotationDataCtl {
	return &AnnotationDataCtl{}
}

func (p *AnnotationDataCtl) Save(c *gin.Context) (interface{}, exports.APIError) {
	var param dto.AnnotationData
	err := Bind(c, &param)
	if err != nil {
		return nil, err
	}
	rsp, err := services.GetAnnotationDataSvc().Save(c, &param)
	if err != nil {
		logging.Error(err).Send()
		return nil, err
	}
	return rsp, nil

}
func (p *AnnotationDataCtl) FindAll(c *gin.Context) (interface{}, exports.APIError) {
	var param dto.AnnotationDataParam
	err := Bind(c, &param)
	if err != nil {
		return nil, err
	}
	rsp, err := services.GetAnnotationDataSvc().Find(c, param)
	if err != nil {
		logging.Error(err).Send()
		return nil, err
	}
	return rsp, nil
}
func (p *AnnotationDataCtl) Delete(c *gin.Context) (interface{}, exports.APIError) {
	var param dto.AnnotationDataId
	err := Bind(c, &param)
	if err != nil {
		return nil, err
	}
	rsp, err := services.GetAnnotationDataSvc().Delete(c, param)
	if err != nil {
		logging.Error(err).Send()
		return nil, err
	}
	return rsp, nil
}
