package pgxutils

import (
	"errors"
	"fmt"
	"net/url"
	"reflect"
)

func DecodeForm[T any](formValues url.Values) (T, error) {
	var data T

	v := reflect.ValueOf(&data).Elem()
	t := reflect.TypeOf(data)

	if t.Kind() != reflect.Struct {
		return data, errors.New("not a struct")
	}

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		fieldValue := v.Field(i)

		jsonTag := field.Tag.Get("json")
		if jsonTag == "" {
			jsonTag = field.Name
		}

		fieldKind := field.Type.Kind()

		switch fieldKind {
		case reflect.Slice:
			sliceValues := formValues[jsonTag]
			if len(sliceValues) > 0 {
				sliceType := field.Type.Elem()
				slice := reflect.MakeSlice(field.Type, len(sliceValues), len(sliceValues))

				for j, sliceValue := range sliceValues {
					elemValue := reflect.New(sliceType).Elem()
					scanMethod := elemValue.Addr().MethodByName("Scan")

					if !scanMethod.IsValid() {
						return data, fmt.Errorf("slice element type %s does not have a Scan method", sliceType)
					}

					results := scanMethod.Call([]reflect.Value{reflect.ValueOf(sliceValue)})

					if len(results) > 0 && !results[0].IsNil() {
						return data, results[0].Interface().(error)
					}

					slice.Index(j).Set(elemValue)
				}

				fieldValue.Set(slice)
			}
		default:
			formValue := formValues.Get(jsonTag)

			if formValue != "" {
				scanMethod := fieldValue.Addr().MethodByName("Scan")

				if !scanMethod.IsValid() {
					return data, errors.New("field " + field.Name + " does not have a Scan method")
				}

				results := scanMethod.Call([]reflect.Value{reflect.ValueOf(formValue)})

				if len(results) > 0 && !results[0].IsNil() {
					return data, results[0].Interface().(error)
				}
			}
		}

	}

	return data, nil
}

// func main() {
// 	type Person struct {
// 		Name pgtype.Text   `json:"name"`
// 		Age  pgtype.Int8   `json:"age"`
// 		Arr  []pgtype.Text `json:"arr"`
// 	}

// 	values := url.Values{"name": []string{"Tushar"}, "age": []string{"22"}, "arr": []string{"value1", "value2"}}

// 	person, err := DecodeForm[Person](values)

// 	fmt.Print(person, err)
// }
