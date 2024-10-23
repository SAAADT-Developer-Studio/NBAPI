package team

import (
	"NBAPI/internal/middleware"
	"context"
	"net/http"
	"strconv"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/sirupsen/logrus"
)

type SeasonRanges string

const (
	seasonFromKey SeasonRanges = "seasonFrom"
	seasonToKey   SeasonRanges = "seasonTo"
)

func parseSeasonYear(value string, _default int) (int, error) {
	if len(value) == 0 {
		return _default, nil
	}
	return strconv.Atoi(value)
}

func SeasonYearMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		seasonFrom := r.URL.Query().Get("seasonFrom")
		seasonTo := r.URL.Query().Get("seasonTo")

		seasonFromInt, err := parseSeasonYear(seasonFrom, 1800)
		if err != nil {
			logrus.Error(err)
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("Invalid seasonFrom parameter"))
		}
		seasonToInt, err := parseSeasonYear(seasonTo, time.Now().Year())
		if err != nil {
			logrus.Error(err)
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("Invalid seasonFrom parameter"))
		}

		ctx := context.WithValue(r.Context(), seasonFromKey, seasonFromInt)
		ctx = context.WithValue(ctx, seasonToKey, seasonToInt)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func Router(router chi.Router) {
	router.Use(middleware.Pagination)
	router.Get("/", TeamsHandler)
	router.Route("/{teamId}", func(r chi.Router) {
		r.Use(SeasonYearMiddleware)
		r.Get("/", TeamHandler)

		r.Route("/stats", func(r chi.Router) {
			r.Get("/pergame", TeamPerGameStatsHandler)
			r.Get("/per100poss", TeamPer100PossStatsHandler)
			r.Get("/totals", TeamTotalsStatsHandler)
		})
	})
}
