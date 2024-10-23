package player

import (
	"NBAPI/internal/middleware"

	"github.com/go-chi/chi/v5"
)

func Router(router chi.Router) {
	router.Use(middleware.SeasonYearMiddleware)
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
		r.Get("/{stat}", PlayerSpecificStatsHandler)
	})

}
