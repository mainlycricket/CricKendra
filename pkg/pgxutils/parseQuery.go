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

		kind := field.Type.Kind()

		columnName := field.Tag.Get("json")
		columnNames = append(columnNames, columnName)

		var datatype string
		if kind == reflect.Slice || kind == reflect.Array {
			datatype = field.Type.Elem().Name()
		} else {
			datatype = field.Type.Name()
		}

		pgType := getPgType(datatype)
		if pgType == "" {
			continue
		}

		values := input.UrlQuery[columnName]
		var argPlaceHolders []string
		for _, value := range values {
			parsed, err := ParseArg(value, datatype)
			if err == nil {
				res.Args = append(res.Args, parsed)
				argPlaceHolders = append(argPlaceHolders, fmt.Sprintf(`$%d`, len(res.Args)))
			}
		}

		if len(argPlaceHolders) > 0 {
			condition := fmt.Sprintf(`%s.%s `, input.TableName, columnName)
			placeholderString := strings.Join(argPlaceHolders, ",")

			if kind == reflect.Slice || kind == reflect.Array {
				condition += fmt.Sprintf(`&& ARRAY[%v]::%s`, placeholderString, pgType)
			} else {
				condition += fmt.Sprintf(`IN (%s)`, placeholderString)
			}

			conditions = append(conditions, condition)
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

	page, _ := strconv.Atoi(input.UrlQuery.Get("__page"))
	if page < 1 {
		page = 1
	}

	limit, _ := strconv.Atoi(input.UrlQuery.Get("__limit"))
	if limit < 1 {
		limit = input.DefaultLimit
	}

	res.RecordsCount = limit
	skip := (page - 1) * limit
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
