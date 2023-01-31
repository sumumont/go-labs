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

package models

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

// UnixTime For gorm time storage
type UnixTime struct {
	time.Time
}

// For gorm jsonb storage
type JsonB map[string]interface{}

// UnixTime implement gorm interfaces
func (t UnixTime) MarshalJSON() ([]byte, error) {
	if t.IsZero() {
		return []byte("null"), nil
	}
	microSec := t.UnixNano() / int64(time.Millisecond)
	return []byte(strconv.FormatInt(microSec, 10)), nil
}

func (t UnixTime) Value() (driver.Value, error) {
	//var zeroTime time.Time
	//if t.Time.UnixNano() == zeroTime.UnixNano() {
	//	return nil, nil
	//}
	if t.IsZero() {
		return nil, nil
	}
	return t.Time, nil
}

func (t *UnixTime) Scan(v interface{}) error {
	value, ok := v.(time.Time)
	if ok {
		*t = UnixTime{Time: value}
		return nil
	}
	return fmt.Errorf("cannot convert %v to timestamp", v)
}

// JsonB implement gorm interfaces
func (j JsonB) Value() (driver.Value, error) {
	valueStr, err := json.Marshal(j)
	return string(valueStr), err
}

func (j *JsonB) Scan(value interface{}) error {
	if value == nil {
		return nil
	}
	return json.Unmarshal(value.([]byte), &j)
}

type JsonMetaData struct {
	//data map[string]interface{}
	data_str []byte
}

func (d *JsonMetaData) Empty() bool {
	return len(d.data_str) <= 2
}
func (d *JsonMetaData) MarshalJSON() ([]byte, error) {
	if len(d.data_str) == 0 {
		return []byte("null"), nil
	}
	return d.data_str, nil
}
func (d *JsonMetaData) UnmarshalJSON(b []byte) error {
	if len(b) >= 2 && (b[0] == '{' || b[0] == '[') {
		d.data_str = b
	} else {
		d.data_str = nil
	}
	return nil
}

func (d JsonMetaData) Value() (driver.Value, error) {
	if len(d.data_str) == 0 {
		return nil, nil
	}

	return d.data_str, nil
}

func (d *JsonMetaData) Scan(v interface{}) error {
	switch ty := v.(type) {
	case string:
		d.data_str = []byte(ty)
	case []byte:
		d.data_str = ty
	default:
		d.data_str = nil
	}
	return nil
}

func (d *JsonMetaData) Fetch(v interface{}) error {
	if d == nil {
		return nil
	}
	return json.Unmarshal([]byte(d.data_str), v)
}
func (d *JsonMetaData) Save(v interface{}) {
	d.data_str, _ = json.Marshal(v)
}

// GormDBDataType gorm db data type
func (JsonMetaData) GormDBDataType(db *gorm.DB, _ *schema.Field) string {
	switch db.Dialector.Name() {
	case "sqlite":
		return "JSON"
	case "mysql":
		return "JSON"
	case "postgres":
		return "JSONB"
	}
	return ""
}

func assertCheck(v bool, msg string) {
	if !v {
		panic(msg)
	}
}

type StringJsonB []string

func (s StringJsonB) Value() (driver.Value, error) {
	marshal, err := json.Marshal(s)
	return string(marshal), err
}

func (s *StringJsonB) Scan(value interface{}) error {
	if value == nil {
		return nil
	}
	err := json.Unmarshal(value.([]byte), &s)
	return err
}
