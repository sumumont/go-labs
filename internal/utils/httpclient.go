/* ******************************************************************************
* 2019 - present Contributed by Apulis Technology (Shenzhen) Co. LTD
*
* This program and the accompanying materials are made available under the
* terms of the MIT License, which is available at
* https://www.opensource.org/licenses/MIT
*
* See the NOTICE file distributed with this work for additional
* information regarding copyright ownership.
* Unless required by applicable law or agreed to in writing, software
* distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
* WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the
* License for the specific language governing permissions and limitations
* under the License.
*
* SPDX-License-Identifier: MIT
******************************************************************************/
package utils

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"github.com/go-labs/internal/logging"
	"io"
	"io/ioutil"
	logger "log"
	"net/http"
	"strconv"
	"strings"
	"sync"
	"time"
)

var httpClient *http.Client

type APISuccessRsp struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

type InnerError struct {
	InnerCode      int
	HttpStatusCode int
	ErrorMsg       string
	Data           map[string]string
}

func (e *InnerError) Error() string {
	return e.ErrorMsg
}

func (e *InnerError) ErrorCode() int {
	return e.InnerCode
}

func (e *InnerError) StatusCode() int {
	return e.HttpStatusCode
}

func (e *InnerError) ErrorData() map[string]string {
	return e.Data
}

const (
	UNABLE_lOCALIZE = "unable to localize msg"
)

// define color log output
const (
	InfoColor    = "\033[1;34m%v\033[0m"
	NoticeColor  = "\033[1;36m%v\033[0m"
	WarningColor = "\033[1;33m%v\033[0m"
	ErrorColor   = "\033[1;31m%v\033[0m"
	DebugColor   = "\033[0;36m%v\033[0m"
)

func GetHttpUrl(host string, path string) string {
	return "http://" + host + path
}

func InitHttpClient() error {
	transport := http.DefaultTransport.(*http.Transport).Clone()
	transport.MaxIdleConns = 100
	transport.MaxConnsPerHost = 100
	transport.MaxIdleConnsPerHost = 100
	transport.TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	httpClient = &http.Client{
		Timeout:   time.Duration(10) * time.Second,
		Transport: transport,
	}
	return nil
}

var one sync.Once

func GetHttpClient() *http.Client {
	if httpClient == nil {
		one.Do(func() {
			_ = InitHttpClient()
		})
	}
	return httpClient
}

func doHttpRequest(url, method string, headers map[string]string, rawBody interface{}) ([]byte, error) {
	var body io.Reader
	// logger.Printf("request [%v] %s", method, url)
	if rawBody != nil {
		if headers == nil {
			headers = make(map[string]string)
		}
		headers["Content-Type"] = "application/json"

		switch t := rawBody.(type) {
		case string:
			body = strings.NewReader(t)
			// logger.Printf("req json: %s", t)
		case []byte:
			body = bytes.NewReader(t)
			// logger.Printf("req json: %s", string(t))
		default:
			data, err := json.Marshal(rawBody)
			if err != nil {
				return nil, err
			}
			// logger.Printf("req json: %s", string(data))
			body = bytes.NewReader(data)
		}
	}
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, err
	}

	if len(headers) != 0 {
		// logger.Printf("req header: %+v", headers)
		for k, v := range headers {
			req.Header.Set(k, v)
		}
	}

	resp, err := GetHttpClient().Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if responseData, err := ioutil.ReadAll(resp.Body); err != nil {
		// logger.Printf("request [%v] [%s] read resp.Body err: %s", method, url, err.Error())
		return nil, err
	} else if resp.StatusCode != http.StatusOK {
		// logger.Printf("req rsp statusCode %d err: %s", resp.StatusCode, string(responseData))
		return responseData, err
	} else {
		// logger.Printf("req rsp: %s", string(responseData))
		return responseData, nil
	}
}

// return body & error code
func AIStudioRequest(url, method string, headers map[string]string, rawBody interface{}, output interface{}) (
	*APISuccessRsp, error) {
	rspData, err := doHttpRequest(url, method, headers, rawBody)
	if len(rspData) == 0 { // no data return
		logger.Printf("[%s] %s http failed: %v", CErr(method), url, err)
		return nil, err
	}
	response := APISuccessRsp{ //Data: output,
	}
	if err := json.Unmarshal(rspData, &response); err != nil {
		logging.Error(err).Str("method", method).Str("url", url).Interface("rawBody", rawBody).Interface("rspData", string(rspData)).Send()
		return nil, err
	}
	if response.Code != 0 {
		logger.Printf("[%s] %s code: %v with req:%v error response: %v",
			CErr(method), url, response.Code, rawBody, response)
		m, _ := response.Data.(map[string]string)
		err = &InnerError{
			InnerCode:      response.Code,
			HttpStatusCode: http.StatusBadRequest,
			ErrorMsg:       response.Msg,
			Data:           m,
		}
		return &response, err
	} else {
		if output != nil {
			if err := JsonConv(response.Data, output); err != nil {
				logging.Error(err).Str("method", method).Str("url", url).Interface("rawBody", rawBody).Interface("rspData", string(rspData)).Send()
				return nil, err
			}
		}
		logger.Printf("[%s] %s succeed", CDbg(method), url)
		response.Data = output
		return &response, nil
	}
}
func DoRequest(url, method string, headers map[string]string, rawBody interface{}, output interface{}) error {

	rspData, err := doHttpRequest(url, method, headers, rawBody)
	if len(rspData) == 0 { // no data return
		return err
	}
	//@todo: only accpet json response ???
	if output != nil {
		if e := json.Unmarshal(rspData, output); e != nil {
			logger.Printf("[%s] %s with req:%v invalid response json data, err: %v", CErr(method), url, rawBody, err)
			return err
		}
	}
	if err != nil {
		logger.Printf("[%s] :%s with req:%v error response:%v", method, url, rawBody, string(rspData))
		if len(headers) > 0 {
			logger.Printf("with headers:%v", headers)
		}
	}

	return err
}
func GetLangCode(code int) string {
	return strconv.Itoa(code)
}

// CErr color output
func CErr(s interface{}) string {
	return fmt.Sprintf(ErrorColor, s)
}

// CDbg color output
func CDbg(s interface{}) string {
	return fmt.Sprintf(DebugColor, s)
}
