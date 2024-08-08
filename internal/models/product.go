package models

type Product struct {
	BaseModelId
	Num        int64 `json:"num"`
	Attributes JsonB `gorm:"type:jsonb" json:"attributes"`
	UserInfo
	BaseModelTime
}

type School struct {
	ID   int64  `gorm:"primary_key" json:"id"`
	Name string `json:"name"`
}
type TheClass struct {
	ID       int64     `gorm:"primary_key" json:"id"`
	Name     string    `json:"name"`
	Students []Student `gorm:"foreignKey:TheClassId" json:"schoolId"`
}
type Student struct {
	ID         int64  `gorm:"primary_key" json:"id"`
	Name       string `json:"name"`
	SchoolId   int64  `json:"schoolId"`
	TheClassId int64  `json:"theClassId"`
}
