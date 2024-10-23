package middleware

import (
	"context"
	"net/http"
	"strconv"
)

type PaginationSearchParam string

const (
	PageCursorKey PaginationSearchParam = "pageCursor"
	PageSizeKey   PaginationSearchParam = "pageSize"
)

const DEFAULT_PAGE_SIZE = 10
const MAX_PAGE_SIZE = 100
const MIN_PAGE_SIZE = 10

func getPageSize(pageSizeQuery string) (int, error) {
	if len(pageSizeQuery) == 0 {
		return DEFAULT_PAGE_SIZE, nil
	}
	return strconv.Atoi(pageSizeQuery)
}

func Pagination(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		pageCursor := r.URL.Query().Get(string(PageCursorKey))
		pageSizeQuery := r.URL.Query().Get(string(PageSizeKey))
		pageSize, err := getPageSize(pageSizeQuery)

		if err != nil {
			http.Error(w, "Invalid page size", http.StatusBadRequest)
			return
		}

		pageSize = max(MIN_PAGE_SIZE, pageSize)
		pageSize = min(MAX_PAGE_SIZE, pageSize)

		ctx := context.WithValue(r.Context(), PageCursorKey, pageCursor)
		ctx = context.WithValue(ctx, PageSizeKey, pageSize)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
