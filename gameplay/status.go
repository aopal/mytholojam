package gameplay

import (
	"errors"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func StatusHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("Status request received")

	vars := mux.Vars(r)
	gameID := vars["gameID"]
	numActionsSeen, _ := strconv.Atoi(vars["actionCounter"])

	if _, ok := gameList[gameID]; !ok {
		w.WriteHeader(400)
		w.Write([]byte("No game with that ID exists.\n"))
		return
	}

	g := gameList[gameID]
	g.lock.RLock()
	defer g.lock.RUnlock()

	status, err, code := getStatus(g, numActionsSeen)

	if err != nil {
		w.WriteHeader(code)
		w.Write([]byte(err.Error()))
	} else {
		w.Write([]byte(status))
	}
}

func getStatus(g *game, numActionsSeen int) (string, error, int) {
	if numActionsSeen > g.NumActions {
		return "", errors.New("Invalid number of actions seen.\n"), 400
	}

	status, err := g.ToJSON(numActionsSeen)
	if err != nil {
		return "", errors.New("Could not get status.\n"), 500
	}

	return string(status), nil, 200
}
