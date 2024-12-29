package pgxutils

import (
	"errors"
	"fmt"
	"net/url"
	"reflect"
	"slices"
	"strconv"
	"strings"
	"time"
)

type QueryInfoInput struct {
	UrlQuery     url.Values
	TableName    string
	DefaultLimit int
	DefaultSort  []string
}

type QueryInfoOutput struct {
	WhereClause      string
	OrderByClause    string
	PaginationClause string
	Args             []any
	RecordsCount     int
}

func ParseQuery[T any](input QueryInfoInput) (QueryInfoOutput, error) {
	defer recover()

	var res QueryInfoOutput

	var data T

	t := reflect.TypeOf(data)

	if t.Kind() != reflect.Struct {
		return res, errors.New("not a struct")
	}

	var columnNames []string
	var conditions []string

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)

		tag := field.Tag.Get("json")
		columnName := strings.Split(tag, ",")[0]
		columnNames = append(columnNames, columnName)

		datatype, pgType, isArray := GetFieldType(field)
		if pgType == "" {
			continue
		}

		var keysWithOperator [][]string
		if isArray {
			keysWithOperator = append(keysWithOperator,
				[]string{columnName, "array_or"},
				[]string{columnName + "__all", "array_all"},
				[]string{columnName + "__exact", "array_exact"},
			)
		} else {
			keysWithOperator = append(keysWithOperator, []string{columnName, "IN"})
		}

		for _, item := range keysWithOperator {
			key, operator := item[0], item[1]

			if values := input.UrlQuery[key]; len(values) > 0 {
				args, placeholders := GetConditionArgs(values, datatype, len(res.Args))
				if len(placeholders) > 0 {
					condition := GetCondition(input.TableName, columnName, operator, pgType, placeholders)
					res.Args = append(res.Args, args...)
					conditions = append(conditions, condition)
				}
			}
		}
	}

	if len(conditions) > 0 {
		res.WhereClause = fmt.Sprintf(`WHERE %s`, strings.Join(conditions, " AND "))
	}

	/* OrderBy Clause */

	sortOrder := input.UrlQuery["__sort"]
	if len(sortOrder) == 0 {
		sortOrder = input.DefaultSort
	}

	var sortColumns []string
	for _, columnName := range sortOrder {
		descFlag := false

		if columnName[0] == '-' {
			descFlag = true
			columnName = columnName[1:]
		}

		if slices.Contains(columnNames, columnName) {
			columnName = fmt.Sprintf(`%s.%s`, input.TableName, columnName)
			if descFlag {
				columnName += " DESC"
			}
			sortColumns = append(sortColumns, columnName)
		}
	}

	res.OrderByClause = fmt.Sprintf(`ORDER BY %s`, strings.Join(sortColumns, ", "))

	/* Pagination Clause */
	skip, limit := GetPaginationParams(&input.UrlQuery)
	res.PaginationClause = fmt.Sprintf(`OFFSET %d ROWS FETCH FIRST %d ROWS ONLY`, skip, (limit + 1))

	return res, nil
}

func getPgType(datatype string) string {
	switch datatype {
	case "Int8":
		return "integer"
	case "Float8":
		return "real"
	case "Bool":
		return "bool"
	case "Date":
		return "date"
	case "Time":
		return "time"
	case "Timestamptz":
		return "timestamptz"
	case "Text":
		return "text"
	default:
		return ""
	}
}

func ParseArg(value string, datatype string) (any, error) {
	switch datatype {
	case "Int8":
		return strconv.ParseInt(value, 10, 64)
	case "Float8":
		return strconv.ParseFloat(value, 64)
	case "Bool":
		return strconv.ParseBool(value)
	case "Date":
		return time.Parse(time.DateOnly, value)
	case "Time":
		return time.Parse(time.TimeOnly, value)
	case "Timestamptz":
		return time.Parse(time.RFC3339, value)
	case "Text":
		return value, nil
	default:
		return nil, errors.New("unsupported type")
	}
}

func GetConditionArgs(values []string, datatype string, argsLen int) ([]any, []string) {
	var argPlaceHolders []string
	var args []any

	for _, value := range values {
		parsed, err := ParseArg(value, datatype)
		if err == nil {
			args = append(args, parsed)
			argsLen++
			argPlaceHolders = append(argPlaceHolders, fmt.Sprintf(`$%d`, argsLen))
		}
	}

	return args, argPlaceHolders
}

func GetCondition(tableName, columnName, operator, pgType string, placeholders []string) string {
	condition := fmt.Sprintf(`%s.%s `, tableName, columnName)
	placeholderString := strings.Join(placeholders, ",")

	switch operator {
	case "array_or":
		condition += fmt.Sprintf(`&& ARRAY[%v]::%s[]`, placeholderString, pgType)
	case "array_exact":
		condition += fmt.Sprintf(`@> ARRAY[%v]::%s[] AND ARRAY[%v]::%s[] @> %s.%s`, placeholderString, pgType, placeholderString, pgType, tableName, columnName)
	case "array_all":
		condition += fmt.Sprintf(`@> ARRAY[%v]::%s[]`, placeholderString, pgType)
	default:
		condition += fmt.Sprintf(`IN (%s)`, placeholderString)
	}

	return condition
}

func GetFieldType(field reflect.StructField) (datatype, pgType string, isArray bool) {
	kind := field.Type.Kind()

	if kind == reflect.Slice || kind == reflect.Array {
		isArray = true
		datatype = field.Type.Elem().Name()
	} else {
		datatype = field.Type.Name()
	}

	pgType = getPgType(datatype)
	return
}
