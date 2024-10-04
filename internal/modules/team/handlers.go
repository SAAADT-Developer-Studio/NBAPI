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
	Team       sqlc.Team              `json:"team"`
	Totals     []sqlc.Total           `json:"totals"`
	Per100Poss []sqlc.Per100Possesion `json:"per_100_possesions"`
	PerGame    []sqlc.PerGame         `json:"per_game"`
}

func TeamHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	abbr := chi.URLParam(r, "teamId")
	team, teamErr := database.Queries.GetTeam(ctx, abbr)
	totalsRows, totalsErr := database.Queries.GetTeamTotals(ctx, abbr)
	per100Rows, per100Err := database.Queries.GetTeamPer100Possesions(ctx, abbr)
	perGameRows, perGameErr := database.Queries.GetTeamPerGame(ctx, abbr)

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

	response := TeamResponse{Team: team, Totals: []sqlc.Total{}, Per100Poss: []sqlc.Per100Possesion{}, PerGame: []sqlc.PerGame{}}
	response.Totals = totalsRows
	response.Per100Poss = per100Rows
	response.PerGame = perGameRows

	render.JSON(w, r, response)
}
