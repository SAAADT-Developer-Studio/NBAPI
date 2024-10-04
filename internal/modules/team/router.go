package team

import (
	"github.com/go-chi/chi/v5"
)

func Router(router chi.Router) {
	router.Get("/", TeamsHandler)
	router.Route("/{teamId}", func(r chi.Router) {
		r.Get("/", TeamHandler)
	})
}
