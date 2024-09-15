package player

import (
	"github.com/go-chi/chi/v5"
)

func Router(router chi.Router) {
	router.Get("/", PlayersHandler)
	router.Route("/{playerId}", func(router chi.Router) {
		router.Get("/", PlayerHandler)
		router.Get("/season/{seasonId}/totals", PlayerHandler)
		router.Get("/season/{seasonId}/per100pos", PlayerHandler)
		router.Get("/season/{seasonId}/pergame", PlayerHandler)
	})
}
