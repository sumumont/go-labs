package models

type Product struct {
	BaseModelId
	Num int64 `json:"num"`
	UserInfo
	BaseModelTime
}
