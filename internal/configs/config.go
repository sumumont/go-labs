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
package configs

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/go-labs/internal/logging"
	"github.com/spf13/viper"
)

var appConfig *AppConfig

type AppConfig struct {
	Port     int            `json:"port"`
	Grpc     int            `json:"grpc"`
	Debug    bool           `json:"debug"`
	ImgProxy ImgProxyConfig `json:"imgProxy"`
}

type ImgProxyConfig struct {
	Enctypted     bool
	Key           string
	Salt          string
	StorageType   string
	RealUrlPrefix string
}

func InitConfig(path string) (*AppConfig, error) {
	logging.Info().Msg("reading config")
	v := viper.NewWithOptions(viper.KeyDelimiter("::"))
	v.SetConfigName("config")
	v.AddConfigPath(path)

	err := v.ReadInConfig()
	if err != nil {
		return nil, err
	}
	appConfig = &AppConfig{}
	v.WatchConfig()
	v.OnConfigChange(func(in fsnotify.Event) {
		logging.Info().Str("fs.event", in.Name).Msg("config.yaml changed")
		err := v.Unmarshal(&appConfig)
		if err != nil {
			logging.Error(err).Send()
		}
	})
	err = v.Unmarshal(&appConfig)
	if err != nil {
		return nil, err
	}
	a := appConfig
	fmt.Println(a)
	//
	//var b AppConfig
	//yfile, _ := os.Open("./configs/config.yaml")	//test.yaml由下一个例子生成
	//defer yfile.Close()
	//ydecode:= yaml.NewDecoder(yfile)
	//ydecode.Decode(&b)
	//fmt.Println(b.Spark.AppName)
	return appConfig, nil
}

func GetAppConfig() *AppConfig {
	return appConfig
}
