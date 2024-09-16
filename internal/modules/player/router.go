package player

import (
	"github.com/go-chi/chi/v5"
)

func Router(router chi.Router) {
	router.Get("/", PlayersHandler)
	router.Route("/{playerId}", func(router chi.Router) {
		router.Get("/", PlayerHandler)
	})
}
