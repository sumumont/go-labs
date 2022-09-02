package dto

import (
	"github.com/go-labs/internal/logging"
	"testing"
)

func TestRecheck(t *testing.T) {
	recheckRsp := RecheckRsp{
		Id:          0,
		ImageName:   "",
		ConnectorId: 0,
		RecheckFlag: "",
		Score:       0,
		RecheckBoxes: []RecheckBox{
			RecheckBox{
				Id:          0,
				BoxType:     "",
				RecheckFlag: "",
				Score:       0,
				Box:         nil,
				Meta:        nil,
				Labels: []RecheckBoxLable{
					{
						Label:     "",
						AnnotType: "",
						Meta:      nil,
					},
				},
			},
		},
	}
	logging.Debug().Interface("recheckRsp", recheckRsp).Send()
}
