package gameplay

import (
	"errors"
	"log"
	"net/http"

	"mytholojam/server/types"

	"github.com/google/uuid"
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
	player1Token, err := uuid.NewRandom()
	if err != nil {
		return nil, errors.New("Could not create game.\n"), 500
	}

	g := types.Game{
		GameID:      gameID,
		Players:     make(map[string]*types.Player),
		ActionOrder: make([]*types.Action, 0),
	}
	g.Lock.Lock()
	defer g.Lock.Unlock()

	gameList[gameID] = &g

	p1 := types.Player{
		Equipment:   make(map[string]*types.Equipment),
		Spirits:     make(map[string]*types.Spirit),
		ID:          player1Token.String(), // "446f5322-ced2-4f9f-83cc-a98f9efd11f9",
		NextActions: nil,
	}
	g.Players[p1.ID] = &p1
	g.Player1 = &p1

	initializeDummyPlayer1(&p1)

	return &g, nil, 200
}
