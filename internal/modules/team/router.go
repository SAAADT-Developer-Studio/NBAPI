package team

import (
	"NBAPI/internal/middleware"

	"github.com/go-chi/chi/v5"
)

func Router(router chi.Router) {
	router.Use(middleware.Pagination)
	router.Get("/", TeamsHandler)
	router.Route("/{teamId}", func(r chi.Router) {
		r.Use(middleware.SeasonYearMiddleware)
		r.Get("/", TeamHandler)

		r.Route("/stats", func(r chi.Router) {
			r.Get("/pergame", TeamPerGameStatsHandler)
			r.Get("/per100poss", TeamPer100PossStatsHandler)
			r.Get("/totals", TeamTotalsStatsHandler)
		})
	})
}
