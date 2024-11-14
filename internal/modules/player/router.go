package player

import (
	"NBAPI/internal/inputs"
	"net/http"

	"github.com/danielgtaylor/huma/v2"
	"github.com/go-chi/chi/v5"
)

func RegisterRoutes(api huma.API) {
	huma.Register(api, huma.Operation{
		Method:  http.MethodGet,
		Path:    "/players",
		Tags:    []string{"players"},
		Summary: "List players",
	}, PlayersHandler)
}

func Router(router chi.Router) {
	router.Use(inputs.SeasonYearMiddleware)
	// router.Get("/", PlayersHandler)
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
