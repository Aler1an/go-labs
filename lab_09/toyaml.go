package main

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
)

func ToYAML(v any) (string, error) {
	return serializeYAML(reflect.ValueOf(v), 0)
}

func serializeYAML(v reflect.Value, indent int) (string, error) {
	space := strings.Repeat("  ", indent)

	switch v.Kind() {
	case reflect.String:
		return fmt.Sprintf("\"%s\"", v.String()), nil
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return fmt.Sprintf("%d", v.Int()), nil
	case reflect.Bool:
		return fmt.Sprintf("%t", v.Bool()), nil
	case reflect.Slice, reflect.Array:
		var result []string

		for i := 0; i < v.Len(); i++ {
			item, err := serializeYAML(v.Index(i), indent+1)
			if err != nil {
				return "", err
			}

			result = append(result, fmt.Sprintf("%s- %s", space, item))
		}

		return strings.Join(result, "\n"), nil
	case reflect.Struct:
		var result []string
		t := v.Type()

		for i := 0; i < v.NumField(); i++ {
			field := t.Field(i)
			tag := field.Tag.Get("json")

			if tag == "" {
				continue
			}

			fieldValue := v.Field(i)

			if fieldValue.Kind() == reflect.Slice || fieldValue.Kind() == reflect.Array {
				value, err := serializeYAML(fieldValue, indent+1)
				if err != nil {
					return "", err
				}

				result = append(result, fmt.Sprintf("%s%s:\n%s", space, tag, value))
			} else {
				value, err := serializeYAML(fieldValue, indent+1)
				if err != nil {
					return "", err
				}

				result = append(result, fmt.Sprintf("%s%s: %s", space, tag, value))
			}
		}

		return strings.Join(result, "\n"), nil

	default:
		return "", errors.New("unsupported type")

	}
}
