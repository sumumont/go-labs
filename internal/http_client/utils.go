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
package http_client

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"github.com/apulisai/sdk/go-utils/logging"
	"github.com/go-labs/pkg/exports"
	"io"
	"io/ioutil"
	"net/http"
)

func DoRequest(url, method string, headers http.Header, body io.Reader) (*http.Response, error) {
	var req *http.Request = nil
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, err
	}
	if headers != nil {
		req.Header = headers.Clone()
	}
	client := &http.Client{}
	client.Transport = &http.Transport{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: true,
		},
	}
	return client.Do(req)
}

func ParseResponse(resp *http.Response, data interface{}) error {
	defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)
	logging.Debug().Str("ParseResponse", string(b)).Send()
	if err != nil {
		return err
	}
	commResp := &exports.CommResponse{
		Data: data,
	}
	err = json.Unmarshal(b, commResp)
	if err == nil {
		if commResp.Code == 0 {
			return nil
		}
		return fmt.Errorf("failed")
	}
	logging.Error(err).Send()
	fmt.Println("json.Unmarshal err", err)
	return fmt.Errorf("http request err:  %d msg: %s", resp.StatusCode, string(b))
}

func Request(url, method string, headers http.Header, param interface{}, data interface{}) error {
	logging.Debug().Str("request.url", url).Send()
	var body io.Reader
	var b []byte
	if param == nil {
		body = nil
	} else {
		b, _ = json.Marshal(param)
		logging.Debug().Interface("http.param", param).Send()
		body = bytes.NewReader(b)
	}
	rsp, err := DoRequest(url, method, headers, body)
	if err != nil {
		return err
	}
	return ParseResponse(rsp, data)
}
