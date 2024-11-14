package team

import (
	"net/http"

	"github.com/danielgtaylor/huma/v2"
)

func RegisterRoutes(api huma.API) {
	huma.Register(api, huma.Operation{
		Method:  http.MethodGet,
		Path:    "/teams",
		Tags:    []string{"teams"},
		Summary: "List all teams",
	}, TeamsHandler)

	huma.Register(api, huma.Operation{
		Method:  http.MethodGet,
		Path:    "/teams/{teamId}",
		Tags:    []string{"teams"},
		Summary: "Get team by ID",
	}, TeamHandler)

	huma.Register(api, huma.Operation{
		Method:  http.MethodGet,
		Path:    "/teams/{teamId}/stats/pergame",
		Tags:    []string{"teams"},
		Summary: "Get team per game stats",
	}, TeamPerGameStatsHandler)

	huma.Register(api, huma.Operation{
		Method:  http.MethodGet,
		Path:    "/teams/{teamId}/stats/per100poss",
		Tags:    []string{"teams"},
		Summary: "Get team per game stats",
	}, TeamPer100PossStatsHandler)

	huma.Register(api, huma.Operation{
		Method:  http.MethodGet,
		Path:    "/teams/{teamId}/stats/totals",
		Tags:    []string{"teams"},
		Summary: "Get team totals stats",
	}, TeamTotalsStatsHandler)

	huma.Register(api, huma.Operation{
		Method:  http.MethodGet,
		Path:    "/teams/{teamId}/opponents/pergame",
		Tags:    []string{"teams"},
		Summary: "Get team opponents per game stats",
	}, TeamPerGameOpponentsHandler)

	huma.Register(api, huma.Operation{
		Method:  http.MethodGet,
		Path:    "/teams/{teamId}/opponents/per100poss",
		Tags:    []string{"teams"},
		Summary: "Get team opponents per game stats",
	}, TeamPer100PossOpponentsHandler)

	huma.Register(api, huma.Operation{
		Method:  http.MethodGet,
		Path:    "/teams/{teamId}/opponents/totals",
		Tags:    []string{"teams"},
		Summary: "Get team opponents totals stats",
	}, TeamTotalsOpponentsHandler)
}
