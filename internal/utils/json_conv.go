package utils

import "encoding/json"

// JsonConv json格式化转换
func JsonConv(origin, out interface{}) error {

	buffer, err := json.Marshal(origin)
	if err != nil {
		return err
	}

	err = json.Unmarshal(buffer, &out)
	if err != nil {
		return err
	}

	return nil
}
