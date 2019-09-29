package gameplay

import (
	"log"
	"net/http"

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
		w.Write([]byte(g.player1.id))
	}
}

func createGame(gameID string) (*game, error, int) {
	// player1Token, err := uuid.NewRandom()
	// if err != nil {
	// 	return nil, errors.New("Could not create game.\n"), 500
	// }

	g := game{
		GameID:      gameID,
		Players:     make(map[string]*player),
		ActionOrder: make([]*action, 0),
	}
	g.lock.Lock()
	defer g.lock.Unlock()

	gameList[gameID] = &g

	p1 := player{
		Equipment:   make(map[string]*equipment),
		Spirits:     make(map[string]*spirit),
		id:          "446f5322-ced2-4f9f-83cc-a98f9efd11f9", //player1Token.String(),
		nextActions: nil,
	}
	g.Players[p1.id] = &p1
	g.player1 = &p1

	initializeDummyPlayer1(&p1)

	return &g, nil, 200
}
