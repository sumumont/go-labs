package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/go-labs/pkg/exports"
)

type BaseController struct{}

func Bind(c *gin.Context, params interface{}) exports.APIError {
	_ = c.ShouldBindUri(params)
	if err := c.ShouldBind(params); err != nil {
		return exports.RaiseAPIErrorAuto(exports.APULIS_IQI_PARAM_ERROR, err.Error())
	}
	return nil
}
