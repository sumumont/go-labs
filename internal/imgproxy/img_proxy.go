package imageproxy

import (
	"github.com/go-labs/internal/configs"
	"github.com/go-labs/internal/logging"
)

var Client ImgProxyClient

func loadconfig() {
	config := configs.GetAppConfig().ImgProxy

	var imgProxyConfig ImgProxyConfig
	imgProxyConfig.Enctypted = config.Enctypted
	imgProxyConfig.Key = config.Key
	imgProxyConfig.Salt = config.Salt
	imgProxyConfig.StorageType = config.StorageType
	imgProxyConfig.RealUrlPrefix = config.RealUrlPrefix

	client, err := CreateImgProxyClient(&imgProxyConfig)
	if err != nil {
		logging.Error(err).Send()
		panic(err)
	}
	Client = *client
	logging.Info().Interface("imageproxyConfig", imgProxyConfig).Msg("Create imageproxy client success")
}
