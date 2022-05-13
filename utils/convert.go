package shrd_utils

import "encoding/json"

func ConvertInterface(from any, to any) error {
	data, err := json.Marshal(from)

	if err != nil {
		return err
	}
	return json.Unmarshal(data, to)
}
