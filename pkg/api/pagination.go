package api

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/iamsumit/sample-go-app/pkg/db"
	"github.com/iamsumit/sample-go-app/pkg/validator"
)

// NewPagination to initialize the pagination with default values.
func NewPagination() db.Pagination {
	return db.Pagination{
		Page:      1,
		PerPage:   20,
		Sort:      "created",
		Direction: "desc",
	}
}

// PaginationParams simple function to get, validate, and compute
// pagination values
func PaginationParams(r *http.Request) (db.Pagination, error) {
	qparams := r.URL.Query()
	pagi := NewPagination()
	if val := qparams.Get("per_page"); val != "" {
		perPage, err := strconv.Atoi(val)

		if err != nil {
			return db.Pagination{}, NewError(
				fmt.Errorf("invalid perPage format: %s", val),
				http.StatusBadRequest,
				nil,
			)
		}

		pagi.PerPage = perPage
	}

	if val := qparams.Get("page"); val != "" {
		page, err := strconv.Atoi(val)
		if err != nil {
			return db.Pagination{}, NewError(
				fmt.Errorf("invalid page format: %s", val),
				http.StatusBadRequest,
				nil,
			)
		}

		pagi.Page = page
	}

	if val := qparams.Get("sort"); val != "" {
		pagi.Sort = strings.ToLower(val)
	}

	if val := qparams.Get("direction"); val != "" {
		pagi.Direction = strings.ToLower(val)
	}

	err := validator.Validate(pagi)
	if err != nil {
		return db.Pagination{}, err
	}

	return pagi, nil
}
