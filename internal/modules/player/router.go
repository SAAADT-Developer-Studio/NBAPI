package player

import (
	"github.com/go-chi/chi/v5"
)

func Router(router chi.Router) {
	router.Get("/", PlayersHandler)

	router.Get("/awards", PlayerAwardWinnerHandler)

	router.Route("/{playerId}", func(r chi.Router) {
		r.Get("/", PlayerHandler)
		r.Get("/award-votes", PlayerAwardHandler)
		r.Get("/{stat}", PlayerSpecificStatsHandler)
	})
}
