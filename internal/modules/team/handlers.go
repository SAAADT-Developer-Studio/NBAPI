package team

import (
	"NBAPI/internal/database"
	"NBAPI/internal/sqlc"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"github.com/sirupsen/logrus"
)

func TeamsHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	search := r.URL.Query().Get("search")
	teams, err := database.Queries.GetTeams(ctx, search)

	if err != nil {
		logrus.Error(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Error fetching teams"))
		return
	}
	if teams == nil {
		teams = []sqlc.Team{}
	}
	render.JSON(w, r, teams)
}

type TeamResponse struct {
	Team       sqlc.Team                `json:"team"`
	Totals     []sqlc.GetTeamTotalsRow  `json:"totals"`
	Per100Poss []sqlc.Per100Possesion   `json:"per_100_possesions"`
	PerGame    []sqlc.GetTeamPerGameRow `json:"per_game"`
}

func TeamHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	abbr := chi.URLParam(r, "teamId")
	seasonFrom := int32(ctx.Value(seasonFromKey).(int))
	seasonTo := int32(ctx.Value(seasonToKey).(int))
	team, teamErr := database.Queries.GetTeam(ctx, abbr)
	totalsRows, totalsErr := database.Queries.GetTeamTotals(ctx,
		sqlc.GetTeamTotalsParams{Abbr: abbr, SeasonYear: seasonFrom, SeasonYear_2: seasonTo},
	)
	per100Rows, per100Err := database.Queries.GetTeamPer100Possesions(ctx,
		sqlc.GetTeamPer100PossesionsParams{Abbr: abbr, SeasonYear: seasonFrom, SeasonYear_2: seasonTo},
	)
	perGameRows, perGameErr := database.Queries.GetTeamPerGame(ctx,
		sqlc.GetTeamPerGameParams{Abbr: abbr, SeasonYear: seasonFrom, SeasonYear_2: seasonTo},
	)

	if teamErr != nil {
		logrus.Error(teamErr)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Sprintf("Error fetching team with id %s", abbr)))
		return
	}

	if totalsErr != nil || per100Err != nil || perGameErr != nil {
		logrus.Error(totalsErr)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Error fetching subtables"))
		return
	}

	response := TeamResponse{Team: team, Totals: []sqlc.GetTeamTotalsRow{}, Per100Poss: []sqlc.Per100Possesion{}, PerGame: []sqlc.GetTeamPerGameRow{}}
	response.Totals = totalsRows
	response.Per100Poss = per100Rows
	response.PerGame = perGameRows

	render.JSON(w, r, response)
}

func TeamPerGameStatsHandler(w http.ResponseWriter, r *http.Request) {
	abbr := chi.URLParam(r, "teamId")
	ctx := r.Context()
	seasonFrom := int32(ctx.Value(seasonFromKey).(int))
	seasonTo := int32(ctx.Value(seasonToKey).(int))

	perGame, err := database.Queries.GetTeamPerGame(ctx,
		sqlc.GetTeamPerGameParams{Abbr: abbr, SeasonYear: seasonFrom, SeasonYear_2: seasonTo},
	)

	if err != nil {
		logrus.Error(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Error fetching team per game stats"))
		return
	}

	render.JSON(w, r, perGame)
}

func TeamPer100PossStatsHandler(w http.ResponseWriter, r *http.Request) {
	abbr := chi.URLParam(r, "teamId")
	ctx := r.Context()
	seasonFrom := int32(ctx.Value(seasonFromKey).(int))
	seasonTo := int32(ctx.Value(seasonToKey).(int))

	per100, err := database.Queries.GetTeamPer100Possesions(ctx,
		sqlc.GetTeamPer100PossesionsParams{Abbr: abbr, SeasonYear: seasonFrom, SeasonYear_2: seasonTo},
	)

	if err != nil {
		logrus.Error(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Error fetching team per game stats"))
		return
	}

	render.JSON(w, r, per100)
}

func TeamTotalsStatsHandler(w http.ResponseWriter, r *http.Request) {
	abbr := chi.URLParam(r, "teamId")
	ctx := r.Context()
	seasonFrom := int32(ctx.Value(seasonFromKey).(int))
	seasonTo := int32(ctx.Value(seasonToKey).(int))

	totals, err := database.Queries.GetTeamTotals(ctx,
		sqlc.GetTeamTotalsParams{Abbr: abbr, SeasonYear: seasonFrom, SeasonYear_2: seasonTo},
	)

	if err != nil {
		logrus.Error(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Error fetching team per game stats"))
		return
	}

	render.JSON(w, r, totals)

}
