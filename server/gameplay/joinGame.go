package gameplay

import (
	"errors"
	"log"
	"net/http"

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
	g.lock.Lock()
	defer g.lock.Unlock()

	err, code := joinGame(g)

	if err != nil {
		w.WriteHeader(code)
		w.Write([]byte(err.Error()))
	} else {
		w.Write([]byte(g.player2.id))
	}
}

func joinGame(g *game) (error, int) {
	if g.player2 != nil {
		return errors.New("Game already full.\n"), 400
	}

	// player2Token, err := uuid.NewRandom()
	// if err != nil {
	// 	return errors.New("Could not join game.\n"), 500
	// }

	p2 := player{
		Equipment:   make(map[string]*equipment),
		Spirits:     make(map[string]*spirit),
		id:          "79cfacdf-d53f-49cd-be9a-c2ad846ef13b", // player2Token.String(),
		nextActions: nil,
	}

	g.Players[p2.id] = &p2
	g.player2 = &p2

	initializeDummyPlayer2(&p2)

	p2.opponent = g.player1
	g.player1.opponent = &p2

	return nil, 200
}
