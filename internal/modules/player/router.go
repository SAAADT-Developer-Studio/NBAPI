package player

import (
	"NBAPI/internal/middleware"

	"github.com/go-chi/chi/v5"
)

func Router(router chi.Router) {
	router.Use(middleware.SeasonYearMiddleware)
	router.Use(middleware.Pagination)
	router.Get("/", PlayersHandler)
	router.Get("/all-stars", AllStarHandler)
	router.Get("/awards", PlayerAwardWinnerHandler)
	router.Route("/all-teams", func(r chi.Router) {
		r.Get("/", AllTeamHandler)
		r.Get("/{awardType}", AllTeamTypeHandler)
	})

	router.Route("/{playerId}", func(r chi.Router) {
		r.Get("/", PlayerHandler)
		r.Get("/award-votes", PlayerAwardHandler)
		r.Get("/all-teams", AllTeamPlayerHandler)
		r.Route("/stats", func(r chi.Router) {
			r.Get("/pergame", PlayerPerGameHandler)
			r.Get("/totals", PlayerTotalsHandler)
			r.Get("/shooting", PlayerShootingHandler)
			r.Get("/per100poss", PlayerPer100Handler)
			r.Get("/per36m", PlayerPer36Handler)
			r.Get("/advanced", PlayerAdvancedHandler)
		})
	})

}
