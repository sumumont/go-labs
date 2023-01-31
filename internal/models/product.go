package models

type Product struct {
	BaseModelId
	Num        int64 `json:"num"`
	Attributes JsonB `gorm:"type:jsonb" json:"attributes"`
	UserInfo
	BaseModelTime
}
