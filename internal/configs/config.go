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
	"github.com/fsnotify/fsnotify"
	"github.com/go-labs/internal/logging"
	"github.com/spf13/viper"
)

var appConfig *AppConfig

type AppConfig struct {
	Port  int
	Grpc  int
	Debug bool
}

func InitConfig() (*AppConfig, error) {
	logging.Info().Msg("reading config")
	v := viper.New()
	v.SetConfigName("config")
	v.AddConfigPath("configs")

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
	return appConfig, nil
}

func GetAppConfig() *AppConfig {
	return appConfig
}
