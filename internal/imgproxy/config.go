package imageproxy

import (
	"fmt"
	"strings"
)

type ImgProxyConfig struct {
	Enctypted bool
	Key       string
	Salt      string

	StorageType   string
	RealUrlPrefix string
}

const storageTypeLocal = "local"
const storageTypeS3 = "s3"
const storageTypeGS = "gs"
const storageTypeABS = "abs"
const storageTypeExternalHttp = "http"
const storageTypeExternalHttps = "https"

var storageTypes = map[string]bool{
	storageTypeLocal:         true,
	storageTypeS3:            true,
	storageTypeGS:            true,
	storageTypeABS:           true,
	storageTypeExternalHttp:  true,
	storageTypeExternalHttps: true,
}

func (imgProxyConfig *ImgProxyConfig) CheckImgProxyConfigValid() error {
	if !storageTypes[imgProxyConfig.StorageType] {
		return fmt.Errorf("Storage type %s not supported ...", imgProxyConfig.StorageType)
	}

	if !strings.HasPrefix(imgProxyConfig.RealUrlPrefix, "http") {
		return fmt.Errorf("Real url prefix not valid: %s", imgProxyConfig.RealUrlPrefix)
	}

	imgProxyConfig.RealUrlPrefix = strings.TrimSuffix(imgProxyConfig.RealUrlPrefix, "/")
	return nil
}
