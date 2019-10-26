package gameplay

import (
	"errors"
	"mytholojam/server/types"
	"net/http"

	"github.com/gorilla/mux"
)

func JoinHandler(w http.ResponseWriter, r *http.Request) {
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

	print(g, "Join request received")

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

	p2, _ := types.NewPlayer()
	initializeDummyPlayer(p2)

	g.Players[p2.ID] = p2
	g.Player2 = p2

	p2.Opponent = g.Player1
	g.Player1.Opponent = p2

	return nil, 200
}
