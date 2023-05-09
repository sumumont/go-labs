package models

type AnnotationData struct {
	Id                 int64  `json:"id" gorm:"primary_key"`
	AnnotationTaskId   int64  `json:"annotationTaskId" gorm:"default:0;uniqueIndex:idx_annotation_data,priority:1"`
	VirtualDatasetId   int64  `json:"virtualDatasetId" gorm:"uniqueIndex:idx_annotation_data,priority:2"`
	VirtualDatasetName string `json:"virtualDatasetName"`
	AttributeKey       string `json:"attributeKey" gorm:"default:imageName;uniqueIndex:idx_annotation_data,priority:3" ` //默认是图片名字段，作为关联数据的字段，也有可能是其他的
	AttributeValue     string `json:"attributeValue" gorm:"uniqueIndex:idx_annotation_data,priority:4"`
	Status             string `json:"status" gorm:"type:varchar(30)"` //distributing:待分发,labeling标注中，labeled已标注
	UserInfo
	BaseModelTime
}
