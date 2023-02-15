package helper

import (
	"fmt"
	"net/url"
	"strconv"
)

const DefaultLimit = 100

func GetPagination(params url.Values) (pageNum, limitNum int, err error) {
	page := params.Get("page")

	if page != "" {
		pageNum, err = strconv.Atoi(page)
		if err != nil || pageNum < 0 {
			return 0, 0, fmt.Errorf("conv page")
		}
	}

	limit := params.Get("limit")

	limitNum = DefaultLimit

	if limit != "" {
		limitNum, err = strconv.Atoi(limit)
		if err != nil || limitNum < 0 {
			return 0, 0, fmt.Errorf("conv limit")
		}
	}

	return
}
