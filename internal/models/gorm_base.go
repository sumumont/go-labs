package models

import (
	"fmt"
	"gorm.io/gorm"
)

type BaseModel struct {
	ID        int64 `gorm:"primaryKey;autoIncrement" json:"id"`
	CreatedAt int64 `gorm:"autoCreateTime:milli" json:"createdAt"`
	UpdatedAt int64 `gorm:"autoUpdateTime:milli" json:"updatedAt"`
}
type BaseModelId struct {
	ID int64 `gorm:"primary_key" json:"id"`
}
type BaseModelTime struct {
	CreatedAt int64 `gorm:"autoCreateTime:milli" json:"createdAt"`
	UpdatedAt int64 `gorm:"autoUpdateTime:milli" json:"updatedAt"`
}
type UserInfo struct {
	UserName  string `gorm:"userName;varchar(64)" json:"userName"`
	UserId    int64  `gorm:"userId" json:"userId"`
	GroupName string `gorm:"groupName;varchar(64)" json:"groupName"`
	GroupId   int64  `gorm:"groupId" json:"groupId"`
	OrgId     int64  `gorm:"orgId" json:"orgId"`
	OrgName   string `gorm:"orgName;varchar(64)" json:"orgName"`
}

type BaseList struct {
	PageNum  int    `json:"pageNum"`
	PageSize int    `json:"pageSize"`
	Sort     string `json:"sort"`
}

func NormalizeSorts(str string) (string, error) {
	var sorts string
	for _, v := range str {
		if v >= 'A' && v <= 'Z' {
			sorts += "_" + string(v+32)
		} else if v == '|' {
			sorts += " "
		} else if v >= 'a' && v <= 'z' || v == '_' || v == '-' || v == ',' {
			sorts += string(v)
		} else {
			return "", fmt.Errorf("invalid sort fields!")
		}
	}
	return sorts, nil
}

func PageList(db *gorm.DB, baseList *BaseList) (*gorm.DB, error) {
	if baseList != nil {
		sort, err := NormalizeSorts(baseList.Sort)
		if err != nil {
			return nil, err
		}
		if sort == "" {
			sort = "created_at desc"
		}
		if len(sort) > 0 {
			db = db.Order(sort)
		}
		if baseList.PageNum > 0 && baseList.PageSize > 0 {
			db = db.Offset((baseList.PageNum - 1) * baseList.PageSize).Limit(baseList.PageSize)
		}
	}
	return db, nil
}
