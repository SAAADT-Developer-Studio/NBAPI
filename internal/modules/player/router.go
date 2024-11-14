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

	// TODO: add pagination
	huma.Register(api, huma.Operation{
		Method:  http.MethodGet,
		Path:    "/players/all-stars",
		Tags:    []string{"players"},
		Summary: "List all-star players",
	}, AllStarHandler)

	// huma.Register(api, huma.Operation{
	// 	Method:  http.MethodGet,
	// 	Path:    "/players/awards",
	// 	Tags:    []string{"players"},
	// 	Summary: "List player award winners",
	// }, PlayerAwardWinnerHandler)

	// huma.Register(api, huma.Operation{
	// 	Method:  http.MethodGet,
	// 	Path:    "/players/all-teams",
	// 	Tags:    []string{"players"},
	// 	Summary: "List all teams",
	// }, AllTeamHandler)

	// huma.Register(api, huma.Operation{
	// 	Method:  http.MethodGet,
	// 	Path:    "/players/all-teams/{awardType}", // TODO: awardType should be an enum
	// 	Tags:    []string{"players"},
	// 	Summary: "List all teams by award type",
	// }, AllTeamTypeHandler)

	// huma.Register(api, huma.Operation{
	// 	Method:  http.MethodGet,
	// 	Path:    "/players/{playerId}",
	// 	Tags:    []string{"players"},
	// 	Summary: "Get player details",
	// }, PlayerHandler)

	// huma.Register(api, huma.Operation{
	// 	Method:  http.MethodGet,
	// 	Path:    "/players/{playerId}/award-votes",
	// 	Tags:    []string{"players"},
	// 	Summary: "Get player award votes",
	// }, PlayerAwardHandler)

	// huma.Register(api, huma.Operation{
	// 	Method:  http.MethodGet,
	// 	Path:    "/players/{playerId}/all-teams",
	// 	Tags:    []string{"players"},
	// 	Summary: "Get player all teams",
	// }, AllTeamPlayerHandler)

	// huma.Register(api, huma.Operation{
	// 	Method:  http.MethodGet,
	// 	Path:    "/players/{playerId}/stats/pergame",
	// 	Tags:    []string{"players"},
	// 	Summary: "Get player per game stats",
	// }, PlayerPerGameHandler)

	// huma.Register(api, huma.Operation{
	// 	Method:  http.MethodGet,
	// 	Path:    "/players/{playerId}/stats/totals",
	// 	Tags:    []string{"players"},
	// 	Summary: "Get player total stats",
	// }, PlayerTotalsHandler)

	// huma.Register(api, huma.Operation{
	// 	Method:  http.MethodGet,
	// 	Path:    "/players/{playerId}/stats/shooting",
	// 	Tags:    []string{"players"},
	// 	Summary: "Get player shooting stats",
	// }, PlayerShootingHandler)

	// huma.Register(api, huma.Operation{
	// 	Method:  http.MethodGet,
	// 	Path:    "/players/{playerId}/stats/per100poss",
	// 	Tags:    []string{"players"},
	// 	Summary: "Get player per 100 possessions stats",
	// }, PlayerPer100Handler)

	// huma.Register(api, huma.Operation{
	// 	Method:  http.MethodGet,
	// 	Path:    "/players/{playerId}/stats/per36m",
	// 	Tags:    []string{"players"},
	// 	Summary: "Get player per 36 minutes stats",
	// }, PlayerPer36Handler)

	// huma.Register(api, huma.Operation{
	// 	Method:  http.MethodGet,
	// 	Path:    "/players/{playerId}/stats/advanced",
	// 	Tags:    []string{"players"},
	// 	Summary: "Get player advanced stats",
	// }, PlayerAdvancedHandler)
}

func Router(router chi.Router) {
	router.Use(inputs.SeasonYearMiddleware)
	// router.Get("/", PlayersHandler)
	// router.Get("/all-stars", AllStarHandler)
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
