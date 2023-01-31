package models

type UserConfig struct {
	BaseModelId
	Type      string       `gorm:"type:varchar(64);uniqueIndex:uIdx_user_config,priority:1" json:"type"`
	Name      string       `gorm:"type:varchar(64);uniqueIndex:uIdx_user_config,priority:2" json:"Name"`
	Config    JsonMetaData `gorm:"type:jsonb" json:"config"`
	UserName  string       `gorm:"type:varchar(64)" json:"userName"`
	OrgName   string       `gorm:"type:varchar(64)" json:"orgName"`
	GroupName string       `gorm:"type:varchar(64)" json:"groupName"`
	UserId    int64        `gorm:"userId;uniqueIndex:uIdx_user_config,priority:3" json:"userId"`
	GroupId   int64        `gorm:"groupId" json:"groupId"`
	OrgId     int64        `gorm:"orgId" json:"orgId"`
	BaseModelTime
}

func (u *UserConfig) SetUser(userInfo UserInfo) {
	u.UserId = userInfo.UserId
	u.UserName = userInfo.UserName
	u.GroupId = userInfo.GroupId
	u.GroupName = userInfo.GroupName
	u.OrgId = userInfo.OrgId
	u.OrgName = userInfo.OrgName
}
