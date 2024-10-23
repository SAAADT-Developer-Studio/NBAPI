package player

import (
	"NBAPI/internal/database"
	"NBAPI/internal/sqlc"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"

	log "github.com/sirupsen/logrus"
)

type Ranges struct {
	Key    string
	FromTo string
	Value  string
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

func PlayersHandler(w http.ResponseWriter, r *http.Request) {
	search := r.URL.Query().Get("search")
	var players []sqlc.Player
	var err error

	if len(search) == 0 {
		players, err = database.Queries.GetPlayers(r.Context())
	} else {
		players, err = database.Queries.GetPlayerBySearch(r.Context(), search)
	}
	if err != nil {
		log.Error(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Error fetching players"))
		return
	}
	fmt.Println("players", players)
	render.JSON(w, r, players)
}

func PlayerHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	seasonFrom := r.URL.Query().Get("seasonFrom")
	seasonTo := r.URL.Query().Get("seasonTo")

	var seasonFromInt int
	var seasonToInt int

	if len(seasonFrom) == 0 {
		seasonFromInt = 1800
	} else {
		seasonFromInt, _ = strconv.Atoi(seasonFrom)
	}

	if len(seasonTo) == 0 {
		seasonToInt = time.Now().Year()
	} else {
		seasonToInt, _ = strconv.Atoi(seasonTo)
	}

	_playerId := chi.URLParam(r, "playerId")
	playerId, err := strconv.Atoi(_playerId)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(fmt.Sprintf("invalid playerId '%d'", playerId)))
		return
	}

	player, playerErr := database.Queries.GetPlayerById(ctx, int32(playerId))
	totals, totalsErr := database.Queries.GetPlayerTotals(ctx, sqlc.GetPlayerTotalsParams{ID: int32(playerId), SeasonYear: int32(seasonFromInt), SeasonYear_2: int32(seasonToInt)})
	perGame, perGameErr := database.Queries.GetPlayerPerGame(ctx, sqlc.GetPlayerPerGameParams{ID: int32(playerId), SeasonYear: int32(seasonFromInt), SeasonYear_2: int32(seasonToInt)})
	per100, per100Err := database.Queries.GetPlayerPer100(ctx, sqlc.GetPlayerPer100Params{ID: int32(playerId), SeasonYear: int32(seasonFromInt), SeasonYear_2: int32(seasonToInt)})
	advanced, advancedErr := database.Queries.GetPlayerAdvanced(ctx, sqlc.GetPlayerAdvancedParams{ID: int32(playerId), SeasonYear: int32(seasonFromInt), SeasonYear_2: int32(seasonToInt)})
	per36, per36Err := database.Queries.GetPlayerPer36(ctx, sqlc.GetPlayerPer36Params{ID: int32(playerId), SeasonYear: int32(seasonFromInt), SeasonYear_2: int32(seasonToInt)})
	shooting, shootingErr := database.Queries.GetPlayerShooting(ctx, sqlc.GetPlayerShootingParams{ID: int32(playerId), SeasonYear: int32(seasonFromInt), SeasonYear_2: int32(seasonToInt)})

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
	seasonFrom := r.URL.Query().Get("seasonFrom")
	seasonTo := r.URL.Query().Get("seasonTo")

	var seasonFromInt int
	var seasonToInt int

	if len(seasonFrom) == 0 {
		seasonFromInt = 1800
	} else {
		seasonFromInt, _ = strconv.Atoi(seasonFrom)
	}

	if len(seasonTo) == 0 {
		seasonToInt = time.Now().Year()
	} else {
		seasonToInt, _ = strconv.Atoi(seasonTo)
	}

	allowedStatLookup := map[string]func() (interface{}, error){
		"pg": func() (interface{}, error) {
			return database.Queries.GetPlayerPerGame(ctx, sqlc.GetPlayerPerGameParams{
				ID:           int32(playerId),
				SeasonYear:   int32(seasonFromInt),
				SeasonYear_2: int32(seasonToInt),
			})
		},
		"per100": func() (interface{}, error) {
			return database.Queries.GetPlayerPer100(ctx, sqlc.GetPlayerPer100Params{
				ID:           int32(playerId),
				SeasonYear:   int32(seasonFromInt),
				SeasonYear_2: int32(seasonToInt),
			})
		},
		"tot": func() (interface{}, error) {
			return database.Queries.GetPlayerTotals(ctx, sqlc.GetPlayerTotalsParams{
				ID:           int32(playerId),
				SeasonYear:   int32(seasonFromInt),
				SeasonYear_2: int32(seasonToInt),
			})
		},
		"per36": func() (interface{}, error) {
			return database.Queries.GetPlayerPer36(ctx, sqlc.GetPlayerPer36Params{
				ID:           int32(playerId),
				SeasonYear:   int32(seasonFromInt),
				SeasonYear_2: int32(seasonToInt),
			})
		},
		"adv": func() (interface{}, error) {
			return database.Queries.GetPlayerAdvanced(ctx, sqlc.GetPlayerAdvancedParams{
				ID:           int32(playerId),
				SeasonYear:   int32(seasonFromInt),
				SeasonYear_2: int32(seasonToInt),
			})
		},
		"sht": func() (interface{}, error) {
			return database.Queries.GetPlayerShooting(ctx, sqlc.GetPlayerShootingParams{
				ID:           int32(playerId),
				SeasonYear:   int32(seasonFromInt),
				SeasonYear_2: int32(seasonToInt),
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
	seasonFrom := r.URL.Query().Get("seasonFrom")
	seasonTo := r.URL.Query().Get("seasonTo")

	var seasonFromInt int
	var seasonToInt int

	if len(seasonFrom) == 0 {
		seasonFromInt = 1800
	} else {
		seasonFromInt, _ = strconv.Atoi(seasonFrom)
	}

	if len(seasonTo) == 0 {
		seasonToInt = time.Now().Year()
	} else {
		seasonToInt, _ = strconv.Atoi(seasonTo)
	}

	awards, err := database.Queries.GetAwardWinners(r.Context(), sqlc.GetAwardWinnersParams{SeasonYear: int32(seasonFromInt), SeasonYear_2: int32(seasonToInt)})
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

	seasonFrom := r.URL.Query().Get("seasonFrom")
	seasonTo := r.URL.Query().Get("seasonTo")

	var seasonFromInt int
	var seasonToInt int

	if len(seasonFrom) == 0 {
		seasonFromInt = 1800
	} else {
		seasonFromInt, _ = strconv.Atoi(seasonFrom)
	}

	if len(seasonTo) == 0 {
		seasonToInt = time.Now().Year()
	} else {
		seasonToInt, _ = strconv.Atoi(seasonTo)
	}

	allTeamPlayer, err := database.Queries.GetPlayerAllTeams(r.Context(), sqlc.GetPlayerAllTeamsParams{PlayerID: int32(playerId), SeasonYear: int32(seasonFromInt), SeasonYear_2: int32(seasonToInt)})
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Error doing your query"))
		return
	}

	render.JSON(w, r, allTeamPlayer)

}

func AllTeamHandler(w http.ResponseWriter, r *http.Request) {
	seasonFrom := r.URL.Query().Get("seasonFrom")
	seasonTo := r.URL.Query().Get("seasonTo")

	var seasonFromInt int
	var seasonToInt int

	if len(seasonFrom) == 0 {
		seasonFromInt = 1800
	} else {
		seasonFromInt, _ = strconv.Atoi(seasonFrom)
	}

	if len(seasonTo) == 0 {
		seasonToInt = time.Now().Year()
	} else {
		seasonToInt, _ = strconv.Atoi(seasonTo)
	}

	allTeams, err := database.Queries.GetAllTeams(r.Context(), sqlc.GetAllTeamsParams{SeasonYear: int32(seasonFromInt), SeasonYear_2: int32(seasonToInt)})
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
	seasonFrom := r.URL.Query().Get("seasonFrom")
	seasonTo := r.URL.Query().Get("seasonTo")
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

	var seasonFromInt int
	var seasonToInt int

	if len(seasonFrom) == 0 {
		seasonFromInt = 1800
	} else {
		seasonFromInt, _ = strconv.Atoi(seasonFrom)
	}

	if len(seasonTo) == 0 {
		seasonToInt = time.Now().Year()
	} else {
		seasonToInt, _ = strconv.Atoi(seasonTo)
	}

	allTeams, err := database.Queries.GetAllTeamsType(r.Context(), sqlc.GetAllTeamsTypeParams{Type: awardType, SeasonYear: int32(seasonFromInt), SeasonYear_2: int32(seasonToInt)})
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Error doing your query"))
		return
	}

	render.JSON(w, r, allTeams)

}
