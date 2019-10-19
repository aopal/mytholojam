package gameplay

import (
	"errors"
	"log"
	"mytholojam/server/types"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

func JoinHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("Join request received")

	vars := mux.Vars(r)
	gameID := vars["gameID"]

	if _, ok := gameList[gameID]; !ok {
		w.WriteHeader(400)
		w.Write([]byte("No game with that ID exists.\n"))
		return
	}

	g := gameList[gameID]
	g.Lock.Lock()
	defer g.Lock.Unlock()

	err, code := joinGame(g)

	if err != nil {
		w.WriteHeader(code)
		w.Write([]byte(err.Error()))
	} else {
		w.Write([]byte(g.Player2.ID))
	}
}

func joinGame(g *types.Game) (error, int) {
	if g.Player2 != nil {
		return errors.New("Game already full.\n"), 400
	}

	player2Token, err := uuid.NewRandom()
	if err != nil {
		return errors.New("Could not join game.\n"), 500
	}

	p2 := types.Player{
		Equipment:   make(map[string]*types.Equipment),
		Spirits:     make(map[string]*types.Spirit),
		ID:          player2Token.String(), // "79cfacdf-d53f-49cd-be9a-c2ad846ef13b",
		NextActions: nil,
	}

	g.Players[p2.ID] = &p2
	g.Player2 = &p2

	initializeDummyPlayer2(&p2)

	p2.Opponent = g.Player1
	g.Player1.Opponent = &p2

	return nil, 200
}
