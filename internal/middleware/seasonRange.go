package middleware

import (
	"context"
	"net/http"
	"strconv"
	"time"

	"github.com/sirupsen/logrus"
)

type SeasonRanges string

const (
	SeasonFromKey SeasonRanges = "seasonFrom"
	SeasonToKey   SeasonRanges = "seasonTo"
)

func parseSeasonYear(value string, defaultValue int) (int, error) {
	if len(value) == 0 {
		return defaultValue, nil
	}
	return strconv.Atoi(value)
}

func SeasonYearMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		seasonFromQuery := r.URL.Query().Get("seasonFrom")
		seasonToQuery := r.URL.Query().Get("seasonTo")

		seasonFrom, err := parseSeasonYear(seasonFromQuery, 1800)
		if err != nil {
			logrus.Error(err)
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("Invalid seasonFrom parameter"))
		}
		seasonTo, err := parseSeasonYear(seasonToQuery, time.Now().Year())
		if err != nil {
			logrus.Error(err)
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("Invalid seasonFrom parameter"))
		}

		ctx := context.WithValue(r.Context(), SeasonFromKey, seasonFrom)
		ctx = context.WithValue(ctx, SeasonToKey, seasonTo)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
