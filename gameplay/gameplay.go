package gameplay

import (
	"log"
	"net/http"
	"sync"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

type game struct {
	player1Token      string // not shown in status calls
	player2Token      string // not shown in status calls
	gameID            string
	lock              sync.Mutex
	player1Equip      map[string]*equipment
	player2Equip      map[string]*equipment
	player1Spirits    map[string]*spirit
	player2Spirits    map[string]*spirit
	player1NextAction *action // not shown in status calls
	player2NextAction *action // not shown in status calls

	actionOrder []*action // nextactions are only appended here once both have been received by the server
	numActions  int
}

type equipment struct {
	id          string
	hp          int
	maxHp       int
	atk         int
	def         int
	moves       []*move
	inhabited   bool
	inhabitedBy *spirit
}

type spirit struct {
	id         string
	hp         int
	maxHp      int
	atk        int
	def        int
	speed      int
	moves      []*move
	inhabiting *equipment
}

type move struct { // the swap move is unique
	name  string
	power int
}

type targetable interface {
	getID() string
}

type action struct {
	user   *spirit
	target *targetable // either a spirit or equipment
	move   *move       // name of attacking move, or the special 'swap' move
}

var gameList map[string]*game
var bufferSize int

func Init() {
	gameList = make(map[string]*game)
	bufferSize = 10
}

func CreateHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("Create request received")

	vars := mux.Vars(r)
	gameID := vars["gameID"]

	if _, ok := gameList[gameID]; ok {
		w.WriteHeader(400)
		w.Write([]byte("A game with that ID already exists.\n"))
		return
	}

	player1Token, err := uuid.NewRandom()
	if err != nil {
		w.WriteHeader(500)
		return
	}

	g := game{player1Token: player1Token.String(), gameID: gameID}
	gameList[gameID] = &g

	w.Write([]byte(""))
}

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

	w.Write([]byte("hey it worked\n"))
}

func StatusHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("Status request received")

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

	w.Write([]byte(""))
}

func ActionHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("Action request received")

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

	w.Write([]byte(""))
}

func (s *spirit) getID() string {
	return s.id
}

func (e *equipment) getID() string {
	return e.id
}
