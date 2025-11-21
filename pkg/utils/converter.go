package utils

import (
	"encoding/json"

	"github.com/google/uuid"
)

func StructToJSON(data interface{}) ([]byte, error) {
	bytes, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}
	return bytes, nil
}

func JSONToStruct(data []byte, v interface{}) error {
	return json.Unmarshal(data, v)
}

func StringToUUID(s string) (uuid.UUID, error) {
	return uuid.Parse(s)
}
