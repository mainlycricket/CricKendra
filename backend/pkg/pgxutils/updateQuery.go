package pgxutils

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
)

func GetSetStatement[T any](idField string, data T) (string, []any, error) {
	v := reflect.ValueOf(data)
	t := v.Type()

	if t.Kind() != reflect.Struct {
		return "", nil, errors.New("not a struct")
	}

	var paramCount int = 1
	var setClauses []string
	var args []any

	for i := 0; i < v.NumField(); i++ {
		fieldType := t.Field(i)
		tag := fieldType.Tag.Get("json")
		columnName := strings.Split(tag, ",")[0]

		if columnName == idField || columnName == "" {
			continue
		}

		fieldValue := v.Field(i)
		var isValueValid bool
		var value any

		if fieldValue.Kind() == reflect.Slice || fieldValue.Kind() == reflect.Array {
			if fieldValue.Len() > 0 {
				isValueValid = true
			}
		} else {
			validField := fieldValue.FieldByName("Valid")
			if validField.IsValid() && validField.Bool() {
				valueMethod := fieldValue.MethodByName("Value")
				if valueMethod.IsValid() {
					results := valueMethod.Call(nil)
					if len(results) == 2 && results[1].IsNil() {
						isValueValid = true
						value = results[0].Interface()
					}
				}
			}
		}

		if isValueValid {
			setClauses = append(setClauses, fmt.Sprintf("%s = $%d", columnName, paramCount))
			args = append(args, value)
			paramCount++
		}
	}

	if len(setClauses) == 0 {
		return "", nil, errors.New("no field to update")
	}

	return strings.Join(setClauses, ", "), args, nil
}
