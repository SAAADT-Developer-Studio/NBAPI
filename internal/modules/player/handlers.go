package player

import (
	"NBAPI/internal/database"
	"NBAPI/internal/sqlc"
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"github.com/jackc/pgx/v5/pgtype"

	log "github.com/sirupsen/logrus"
)

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

	_playerId := chi.URLParam(r, "playerId")
	playerId, err := strconv.Atoi(_playerId)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(fmt.Sprintf("invalid playerId '%d'", playerId)))
		return
	}

	player, err := database.Queries.GetPlayer(ctx, int32(playerId))

	if err != nil {
		log.Error(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Sprintf("error fetching player with id %d", playerId)))
		return
	}

	render.JSON(w, r, player)
}
