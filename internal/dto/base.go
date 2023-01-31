package dto

import "github.com/go-labs/internal/models"

type BaseListDto struct {
	PageNum  int    `form:"pageNum" json:"pageNum"`
	PageSize int    `form:"pageSize" json:"pageSize"`
	Order    string `form:"order" json:"order"`
	OrderBy  string `form:"orderBy" json:"orderBy"`
	Sort     string `form:"sort"`
}

func (b *BaseListDto) Conversion() *models.BaseList {
	return &models.BaseList{
		PageNum:  b.PageNum,
		PageSize: b.PageSize,
		Sort:     b.Sort,
	}
}

type BaseListResp struct {
	Items interface{} `json:"items"`
	Total int64       `json:"total"`
}
