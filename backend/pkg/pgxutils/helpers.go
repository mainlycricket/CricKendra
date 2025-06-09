package pgxutils

import (
	"net/url"
	"strconv"
)

func GetPaginationParams(params *url.Values) (int, int) {
	var page, limit, skip int

	page, _ = strconv.Atoi(params.Get("__page"))
	page = max(page, 1)

	limit, _ = strconv.Atoi(params.Get("__limit"))
	if limit < 1 {
		limit = 50
	}

	skip = (page - 1) * limit

	return skip, limit
}
