package player

import (
	"github.com/go-chi/chi/v5"
)

func Router(router chi.Router) {
	router.Get("/", PlayersHandler)
<<<<<<< Updated upstream
	router.Route("/{playerId}", func(r chi.Router) {
		r.Get("/", PlayerHandler)
=======

	router.Get("/awards", PlayerAwardWinnerHandler)
	router.Route("/all-teams", func(r chi.Router) {
		r.Get("/", AllTeamHandler)
		r.Get("/{awardType}", AllTeamTypeHandler)
	})

	router.Route("/{playerId}", func(r chi.Router) {
		r.Get("/", PlayerHandler)
		r.Get("/award-votes", PlayerAwardHandler)
		r.Get("/all-teams", AllTeamPlayerHandler)
>>>>>>> Stashed changes
		r.Get("/{stat}", PlayerSpecificStatsHandler)
	})
}
