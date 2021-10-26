package utils

import "encoding/json"

func ToJson(body interface{}) ([]byte, error) {
	bytes, err := json.Marshal(body)

	if err != nil {
		return nil, err
	}

	return bytes, nil
}
