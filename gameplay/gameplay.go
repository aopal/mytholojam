package gameplay

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

type targetable interface {
	getID() string
}

var gameList map[string]*game
var bufferSize int
var moveList map[string]*move
var equipList map[string]*equipmentTemplate
var spiritList map[string]*spiritTemplate

func Init() {
	gameList = make(map[string]*game)
	bufferSize = 10
	moveList = make(map[string]*move)
	equipList = make(map[string]*equipmentTemplate)
	spiritList = make(map[string]*spiritTemplate)

	loadResources()
}

func loadResources() {
	moveF, _ := ioutil.ReadFile("resources/moves.json")
	equipF, _ := ioutil.ReadFile("resources/equipment.json")
	spiritF, _ := ioutil.ReadFile("resources/spirits.json")

	_ = json.Unmarshal([]byte(moveF), &moveList)
	_ = json.Unmarshal([]byte(equipF), &equipList)
	_ = json.Unmarshal([]byte(spiritF), &spiritList)

	// t, _ := json.Marshal(moveList)
	// fmt.Println(string(t), moveList["average"])

	// t2, _ := json.Marshal(spiritList)
	// fmt.Println(string(t2), spiritList["warrior"])

	// t3, _ := json.Marshal(equipList)
	// fmt.Println(string(t3), equipList["bow"])
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

	g, err, code := createGame(gameID)

	if err != nil {
		w.WriteHeader(code)
		w.Write([]byte(err.Error()))
	} else {
		w.Write([]byte(g.player1Token))
	}
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

	err, code := joinGame(g)

	if err != nil {
		w.WriteHeader(code)
		w.Write([]byte(err.Error()))
	} else {
		w.Write([]byte(g.player2Token))
	}
}

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
	g.lock.Lock()
	defer g.lock.Unlock()

	status, err, code := getStatus(g, numActionsSeen)

	if err != nil {
		w.WriteHeader(code)
		w.Write([]byte(err.Error()))
	} else {
		w.Write([]byte(status))
	}
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

func createGame(gameID string) (*game, error, int) {
	player1Token, err := uuid.NewRandom()
	if err != nil {
		return nil, errors.New("Could not create game.\n"), 500
	}

	g := game{player1Token: player1Token.String(), GameID: gameID}
	gameList[gameID] = &g

	g.Player1Equip = make(map[string]*equipment)
	g.Player1Spirits = make(map[string]*spirit)

	e1 := equipList["sword"].NewEquipment()
	e2 := equipList["shield"].NewEquipment()
	e3 := equipList["bow"].NewEquipment()

	s1 := spiritList["warrior"].NewSpirit()
	s2 := spiritList["cleric"].NewSpirit()

	g.Player1Equip[e1.ID] = e1
	g.Player1Equip[e2.ID] = e3
	g.Player1Equip[e3.ID] = e3

	g.Player1Spirits[s1.ID] = s1
	g.Player1Spirits[s2.ID] = s2

	s1.Inhabit(e1)
	s2.Inhabit(e2)

	return &g, nil, 200
}

func joinGame(g *game) (error, int) {
	if g.player2Token != "" {
		return errors.New("Game already full.\n"), 400
	}

	player2Token, err := uuid.NewRandom()
	if err != nil {
		return errors.New("Could not join game.\n"), 500
	}

	g.player2Token = player2Token.String()

	g.Player2Equip = make(map[string]*equipment)
	g.Player2Spirits = make(map[string]*spirit)

	e1 := equipList["helmet"].NewEquipment()
	e2 := equipList["breastplate"].NewEquipment()
	e3 := equipList["axe"].NewEquipment()

	s1 := spiritList["thief"].NewSpirit()
	s2 := spiritList["mage"].NewSpirit()

	g.Player2Equip[e1.ID] = e1
	g.Player2Equip[e2.ID] = e3
	g.Player2Equip[e3.ID] = e3

	g.Player2Spirits[s1.ID] = s1
	g.Player2Spirits[s2.ID] = s2

	s1.Inhabit(e3)
	s2.Inhabit(e1)

	return nil, 200
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
