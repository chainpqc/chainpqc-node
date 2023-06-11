package common

import (
	"encoding/json"
)

func Marshal(v any, prefix [2]byte) ([]byte, error) {

	b, err := json.Marshal(v)
	if err != nil {
		return nil, err
	}
	return b, nil
}

func Unmarshal(b []byte, prefix [2]byte, v any) error {

	err := json.Unmarshal(b, v)
	if err != nil {
		return err
	}
	return nil
}
