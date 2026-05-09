package main

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
)

func ToJSON(v any) (string, error) {
	return serialize(reflect.ValueOf(v))
}

func serialize(v reflect.Value) (string, error) {
	switch v.Kind() {

	case reflect.String:
		return fmt.Sprintf("\"%s\"", v.String()), nil

	case reflect.Int, reflect.Int64:
		return fmt.Sprintf("%d", v.Int()), nil

	case reflect.Bool:
		return fmt.Sprintf("%t", v.Bool()), nil

	case reflect.Slice:
		var result []string
		for i := 0; i < v.Len(); i++ {
			elem, _ := serialize(v.Index(i))
			result = append(result, elem)
		}
		return "[" + strings.Join(result, ",") + "]", nil

	case reflect.Struct:
		var result []string
		t := v.Type()

		for i := 0; i < v.NumField(); i++ {
			field := t.Field(i)
			tag := field.Tag.Get("json")
			if tag == "" {
				continue
			}

			value, _ := serialize(v.Field(i))
			result = append(result, fmt.Sprintf("\"%s\": %s", tag, value))
		}

		return "{ " + strings.Join(result, ", ") + " }", nil

	default:
		return "", errors.New("unsupported type")
	}
}
