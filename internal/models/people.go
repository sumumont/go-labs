package models

type People struct {
	ID   int64  `aorm:"primary_key" gorm:"primary_key" json:"id"`
	Name string `gorm:"type:varchar(64);" json:"Name"`
	UserInfo
	BaseModelTime
	Infos []Info `json:"info" aorm:"type:child;"`
}
type Info struct {
	PeopleId int64 `aorm:"people_id" gorm:"column:__info;type:varchar(64);" json:"peopleId"`
	Age      int   `gorm:"type:varchar(64);" json:"age"`
}
