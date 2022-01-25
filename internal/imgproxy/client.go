package imageproxy

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"net/url"
	"strings"
)

const insecureValue = "insecure"

type ImgProxyClient struct {
	config    *ImgProxyConfig
	keyBytes  []byte
	saltBytes []byte
}

type imgProxyOperation struct {
	client     *ImgProxyClient
	operations []string
}

type imgProxyPathToSign struct {
	client     *ImgProxyClient
	operations []string
	filePath   string
}

func CreateImgProxyClient(imgProxyConfig *ImgProxyConfig) (*ImgProxyClient, error) {
	err := imgProxyConfig.CheckImgProxyConfigValid()
	if err != nil {
		return nil, err
	}

	keyBytes, err := hex.DecodeString(imgProxyConfig.Key)
	if err != nil {
		return nil, err
	}
	saltBytes, err := hex.DecodeString(imgProxyConfig.Salt)
	if err != nil {
		return nil, err
	}

	return &ImgProxyClient{
		config:    imgProxyConfig,
		keyBytes:  keyBytes,
		saltBytes: saltBytes,
	}, nil
}

func (imgProxyClient *ImgProxyClient) NewOperation() *imgProxyOperation {
	var operations []string
	return &imgProxyOperation{client: imgProxyClient, operations: operations}
}

func (imgProxyOperation *imgProxyOperation) Resize(resizeType ResizeType, width, height int, enlarge bool) *imgProxyOperation {
	resizeStr := fmt.Sprintf("rs:%s:%d:%d:%t", resizeType, width, height, enlarge)
	imgProxyOperation.operations = append(imgProxyOperation.operations, resizeStr)
	return imgProxyOperation
}

func (imgProxyOperation *imgProxyOperation) Quality(quality int) *imgProxyOperation {
	qualityStr := fmt.Sprintf("q:%d", quality)
	imgProxyOperation.operations = append(imgProxyOperation.operations, qualityStr)
	return imgProxyOperation
}

func (imgProxyOperation *imgProxyOperation) AddFilePath(filePath string) *imgProxyPathToSign {
	return &imgProxyPathToSign{
		client:     imgProxyOperation.client,
		operations: imgProxyOperation.operations,
		filePath:   filePath,
	}
}

func (imgProxyPathToSign *imgProxyPathToSign) GenerateURL() string {
	operationPath := strings.Join(imgProxyPathToSign.operations, "/")
	fileUrl := encodeFileURL(imgProxyPathToSign.client.config.StorageType, imgProxyPathToSign.filePath)
	pathToSign := fmt.Sprintf("/%s/%s", operationPath, fileUrl)
	sign := imgProxyPathToSign.client.signForPath(pathToSign)
	return fmt.Sprintf("%s/%s%s", imgProxyPathToSign.client.config.RealUrlPrefix, sign, pathToSign)
}

func encodeFileURL(storageType, filePath string) string {
	filePath = strings.TrimPrefix(filePath, "/")
	sourceUrl := fmt.Sprintf("%s:///%s", storageType, url.PathEscape(filePath))
	return base64.RawURLEncoding.EncodeToString([]byte(sourceUrl))
}

func (imgProxyClient *ImgProxyClient) signForPath(path string) string {
	if !imgProxyClient.config.Enctypted {
		return insecureValue
	}

	mac := hmac.New(sha256.New, imgProxyClient.keyBytes)
	_, _ = mac.Write(imgProxyClient.saltBytes)
	_, _ = mac.Write([]byte(path))
	return base64.RawURLEncoding.EncodeToString(mac.Sum(nil))
}
