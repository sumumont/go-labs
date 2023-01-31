package models

type People struct {
	BaseModelId
	Name string `gorm:"type:varchar(64);" json:"Name"`
	UserInfo
	BaseModelTime
}
