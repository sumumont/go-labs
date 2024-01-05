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
package routers

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/go-labs/internal/configs"
	"github.com/go-labs/internal/logging"
)

func InitRouter() *gin.Engine {

	if !configs.GetAppConfig().Debug {
		gin.SetMode(gin.ReleaseMode)
	} else {
		gin.SetMode(gin.DebugMode)
	}
	//@todo: init common user identity authentication logic here
	r := gin.New()

	r.MaxMultipartMemory = 8 << 20 // 8 MiB

	r.Use(cors.Default())
	r.RedirectTrailingSlash = false
	//r.Use(Auth())

	//r.NoMethod(HandleNotFound)
	//r.NoRoute(HandleNotFound)

	r.Use(gin.Recovery())
	r.Use(logging.Middleware())
	r.POST("/test", func(c *gin.Context) {
		var req interface{}
		c.ShouldBindJSON(&req)
		logging.Info().Msg(c.Request.URL.RawQuery)
	})
	r.POST("/resource/*filepath", func(c *gin.Context) {
		var req interface{}
		c.ShouldBindJSON(&req)
		logging.Info().Str("filepath", c.Query("filepath")).Msg(c.Request.URL.RawQuery)
	})
	return r
}
