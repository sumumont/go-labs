package dto

type RecheckRsp struct {
	Id           int64        `json:"id"`
	ImageName    string       `json:"imageName"`
	ConnectorId  int64        `json:"connectorId"`
	RecheckFlag  string       `json:"recheckFlag"`
	Score        float64      `json:"score"`
	RecheckBoxes []RecheckBox `json:"recheckBoxes"`
}

type RecheckBox struct {
	Id          int64             `json:"id"`
	BoxType     string            `json:"boxType"`     // infer推理框 annotation标注框
	RecheckFlag string            `json:"recheckFlag"` // 复判标记
	Score       float64           `json:"score"`       // 置信度
	Box         interface{}       `json:"box"`         // 矩形框的信息 复判框的详情
	Meta        interface{}       `json:"meta"`        // 额外的扩展信息
	Labels      []RecheckBoxLable `json:"labels"`
}

type RecheckBoxLable struct {
	Label     string      `json:"label"`     // 标注的标签
	AnnotType string      `json:"annotType"` // 标注类型，detection, classfication
	Meta      interface{} `json:"meta"`      // 额外的扩展信息
}
