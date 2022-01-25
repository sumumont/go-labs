package imageproxy

import (
	"fmt"
	"github.com/go-labs/internal/configs"
	"github.com/go-labs/internal/logging"
	"testing"
)

func TestImgProxy(t *testing.T) {
	config, err := configs.InitConfig("../../configs")
	if err != nil {
		panic(err)
	}
	logging.Info().Interface("config", config).Send()
	loadconfig()
	url := "dataset/ipc/2022012421/qijian_1643029507.jpg"
	realUrl := Client.NewOperation().Resize(ResizeTypeFit, 6000, 6000, true).AddFilePath("/" + url).GenerateURL()
	logging.Info().Msg(realUrl)
	fmt.Println(realUrl)
}
