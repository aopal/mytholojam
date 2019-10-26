package gameplay

import (
	"log"
	"net/http"

	"mytholojam/server/types"

	"github.com/gorilla/mux"
)

func CreateHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("Create request received")
	glLock.Lock()
	defer glLock.Unlock()

	vars := mux.Vars(r)
	gameID := vars["gameID"]

	if _, ok := gameList[gameID]; ok {
		w.WriteHeader(400)
		w.Write([]byte("A game with that ID already exists.\n"))
		return
	}

	g, err, code := createGame(gameID)

	if err != nil {
		w.WriteHeader(code)
		w.Write([]byte(err.Error()))
	} else {
		w.Write([]byte(g.Player1.ID))
	}
}

func createGame(gameID string) (*types.Game, error, int) {
	g := types.NewGame(gameID)

	g.Lock.Lock()
	defer g.Lock.Unlock()

	gameList[gameID] = g

	p1, _ := types.NewPlayer()
	initializeDummyPlayer(p1)

	g.Players[p1.ID] = p1
	g.Player1 = p1

	return g, nil, 200
}
