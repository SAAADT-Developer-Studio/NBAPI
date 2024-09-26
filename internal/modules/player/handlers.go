package player

import (
	"NBAPI/internal/database"
	"NBAPI/internal/sqlc"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"github.com/jackc/pgx/v5/pgtype"

	log "github.com/sirupsen/logrus"
)

type Ranges struct {
    Key   string
	FromTo string
    Value string
}

func PlayersHandler(w http.ResponseWriter, r *http.Request) {
	search := r.URL.Query().Get("search")
	var players []sqlc.Player
	var err error
	
	if len(search) == 0 {
		players, err = database.Queries.GetPlayers(r.Context())
	} else {
		players, err = database.Queries.GetPlayerBySearch(r.Context(), pgtype.Text{String: search})
	}
	if err != nil {
		log.Error(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Error fetching players"))
		return
	}
		render.JSON(w, r, players)
}

func PlayerHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	seasonFrom := r.URL.Query().Get("seasonFrom")
	seasonTo := r.URL.Query().Get("seasonTo")

	var seasonFromInt int
	var seasonToInt int

	if(len(seasonFrom) == 0){
		seasonFromInt = 1800
	} else {
		seasonFromInt,_ = strconv.Atoi(seasonFrom)
	}

	if(len(seasonTo) == 0){
		seasonToInt = time.Now().Year()
	} else {
		seasonToInt,_ = strconv.Atoi(seasonTo)
	}

	_playerId := chi.URLParam(r, "playerId")
	playerId, err := strconv.Atoi(_playerId)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(fmt.Sprintf("invalid playerId '%d'", playerId)))
		return
	}

	player, err := database.Queries.GetPlayer(ctx, sqlc.GetPlayerParams{ID: int32(playerId), SeasonYear: int32(seasonFromInt), SeasonYear_2: int32(seasonToInt) })

	if err != nil {
		log.Error(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Sprintf("error fetching player with id %d", playerId)))
		return
	}

	render.JSON(w, r, player)
}

func PlayerStatsHandler(w http.ResponseWriter, r *http.Request){
	ctx := r.Context()
	var validParams []Ranges
	var currentString = ""
	var isValidParam = true
	var params = r.URL.Query()
	for key, val := range params {
		if(isValidParam){
			if(len(key) > 4 && key[len(key)-4:] == "From"){
				isValidParam = false
				currentString = key[:len(key)-4]
				validParams = append(validParams, Ranges{Key: currentString, Value: val[0], FromTo: "From" })

			} else {
				w.WriteHeader(http.StatusBadRequest)
				w.Write([]byte("Use this endpoint in this style xFrom=field1&xTo=field2..."))
				return
			}
		} else {
			if(len(key) > 2 && key[len(key)-2:] == "To"){
				isValidParam = true
				if(currentString == key[:len(key)-2]){
					validParams = append(validParams, Ranges{Key: currentString, Value: val[0], FromTo: "To" })
				} else {
					w.WriteHeader(http.StatusBadRequest)
					w.Write([]byte(fmt.Sprintf("Did you mean %sFrom=value&&%sTo=value", currentString, currentString)))
					return
				}
			} else {
				w.WriteHeader(http.StatusBadRequest)
				w.Write([]byte("Use this endpoint in this style xFrom=field1&xTo=field2..."))
				return
			}
		}

		if(!isValidParam){
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(fmt.Sprintf("you forgot to add &&%sTo", currentString)))
			return
		}

	var baseQuery = `select player.id,
player.fullName,
player_totals.season_year,
totals.gp,
totals.gs,
totals.mp as total_mp,
totals.fg as totals_fg,
totals.fga as totals_fga,
totals.p3 as totals_p3,
totals.pa3 as totals_pa3,
totals.p2 as totals_p2,
totals.pa2 as totals_pa2,
totals.ft as totals_ft,
totals.fta as totals_fta,
totals.orb as totals_orb,
totals.drb as totals_drb,
totals.trb as totals_trb,
totals.stl as totals_stl,
totals.blk as totals_blk,
totals.ast as totals_ast,
totals.tov as totals_tov,
totals.pf as totals_pf,
totals.pts as totals_pts,
per_game.mp as per_game_mp,
per_game.fg as per_game_fg,
per_game.fga as per_game_fga,
per_game.fg_percent as per_game_fg_percent,
per_game.p3 as per_game_p3,
per_game.pa3 as per_game_pa3,
per_game.p_percent3 as per_game_p_percent3,
per_game.p2 as per_game_p2,
per_game.pa2 as per_game_pa2,
per_game.p_percent2 as per_game_p_percent2,
per_game.efg_percent as per_game_efg_percent,
per_game.ft as per_game_ft,
per_game.fta as per_game_fta,
per_game.ft_percent as per_game_ft_percent,
per_game.orb as per_game_orb,
per_game.drb as per_game_drb,
per_game.trb as per_game_trb,
per_game.ast as per_game_ast,
per_game.stl as per_game_stl,
per_game.blk as per_game_blk,
per_game.tov as per_game_tov,
per_game.pf as per_game_pf,
per_game.pts as per_game_pts,
per_100_possesions.fg as per_100_possesions_fg,
per_100_possesions.fga as per_100_possesions_fga,
per_100_possesions.p3 as per_100_possesions_p3,
per_100_possesions.pa3 as per_100_possesions_pa3,
per_100_possesions.p2 as per_100_possesions_p2,
per_100_possesions.pa2 as per_100_possesions_pa2,
per_100_possesions.ft as per_100_possesions_ft,
per_100_possesions.fta as per_100_possesions_fta,
per_100_possesions.orb as per_100_possesions_orb,
per_100_possesions.drb as per_100_possesions_drb,
per_100_possesions.trb as per_100_possesions_trb,
per_100_possesions.stl as per_100_possesions_stl,
per_100_possesions.blk as per_100_possesions_blk,
per_100_possesions.ast as per_100_possesions_ast,
per_100_possesions.tov as per_100_possesions_tov,
per_100_possesions.pf as per_100_possesions_pf,
per_100_possesions.pts as per_100_possesions_pts,
per_100_possesions.o_rtg as per_100_possesions_o_rtg,
per_100_possesions.d_rtg as per_100_possesions_d_rtg,
per_36.fg as per_36_fg,
per_36.fga as per_36_fga,
per_36.p3 as per_36_p3,
per_36.pa3 as per_36_pa3,
per_36.p2 as per_36_p2,
per_36.pa2 as per_36_pa2,
per_36.ft as per_36_ft,
per_36.fta as per_36_fta,
per_36.orb as per_36_orb,
per_36.drb as per_36_drb,
per_36.trb as per_36_trb,
per_36.stl as per_36_stl,
per_36.blk as per_36_blk,
per_36.ast as per_36_ast,
per_36.tov as per_36_tov,
per_36.pf as per_36_pf,
per_36.pts as per_36_pts,
player_shooting.avg_dist_fga,
player_shooting.percent_fga_from_2p_range,
player_shooting.percent_fga_from_0_3_range,
player_shooting.percent_fga_from_3_10_range,
player_shooting.percent_fga_from_10_16_range,
player_shooting.percent_fga_from_16_3p_range,
player_shooting.percent_fga_from_3p_range,
player_shooting.fg_percent_from_2p_range,
player_shooting.fg_percent_from_0_3_range,
player_shooting.fg_percent_from_3_10_range,
player_shooting.fg_percent_from_10_16_range,
player_shooting.fg_percent_from_16_3p_range,
player_shooting.fg_percent_from_3p_range,
player_shooting.percent_assisted_2p_fg,
player_shooting.percent_assisted_3p_fg,
player_shooting.percent_dunks_of_fga,
player_shooting.num_of_dunks,
player_shooting.percent_corner_3s_of_3pa,
player_shooting.corner_3_point_percent,
player_shooting.num_heaves_attempted,
player_shooting.num_heaves_made,
advanced."per",
advanced.ts_percent,
advanced.p_ar3,
advanced.f_tr,
advanced.orb_percent,
advanced.drb_percent,
advanced.trb_percent,
advanced.ast_percent,
advanced.stl_percent,
advanced.blk_percent,
advanced.tov_percent,
advanced.usg_percent,
advanced.ows,
advanced.dws,
advanced.ws,
advanced.ws48,
advanced.obpm,
advanced.dbpm,
advanced.bpm,
advanced.vorp
from
player
left join player_totals on player.id = totals.player_id
left join totals on totals.id = player_totals.total_id
left join player_per_game on player.id = player_per_game.player_id
left join per_game on per_game.id = player_per_game.per_game_id
left join player_per_100_possesions on player.id = player_per_100_possesions.player_id
left join per_100_possesions on per_100_possesions.id = player_per_100_possesions.per_100_id
left join per_36 on per_36.player_id = player.id
left join player_shooting on player_shooting.player_id = player.id
left join player_advanced on player_advanced.player_id = player_id
left join advanced on advanced.id = player_advanced.advanced_id
where `

		for idx := 0; idx < len(validParams); idx += 2 {
			baseQuery += validParams[idx].Key + " between " + validParams[idx].Value + " and " + validParams[idx + 1].Value + " " 
		}

		data, err := database.DB.Query(ctx, baseQuery)

		if err != nil {
			log.Error(err)
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(fmt.Sprintf("error doing your query: %s", baseQuery)))
			return
		}
	
		render.JSON(w, r, data)
	}
}

func PlayerSpecificStatsHandler(w http.ResponseWriter, r *http.Request){
	ctx := r.Context()
	playerId := chi.URLParam(r, "playerId")
	stats := r.URL.Query().Get("statType")

	allowedStatLookup := map[string]string{
        "PerGame":      "PerGame",
        "Per100":  "Per100Possesion",
        "Total":     "Total",
        "Per36":   "PlayerPer36",
        "Advanced":     "Advanced",
        "Shooting":     "PlayerShooting",
    }

	queriedStats := strings.Split(stats, ",")

	if(len(queriedStats) == 0){
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("The correct format is /player/{playerId}/stat?statType=tot&&statType=ppg..."))
		return
	}

	baseQuery := "select * from player"	
	allAllowed := true
    for _, stat := range queriedStats {
        if _, ok := allowedStatLookup[stat]; !ok {
            allAllowed = false
            break
        } else {
			if(stat != "Shooting" && stat != "Per36"){
				baseQuery += fmt.Sprintf(" left join Player%s on player.id = Player%s.PlayerID left join %s on %s.id = Player%s.%sID", allowedStatLookup[stat], allowedStatLookup[stat], allowedStatLookup[stat], allowedStatLookup[stat], allowedStatLookup[stat], stat)
			} else {
				baseQuery += fmt.Sprintf(" left join Player%s on player.id = Player%s.PlayerID", allowedStatLookup[stat], allowedStatLookup[stat])
			}
		}
    }

    if(!allAllowed){
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("The only allowed stat types are: pg, per100, tot, per36, adv, sht"))
		return
	}

	baseQuery += " where player.id = " + playerId 

	data, err := database.DB.Query(ctx, baseQuery)

		if err != nil {
			log.Error(err)
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(fmt.Sprintf("error doing your query: %s", baseQuery)))
			return
		}
	
		render.JSON(w, r, data)
}

