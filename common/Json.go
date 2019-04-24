package common

import (
	"encoding/json"
)

// Decodes JSON data.
func DecodeJson(data []byte, v interface{}) error {
	if data != nil {
		var err = json.Unmarshal(data, v)

		if err != nil {
			return err
		}
	}

	return nil
}
