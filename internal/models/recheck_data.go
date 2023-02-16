package models

type RecheckData struct {
	Id            string   `json:"id"`
	Label         string   `json:"label"`
	DataChannelId string   `json:"dataChannelId"`
	InferId       string   `json:"inferId"`
	ImageName     string   `json:"imageName"`
	DeviceName    string   `json:"deviceName"`
	FovID         float64  `json:"fovID"`
	BoardID       float64  `json:"boardID"`
	RecheckResult string   `json:"recheckResult"`
	Boxes         []Box    `json:"boxes"`
	Location      Location `json:"location"`
}

type Box struct {
	Id            string   `json:"id"`
	Result        string   `json:"result"`
	Score         float64  `json:"score"`
	RecheckResult string   `json:"recheckResult"`
	Location      Location `json:"location"`
	Label         string   `json:"label"`
}
type Location struct {
	X      float64 `json:"X"`
	Y      float64 `json:"Y"`
	Width  float64 `json:"Width"`
	Height float64 `json:"Height"`
}
