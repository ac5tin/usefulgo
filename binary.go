package utils

import (
	"bytes"
	"encoding/gob"
)

// GetBytes returns binary byte slices
func GetBytes[T any](key T) (*[]byte, error) {
	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	err := enc.Encode(key)
	if err != nil {
		return nil, err
	}
	b := buf.Bytes()
	return &b, nil
}
