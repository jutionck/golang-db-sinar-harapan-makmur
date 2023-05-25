package common

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jutionck/golang-db-sinar-harapan-makmur/model/dto"
	"strconv"
)

func ValidateRequestQueryParams(c *gin.Context) (dto.RequestQueryParams, error) {
	page, err := strconv.Atoi(c.DefaultQuery("page", "1"))
	if err != nil || page <= 0 {
		return dto.RequestQueryParams{}, fmt.Errorf("invalid page number")
	}

	limit, err := strconv.Atoi(c.DefaultQuery("limit", "5"))
	if err != nil || limit <= 0 {
		return dto.RequestQueryParams{}, fmt.Errorf("invalid limit value")
	}

	order := c.DefaultQuery("order", "id")
	sort := c.DefaultQuery("sort", "asc")

	return dto.RequestQueryParams{
		QueryParams: dto.QueryParams{
			Order: order,
			Sort:  sort,
		},
		PaginationParam: dto.PaginationParam{
			Page:  page,
			Limit: limit,
		},
	}, nil
}
