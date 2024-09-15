package player

import (
	"NBAPI/internal/database"
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"

	log "github.com/sirupsen/logrus"
)

func PlayersHandler(w http.ResponseWriter, r *http.Request) {
	players, err := database.Queries.GetPlayers(r.Context())
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
