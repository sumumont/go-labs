package dto

type AnnotationData struct {
	AnnotationTaskId   int64  `json:"annotationTaskId"`
	AttributeKey       string `json:"attributeKey"`
	AttributeValue     string `json:"attributeValue"`
	Id                 int64  `json:"id"`
	Status             string `json:"status"`
	VirtualDatasetId   int64  `json:"virtualDatasetId"`
	VirtualDatasetName string `json:"virtualDatasetName"`
}

type AnnotationDataId struct {
	Id int64 `json:"id" uri:"id"`
}

type AnnotationDataParam struct {
	BaseListDto
	AnnotationTaskId   int64  `json:"annotationTaskId" form:"annotationTaskId"`
	AttributeKey       string `json:"attributeKey" form:"attributeKey"`
	AttributeValue     string `json:"attributeValue" form:"attributeValue"`
	Id                 int64  `json:"id" form:"id"`
	Status             string `json:"status" form:"status"`
	VirtualDatasetId   int64  `json:"virtualDatasetId" form:"virtualDatasetId"`
	VirtualDatasetName string `json:"virtualDatasetName" form:"virtualDatasetName"`
}
