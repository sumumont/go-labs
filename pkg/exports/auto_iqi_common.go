/*******************************************************************************
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

package exports

import "fmt"

const APULIS_IQI_MODULE_ID = 1022
const APULIS_IQI_API_VERSION = "/api/v1"
const APULIS_IQI_MODULE_NAME = "apulis-iqi"

// Define HTTP rest service used client structures
type RequestObject = map[string]interface{}
type RespObject = map[string]interface{}
type GObject = map[string]interface{}
type QueryFilterMap = map[string]string
type RequestTags = map[string]string

type SearchCond struct {
	Offset     uint
	TotalCount int64
	Next       string
	//start from 1~N
	PageNum  uint   `form:"pageNum"`
	PageSize uint   `form:"pageSize"`
	Sort     string `form:"sort"`
	// list by app group
	Group string `form:"group"`
	// indicate "group" list match recursively !
	MatchAll bool `form:"matchAll"`
	// search by keyword
	SearchWord string `form:"searchWord"`
	//enumeration for need detail return
	Detail int32 `form:"detail"`
	//enumeration for deleted item search
	Show int32 `form:"show"`
	// filters by predefined key=value pairs
	EqualFilters map[string]string
	// filters by advacned operator
	AdvanceOpFilters map[string]interface{}
}

type CommResponse struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

type PagedResult struct {
	// total items matched
	Total uint `json:"total"`
	// ceil(total/pageSize) ,be zero if request pageSize is zero
	TotalPages uint `json:"totalPages,omitempty"`
	// request pageNum
	PageNum uint `json:"pageNum,omitempty"`
	// request pageSize, if zero indicate none paged querys
	PageSize uint `json:"pageSize,omitempty"`
	// used for next pagedQuery hints
	Next string `json:"next,omitempty"`
	// used for return data
	Items interface{} `json:"items"`
}

func MakeIQIProjectContext(projectId uint64) string {
	return fmt.Sprintf("%s/%d", APULIS_IQI_MODULE_NAME, projectId)
}
