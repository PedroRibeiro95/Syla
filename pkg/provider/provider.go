package provider

import (
	"encoding/json"
)

// MarshalToJSON ...
func MarshalToJSON(i interface{}) ([]byte, error) {
	jsonMarshalled, err := json.Marshal(i)
	if err != nil {
		return []byte{}, err
	}
	return jsonMarshalled, nil
}
