package player

import (
	"NBAPI/internal/database"
	"NBAPI/internal/middleware"
	"NBAPI/internal/sqlc"
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"

	log "github.com/sirupsen/logrus"
)

type PlayersResponse struct {
	Players  []sqlc.Player `json:"players"`
	NextPage *int32        `json:"next_page"`
}

func PlayersHandler(w http.ResponseWriter, r *http.Request) {
	search := r.URL.Query().Get("search")
	pageSize := int32(r.Context().Value(middleware.PageSizeKey).(int))

	pageCursorQuery := r.Context().Value(middleware.PageCursorKey).(string)
	if len(pageCursorQuery) == 0 {
		pageCursorQuery = "0" // not the best but it works so shut up
	}
	pageCursor, err := strconv.Atoi(pageCursorQuery)
	if err != nil {
		http.Error(w, "Invalid page cursor", http.StatusBadRequest)
		return
	}

	players, err := database.Queries.GetPlayers(r.Context(), sqlc.GetPlayersParams{Search: search, PageSize: pageSize, Cursor: int32(pageCursor)})
	if err != nil {
		log.Error(err)
		http.Error(w, "Error fetching players", http.StatusInternalServerError)
		return
	}

	var nextPage *int32
	if len(players) > 0 {
		nextPage = &(players[len(players)-1].ID)
		players = players[:len(players)-1]
	}
	render.JSON(w, r, PlayersResponse{Players: players, NextPage: nextPage})
}

type PlayerResponse struct {
	Player   sqlc.Player            `json:"player"`
	Totals   []sqlc.Total           `json:"totals"`
	PerGame  []sqlc.PerGame         `json:"perGame"`
	Per100   []sqlc.Per100Possesion `json:"per100"`
	Advanced []sqlc.Advanced        `json:"advanced"`
	Per36    []sqlc.PlayerPer36     `json:"per36"`
	Shooting []sqlc.PlayerShooting  `json:"shooting"`
}

func PlayerHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	seasonFrom := int32(ctx.Value(middleware.SeasonFromKey).(int))
	seasonTo := int32(ctx.Value(middleware.SeasonToKey).(int))

	_playerId := chi.URLParam(r, "playerId")
	playerId, err := strconv.Atoi(_playerId)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(fmt.Sprintf("invalid playerId '%d'", playerId)))
		return
	}

	player, playerErr := database.Queries.GetPlayerById(ctx, int32(playerId))
	totals, totalsErr := database.Queries.GetPlayerTotals(ctx, sqlc.GetPlayerTotalsParams{ID: int32(playerId), SeasonYear: seasonFrom, SeasonYear_2: seasonTo})
	perGame, perGameErr := database.Queries.GetPlayerPerGame(ctx, sqlc.GetPlayerPerGameParams{ID: int32(playerId), SeasonYear: seasonFrom, SeasonYear_2: seasonTo})
	per100, per100Err := database.Queries.GetPlayerPer100(ctx, sqlc.GetPlayerPer100Params{ID: int32(playerId), SeasonYear: seasonFrom, SeasonYear_2: seasonTo})
	advanced, advancedErr := database.Queries.GetPlayerAdvanced(ctx, sqlc.GetPlayerAdvancedParams{ID: int32(playerId), SeasonYear: seasonFrom, SeasonYear_2: seasonTo})
	per36, per36Err := database.Queries.GetPlayerPer36(ctx, sqlc.GetPlayerPer36Params{ID: int32(playerId), SeasonYear: seasonFrom, SeasonYear_2: seasonTo})
	shooting, shootingErr := database.Queries.GetPlayerShooting(ctx, sqlc.GetPlayerShootingParams{ID: int32(playerId), SeasonYear: seasonFrom, SeasonYear_2: seasonTo})

	if playerErr != nil {
		log.Error(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Sprintf("error fetching player with id %d", playerId)))
		return
	}

	if shootingErr != nil {
		log.Error(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Sprintf("error fetching player shooting with id %d", playerId)))
		return
	}

	if totalsErr != nil {
		log.Error(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Sprintf("error fetching player totals with id %d", playerId)))
		return
	}

	if perGameErr != nil {
		log.Error(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Sprintf("error fetching player Per Game with id %d", playerId)))
		return
	}

	if per100Err != nil {
		log.Error(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Sprintf("error fetching player Per 100 possesions with id %d", playerId)))
		return
	}

	if advancedErr != nil {
		log.Error(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Sprintf("error fetching player Advanced with id %d", playerId)))
		return
	}

	if per36Err != nil {
		log.Error(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Sprintf("error fetching player Per 36 with id %d", playerId)))
		return
	}

	playerResponse := PlayerResponse{
		Player:   sqlc.Player{},
		Totals:   []sqlc.Total{},
		PerGame:  []sqlc.PerGame{},
		Per100:   []sqlc.Per100Possesion{},
		Advanced: []sqlc.Advanced{},
		Per36:    []sqlc.PlayerPer36{},
		Shooting: []sqlc.PlayerShooting{},
	}

	playerResponse.Player = player
	if len(totals) > 0 {
		playerResponse.Totals = totals
	}

	if len(perGame) > 0 {
		playerResponse.PerGame = perGame
	}

	if len(per100) > 0 {
		playerResponse.Per100 = per100
	}

	if len(advanced) > 0 {
		playerResponse.Advanced = advanced
	}
	if len(per36) > 0 {
		playerResponse.Per36 = per36
	}
	if len(per36) > 0 {
		playerResponse.Shooting = shooting
	}

	render.JSON(w, r, playerResponse)
}

func keyExists(key string, m map[string]func() (interface{}, error)) bool {
	_, exists := m[key]
	return exists
}

func PlayerSpecificStatsHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	_playerId := chi.URLParam(r, "playerId")
	playerId, playerIdErr := strconv.Atoi(_playerId)

	if playerIdErr != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Your playerId is not a number"))
		return
	}

	stat := chi.URLParam(r, "stat")
	seasonFrom := int32(ctx.Value(middleware.SeasonFromKey).(int))
	seasonTo := int32(ctx.Value(middleware.SeasonToKey).(int))

	allowedStatLookup := map[string]func() (interface{}, error){
		"pg": func() (interface{}, error) {
			return database.Queries.GetPlayerPerGame(ctx, sqlc.GetPlayerPerGameParams{
				ID:           int32(playerId),
				SeasonYear:   seasonFrom,
				SeasonYear_2: seasonTo,
			})
		},
		"per100": func() (interface{}, error) {
			return database.Queries.GetPlayerPer100(ctx, sqlc.GetPlayerPer100Params{
				ID:           int32(playerId),
				SeasonYear:   seasonFrom,
				SeasonYear_2: seasonTo,
			})
		},
		"tot": func() (interface{}, error) {
			return database.Queries.GetPlayerTotals(ctx, sqlc.GetPlayerTotalsParams{
				ID:           int32(playerId),
				SeasonYear:   seasonFrom,
				SeasonYear_2: seasonTo,
			})
		},
		"per36": func() (interface{}, error) {
			return database.Queries.GetPlayerPer36(ctx, sqlc.GetPlayerPer36Params{
				ID:           int32(playerId),
				SeasonYear:   seasonFrom,
				SeasonYear_2: seasonTo,
			})
		},
		"adv": func() (interface{}, error) {
			return database.Queries.GetPlayerAdvanced(ctx, sqlc.GetPlayerAdvancedParams{
				ID:           int32(playerId),
				SeasonYear:   seasonFrom,
				SeasonYear_2: seasonTo,
			})
		},
		"sht": func() (interface{}, error) {
			return database.Queries.GetPlayerShooting(ctx, sqlc.GetPlayerShootingParams{
				ID:           int32(playerId),
				SeasonYear:   seasonFrom,
				SeasonYear_2: seasonTo,
			})
		},
	}

	if len(stat) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("The correct format is /player/{playerId}/stat?statType=tot&&statType=ppg..."))
		return
	}

	if !keyExists(stat, allowedStatLookup) {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("The only allowed stat types are: pg, per100, tot, per36, adv, sht"))
		return
	}

	data, dataErr := allowedStatLookup[stat]()

	if dataErr != nil {
		log.Error(dataErr)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Sprintf("error: %s", dataErr)))
		return
	}

	render.JSON(w, r, data)
}

func PlayerAwardHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	_playerId := chi.URLParam(r, "playerId")
	playerId, playerIdErr := strconv.Atoi(_playerId)

	if playerIdErr != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Your playerId is not a number"))
		return
	}

	playerAward, err := database.Queries.GetPlayerAwards(ctx, sqlc.GetPlayerAwardsParams{PlayerID: int32(playerId)})

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Error doing your query"))
		return
	}

	render.JSON(w, r, playerAward)
}

func PlayerAwardWinnerHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	seasonFrom := int32(ctx.Value(middleware.SeasonFromKey).(int))
	seasonTo := int32(ctx.Value(middleware.SeasonToKey).(int))

	awards, err := database.Queries.GetAwardWinners(r.Context(), sqlc.GetAwardWinnersParams{SeasonYear: seasonFrom, SeasonYear_2: seasonTo})

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Error doing your query"))
		return
	}

	render.JSON(w, r, awards)

}

func AllTeamPlayerHandler(w http.ResponseWriter, r *http.Request) {
	_playerId := chi.URLParam(r, "playerId")
	playerId, playerIdErr := strconv.Atoi(_playerId)

	if playerIdErr != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Your playerId is not a number"))
		return
	}

	ctx := r.Context()
	seasonFrom := int32(ctx.Value(middleware.SeasonFromKey).(int))
	seasonTo := int32(ctx.Value(middleware.SeasonToKey).(int))

	allTeamPlayer, err := database.Queries.GetPlayerAllTeams(r.Context(), sqlc.GetPlayerAllTeamsParams{PlayerID: int32(playerId), SeasonYear: seasonFrom, SeasonYear_2: seasonTo})
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Error doing your query"))
		return
	}

	render.JSON(w, r, allTeamPlayer)

}

func AllTeamHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	seasonFrom := int32(ctx.Value(middleware.SeasonFromKey).(int))
	seasonTo := int32(ctx.Value(middleware.SeasonToKey).(int))

	allTeams, err := database.Queries.GetAllTeams(r.Context(), sqlc.GetAllTeamsParams{SeasonYear: seasonFrom, SeasonYear_2: seasonTo})
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Error doing your query"))
		return
	}

	render.JSON(w, r, allTeams)

}

func includes(arr []string, element string) bool {
	for _, item := range arr {
		if item == element {
			return true
		}
	}
	return false
}

func AllTeamTypeHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	seasonFrom := int32(ctx.Value(middleware.SeasonFromKey).(int))
	seasonTo := int32(ctx.Value(middleware.SeasonToKey).(int))
	awardType := chi.URLParam(r, "awardType")

	allowedAwardTypes :=
		[]string{
			"All-Rookie",
			"All-BAA",
			"All-Defense",
			"All-NBA",
			"All-ABA"}

	if !includes(allowedAwardTypes, awardType) {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("The only allowed award types are: All-Rookie, All-BAA, All-Defense, All-NBA, All-ABA"))
		return
	}

	allTeams, err := database.Queries.GetAllTeamsType(r.Context(), sqlc.GetAllTeamsTypeParams{Type: awardType, SeasonYear: seasonFrom, SeasonYear_2: seasonTo})
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Error doing your query"))
		return
	}

	render.JSON(w, r, allTeams)

}

func AllStarHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	seasonFrom := int32(ctx.Value(middleware.SeasonFromKey).(int))
	seasonTo := int32(ctx.Value(middleware.SeasonToKey).(int))
	search := r.URL.Query().Get("search")

	allStars, err := database.Queries.GetAllStars(r.Context(), sqlc.GetAllStarsParams{Lower: search, SeasonYear: seasonFrom, SeasonYear_2: seasonTo})
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Error doing your query"))
		return
	}

	render.JSON(w, r, allStars)

}
