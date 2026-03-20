package main

import (
	"fmt"

	"github.com/Aler1an/go-labs/lab_06/models"
)

func ToJSONEasy(v any) (string, error) {
	switch val := v.(type) {
	case *models.Server:
		data, err := val.MarshalJSON()
		return string(data), err
	default:
		return "", fmt.Errorf("unsupported type")
	}
}
