package tools

import (
	"errors"
	"strconv"

	"github.com/SilleCao/golang/go-micro-demo/internal/pkg/common"
	"github.com/asaskevich/govalidator"
)

func GetPagination(page string, size string) (*common.Pagination, error) {
	return GetPaginationWithSort(page, size, "")
}

func GetPaginationWithSort(page string, size string, sort string) (*common.Pagination, error) {
	pagination := &common.Pagination{
		Sort: sort,
	}
	if !govalidator.IsInt(page) {
		return pagination, errors.New("invaild parameters")
	}
	if !govalidator.IsInt(size) {
		return pagination, errors.New("invaild parameters")
	}
	pageNum, _ := strconv.Atoi(page)
	sizeNum, _ := strconv.Atoi(size)

	pagination.Page = pageNum
	pagination.Size = sizeNum
	return pagination, nil
}
