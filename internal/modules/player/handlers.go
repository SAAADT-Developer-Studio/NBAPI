package player

import (
	"NBAPI/internal/database"
	"NBAPI/internal/inputs"
	"NBAPI/internal/sqlc"
	"context"
	"fmt"
	"net/http"
	"slices"
	"strconv"

	"github.com/danielgtaylor/huma/v2"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"

	log "github.com/sirupsen/logrus"
)

type PlayersResponseBody struct {
	Players  []sqlc.Player `json:"players"`
	NextPage *int32        `json:"next_page"`
}

type PlayersResponse struct {
	Body PlayersResponseBody
}

type PlayersInput struct {
	inputs.PaginationParams
	Cursor int    `query:"pageCursor" default:"0" doc:"Page cursor for pagination."`
	Search string `query:"search" doc:"Filter results based on a search string."`
}

func PlayersHandler(ctx context.Context, input *PlayersInput) (*PlayersResponse, error) {
	search := input.Search
	pageSize := int32(input.Limit)
	pageCursor := input.Cursor

	players, err := database.Queries.GetPlayers(ctx, sqlc.GetPlayersParams{Search: search, PageSize: pageSize, Cursor: int32(pageCursor)})
	if err != nil {
		return nil, err
	}

	var nextPage *int32
	if len(players) > int(pageSize) {
		nextPage = &(players[pageSize].ID)
		players = players[:pageSize]
	}

	response := &PlayersResponse{
		Body: PlayersResponseBody{Players: players, NextPage: nextPage},
	}
	return response, nil
}

type PlayerResponseBody struct {
	Player   sqlc.Player            `json:"player"`
	Totals   []sqlc.Total           `json:"totals"`
	PerGame  []sqlc.PerGame         `json:"perGame"`
	Per100   []sqlc.Per100Possesion `json:"per100"`
	Advanced []sqlc.Advanced        `json:"advanced"`
	Per36    []sqlc.PlayerPer36     `json:"per36"`
	Shooting []sqlc.PlayerShooting  `json:"shooting"`
}

type PlayerResponse struct {
	Body PlayerResponseBody
}

type PlayerInput struct {
	inputs.SeasonRangeParams
	PlayerId int `path:"playerId" doc:"The ID of the player to fetch."`
}

func PlayerHandler(ctx context.Context, input *PlayerInput) (*PlayerResponse, error) {
	seasonFrom := int32(input.SeasonFrom)
	seasonTo := int32(input.SeasonTo)
	playerId := int32(input.PlayerId)

	player, playerErr := database.Queries.GetPlayerById(ctx, int32(playerId))
	totals, totalsErr := database.Queries.GetPlayerTotals(ctx, sqlc.GetPlayerTotalsParams{ID: int32(playerId), SeasonYear: seasonFrom, SeasonYear_2: seasonTo})
	perGame, perGameErr := database.Queries.GetPlayerPerGame(ctx, sqlc.GetPlayerPerGameParams{ID: int32(playerId), SeasonYear: seasonFrom, SeasonYear_2: seasonTo})
	per100, per100Err := database.Queries.GetPlayerPer100(ctx, sqlc.GetPlayerPer100Params{ID: int32(playerId), SeasonYear: seasonFrom, SeasonYear_2: seasonTo})
	advanced, advancedErr := database.Queries.GetPlayerAdvanced(ctx, sqlc.GetPlayerAdvancedParams{ID: int32(playerId), SeasonYear: seasonFrom, SeasonYear_2: seasonTo})
	per36, per36Err := database.Queries.GetPlayerPer36(ctx, sqlc.GetPlayerPer36Params{ID: int32(playerId), SeasonYear: seasonFrom, SeasonYear_2: seasonTo})
	shooting, shootingErr := database.Queries.GetPlayerShooting(ctx, sqlc.GetPlayerShootingParams{ID: int32(playerId), SeasonYear: seasonFrom, SeasonYear_2: seasonTo})

	if playerErr != nil {
		log.Error(playerErr)
		return nil, huma.Error500InternalServerError("error fetching player", playerErr)
	}

	if shootingErr != nil {
		log.Error(shootingErr)
		return nil, huma.Error500InternalServerError("error fetching player shooting", shootingErr)
	}

	if totalsErr != nil {
		log.Error(totalsErr)
		return nil, huma.Error500InternalServerError("error fetching player totals", totalsErr)
	}

	if perGameErr != nil {
		log.Error(perGameErr)
		return nil, huma.Error500InternalServerError("error fetching player Per Game", perGameErr)
	}

	if per100Err != nil {
		log.Error(per100Err)
		return nil, huma.Error500InternalServerError("error fetching player Per 100 possesions", per100Err)
	}

	if advancedErr != nil {
		log.Error(advancedErr)
		return nil, huma.Error500InternalServerError("error fetching player Advanced", advancedErr)
	}

	if per36Err != nil {
		log.Error(per36Err)
		return nil, huma.Error500InternalServerError("error fetching player Per 36", per36Err)
	}

	playerResponse := PlayerResponseBody{
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
	if len(shooting) > 0 {
		playerResponse.Shooting = shooting
	}

	return &PlayerResponse{Body: playerResponse}, nil
}

func PlayerPerGameHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	_playerId := chi.URLParam(r, "playerId")
	playerId, playerIdErr := strconv.Atoi(_playerId)

	if playerIdErr != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Your playerId is not a number"))
		return
	}

	seasonFrom := int32(ctx.Value(inputs.SeasonFromKey).(int))
	seasonTo := int32(ctx.Value(inputs.SeasonToKey).(int))

	ppg, err := database.Queries.GetPlayerPerGame(ctx, sqlc.GetPlayerPerGameParams{
		ID:           int32(playerId),
		SeasonYear:   seasonFrom,
		SeasonYear_2: seasonTo,
	})

	if err != nil {
		log.Error(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Sprintf("error: %s", err)))
		return
	}

	render.JSON(w, r, ppg)
}

func PlayerPer100Handler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	_playerId := chi.URLParam(r, "playerId")
	playerId, playerIdErr := strconv.Atoi(_playerId)

	if playerIdErr != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Your playerId is not a number"))
		return
	}

	seasonFrom := int32(ctx.Value(inputs.SeasonFromKey).(int))
	seasonTo := int32(ctx.Value(inputs.SeasonToKey).(int))

	per100, err := database.Queries.GetPlayerPer100(ctx, sqlc.GetPlayerPer100Params{
		ID:           int32(playerId),
		SeasonYear:   seasonFrom,
		SeasonYear_2: seasonTo,
	})

	if err != nil {
		log.Error(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Sprintf("error: %s", err)))
		return
	}

	render.JSON(w, r, per100)
}

func PlayerTotalsHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	_playerId := chi.URLParam(r, "playerId")
	playerId, playerIdErr := strconv.Atoi(_playerId)

	if playerIdErr != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Your playerId is not a number"))
		return
	}

	seasonFrom := int32(ctx.Value(inputs.SeasonFromKey).(int))
	seasonTo := int32(ctx.Value(inputs.SeasonToKey).(int))

	totals, err := database.Queries.GetPlayerTotals(ctx, sqlc.GetPlayerTotalsParams{
		ID:           int32(playerId),
		SeasonYear:   seasonFrom,
		SeasonYear_2: seasonTo,
	})

	if err != nil {
		log.Error(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Sprintf("error: %s", err)))
		return
	}

	render.JSON(w, r, totals)
}

func PlayerPer36Handler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	_playerId := chi.URLParam(r, "playerId")
	playerId, playerIdErr := strconv.Atoi(_playerId)

	if playerIdErr != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Your playerId is not a number"))
		return
	}

	seasonFrom := int32(ctx.Value(inputs.SeasonFromKey).(int))
	seasonTo := int32(ctx.Value(inputs.SeasonToKey).(int))

	per36, err := database.Queries.GetPlayerPer36(ctx, sqlc.GetPlayerPer36Params{
		ID:           int32(playerId),
		SeasonYear:   seasonFrom,
		SeasonYear_2: seasonTo,
	})

	if err != nil {
		log.Error(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Sprintf("error: %s", err)))
		return
	}

	render.JSON(w, r, per36)
}

func PlayerAdvancedHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	_playerId := chi.URLParam(r, "playerId")
	playerId, playerIdErr := strconv.Atoi(_playerId)

	if playerIdErr != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Your playerId is not a number"))
		return
	}

	seasonFrom := int32(ctx.Value(inputs.SeasonFromKey).(int))
	seasonTo := int32(ctx.Value(inputs.SeasonToKey).(int))

	advanced, err := database.Queries.GetPlayerAdvanced(ctx, sqlc.GetPlayerAdvancedParams{
		ID:           int32(playerId),
		SeasonYear:   seasonFrom,
		SeasonYear_2: seasonTo,
	})

	if err != nil {
		log.Error(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Sprintf("error: %s", err)))
		return
	}

	render.JSON(w, r, advanced)
}

func PlayerShootingHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	_playerId := chi.URLParam(r, "playerId")
	playerId, playerIdErr := strconv.Atoi(_playerId)

	if playerIdErr != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Your playerId is not a number"))
		return
	}

	seasonFrom := int32(ctx.Value(inputs.SeasonFromKey).(int))
	seasonTo := int32(ctx.Value(inputs.SeasonToKey).(int))

	shooting, err := database.Queries.GetPlayerShooting(ctx, sqlc.GetPlayerShootingParams{
		ID:           int32(playerId),
		SeasonYear:   seasonFrom,
		SeasonYear_2: seasonTo,
	})

	if err != nil {
		log.Error(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Sprintf("error: %s", err)))
		return
	}

	render.JSON(w, r, shooting)
}

type PlayerAwardResponseBody struct {
	Awards []sqlc.PlayerAward `json:"awards"`
}

type PlayerAwardResponse struct {
	Body PlayerAwardResponseBody
}

func PlayerAwardHandler(ctx context.Context, input *PlayerInput) (*PlayerAwardResponse, error) {
	seasonFrom := int32(input.SeasonFrom)
	seasonTo := int32(input.SeasonTo)
	playerId := int32(input.PlayerId)

	playerAwards, err := database.Queries.GetPlayerAwards(ctx, sqlc.GetPlayerAwardsParams{PlayerID: playerId, SeasonYear: seasonFrom, SeasonYear_2: seasonTo})
	if err != nil {
		return nil, huma.Error500InternalServerError("error fetching player awards", err)
	}

	response := &PlayerAwardResponse{
		Body: PlayerAwardResponseBody{Awards: playerAwards},
	}
	return response, nil
}

type PlayerAwardWinnersResponseBody struct {
	Awards []sqlc.PlayerAward `json:"awards"`
}

type PlayerAwardWinnersResponse struct {
	Body PlayerAwardWinnersResponseBody
}

func PlayerAwardWinnerHandler(ctx context.Context, input *inputs.SeasonRangeParams) (*PlayerAwardWinnersResponse, error) {
	seasonFrom := int32(input.SeasonFrom)
	seasonTo := int32(input.SeasonTo)

	awards, err := database.Queries.GetPlayerAwardWinners(ctx, sqlc.GetPlayerAwardWinnersParams{SeasonYear: seasonFrom, SeasonYear_2: seasonTo})
	if err != nil {
		return nil, err
	}

	response := &PlayerAwardWinnersResponse{
		Body: PlayerAwardWinnersResponseBody{Awards: awards},
	}
	return response, nil
}

type AllTeamPlayerResponseBody struct {
	Teams []sqlc.GetPlayerAllTeamsRow `json:"teams"`
}

type AllTeamPlayerResponse struct {
	Body AllTeamPlayerResponseBody
}

func AllTeamPlayerHandler(ctx context.Context, input *PlayerInput) (*AllTeamPlayerResponse, error) {
	seasonFrom := int32(input.SeasonFrom)
	seasonTo := int32(input.SeasonTo)
	playerId := int32(input.PlayerId)

	allTeamPlayer, err := database.Queries.GetPlayerAllTeams(ctx, sqlc.GetPlayerAllTeamsParams{PlayerID: playerId, SeasonYear: seasonFrom, SeasonYear_2: seasonTo})
	if err != nil {
		return nil, huma.Error500InternalServerError("error fetching player all teams", err)
	}

	response := &AllTeamPlayerResponse{
		Body: AllTeamPlayerResponseBody{Teams: allTeamPlayer},
	}
	return response, nil
}

type AllTeamsResponseBody struct {
	Teams []sqlc.GetAllTeamsRow `json:"teams"`
}

type AllTeamsResponse struct {
	Body AllTeamsResponseBody
}

func AllTeamHandler(ctx context.Context, input *inputs.SeasonRangeParams) (*AllTeamsResponse, error) {
	seasonFrom := int32(input.SeasonFrom)
	seasonTo := int32(input.SeasonTo)

	allTeams, err := database.Queries.GetAllTeams(ctx, sqlc.GetAllTeamsParams{SeasonYear: seasonFrom, SeasonYear_2: seasonTo})
	if err != nil {
		return nil, err
	}

	response := &AllTeamsResponse{
		Body: AllTeamsResponseBody{Teams: allTeams},
	}
	return response, nil
}

type AllTeamsTypeResponse struct {
	Body []sqlc.GetAllTeamsTypeRow
}

type AllTeamsTypeInput struct {
	inputs.SeasonRangeParams
	AwardType string `query:"awardType" enum:"All-Rookie,All-BAA,All-Defense,All-NBA,All-ABA" doc:"The type of award to filter by."`
}

func AllTeamTypeHandler(ctx context.Context, input *AllTeamsTypeInput) (*AllTeamsTypeResponse, error) {
	seasonFrom := int32(input.SeasonFrom)
	seasonTo := int32(input.SeasonTo)
	awardType := input.AwardType

	allowedAwardTypes :=
		[]string{
			"All-Rookie",
			"All-BAA",
			"All-Defense",
			"All-NBA",
			"All-ABA"}

	if !slices.Contains(allowedAwardTypes, awardType) {
		return nil, huma.Error400BadRequest("The only allowed award types are: All-Rookie, All-BAA, All-Defense, All-NBA, All-ABA")
	}

	allTeams, err := database.Queries.GetAllTeamsType(ctx, sqlc.GetAllTeamsTypeParams{Type: awardType, SeasonYear: seasonFrom, SeasonYear_2: seasonTo})
	if err != nil {
		return nil, err
	}

	response := &AllTeamsTypeResponse{
		Body: allTeams,
	}
	return response, nil
}

type AllStarsInput struct {
	inputs.SeasonRangeParams
	Search string `query:"search" doc:"Filter results based on a search string."`
}

type AllStarsResponseBody struct {
	Players []sqlc.AllStar `json:"players"`
}

type AllStarsResponse struct {
	Body AllStarsResponseBody
}

func AllStarHandler(ctx context.Context, input *AllStarsInput) (*AllStarsResponse, error) {
	search := input.Search
	seasonFrom := int32(input.SeasonFrom)
	seasonTo := int32(input.SeasonTo)

	allStars, err := database.Queries.GetAllStars(ctx, sqlc.GetAllStarsParams{Lower: search, SeasonYear: seasonFrom, SeasonYear_2: seasonTo})
	if err != nil {
		return nil, err
	}

	response := &AllStarsResponse{
		Body: AllStarsResponseBody{Players: allStars},
	}
	return response, nil
}
