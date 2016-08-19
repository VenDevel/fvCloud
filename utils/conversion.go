package utils

import (
	"encoding/json"
)

func ByteJsonToMap(body []byte) (map[string]interface{}, error) {
	var da map[string]interface{}
	err := json.Unmarshal(body, &da)
	return da, err
}
