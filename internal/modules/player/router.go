package player

import (
	"net/http"

	"github.com/danielgtaylor/huma/v2"
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

	huma.Register(api, huma.Operation{
		Method:  http.MethodGet,
		Path:    "/players/awards",
		Tags:    []string{"players"},
		Summary: "List player award winners",
	}, PlayerAwardWinnerHandler)

	huma.Register(api, huma.Operation{
		Method:  http.MethodGet,
		Path:    "/players/all-teams",
		Tags:    []string{"players"},
		Summary: "List players all teams", // TODO: what should sumamry be?
	}, AllTeamHandler)

	huma.Register(api, huma.Operation{
		Method:  http.MethodGet,
		Path:    "/players/all-teams/{awardType}", // TODO: awardType should be an enum
		Tags:    []string{"players"},
		Summary: "List all teams by award type",
	}, AllTeamTypeHandler)

	huma.Register(api, huma.Operation{
		Method:  http.MethodGet,
		Path:    "/players/{playerId}",
		Tags:    []string{"players"},
		Summary: "Get player details",
	}, PlayerHandler)

	huma.Register(api, huma.Operation{
		Method:  http.MethodGet,
		Path:    "/players/{playerId}/award-votes",
		Tags:    []string{"players"},
		Summary: "Get player award votes",
	}, PlayerAwardHandler)

	huma.Register(api, huma.Operation{
		Method:  http.MethodGet,
		Path:    "/players/{playerId}/all-teams",
		Tags:    []string{"players"},
		Summary: "Get player all teams",
	}, AllTeamPlayerHandler)

	huma.Register(api, huma.Operation{
		Method:  http.MethodGet,
		Path:    "/players/{playerId}/stats/pergame",
		Tags:    []string{"players"},
		Summary: "Get player per game stats",
	}, PlayerPerGameHandler)

	huma.Register(api, huma.Operation{
		Method:  http.MethodGet,
		Path:    "/players/{playerId}/stats/totals",
		Tags:    []string{"players"},
		Summary: "Get player total stats",
	}, PlayerTotalsHandler)

	huma.Register(api, huma.Operation{
		Method:  http.MethodGet,
		Path:    "/players/{playerId}/stats/shooting",
		Tags:    []string{"players"},
		Summary: "Get player shooting stats",
	}, PlayerShootingHandler)

	huma.Register(api, huma.Operation{
		Method:  http.MethodGet,
		Path:    "/players/{playerId}/stats/per100poss",
		Tags:    []string{"players"},
		Summary: "Get player per 100 possessions stats",
	}, PlayerPer100Handler)

	huma.Register(api, huma.Operation{
		Method:  http.MethodGet,
		Path:    "/players/{playerId}/stats/per36m",
		Tags:    []string{"players"},
		Summary: "Get player per 36 minutes stats",
	}, PlayerPer36Handler)

	huma.Register(api, huma.Operation{
		Method:  http.MethodGet,
		Path:    "/players/{playerId}/stats/advanced",
		Tags:    []string{"players"},
		Summary: "Get player advanced stats",
	}, PlayerAdvancedHandler)
}
