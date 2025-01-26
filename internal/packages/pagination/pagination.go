package pagination

import (
	"fmt"
	"strconv"
)

func PerPageNumber(param string) (int, error) {
	var perPage int
	if param != "" {
		var err error
		perPage, err = strconv.Atoi(param)
		if err != nil {
			return 0, err
		}
	}
	if perPage <= 0 {
		perPage = 10
	}
	if perPage > 100 {
		perPage = 100
	}
	fmt.Println(perPage)
	return perPage, nil
}

func Paginate(param string) (int, error) {
	var page int
	if param != "" {
		var err error
		page, err = strconv.Atoi(param)
		if err != nil {
			return 0, err
		}
	}
	if page <= 0 {
		page = 1
	}
	return page, nil
}

func GetTotalPages(totalRecords int64, perPage int) int64 {
	return (totalRecords + int64(perPage) - 1) / int64(perPage)
}
