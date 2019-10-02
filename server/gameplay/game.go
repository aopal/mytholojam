package gameplay

import (
	"encoding/json"
	"sync"
)

type game struct {
	GameID      string             `json:"gameID"`
	Players     map[string]*player `json:"players"`
	NumActions  int                `json:"numActions"`
	ActionOrder []*action          `json:"newActions"` // nextactions are only appended here once both have been received by the server

	player1 *player
	player2 *player
	lock    sync.RWMutex
}

type player struct {
	Equipment   map[string]*equipment `json:"equipment"`
	Spirits     map[string]*spirit    `json:"spirits"`
	id          string
	opponent    *player
	nextActions []*action
}

func (g *game) ToJSON(numActionsSeen int) ([]byte, error) {
	return json.Marshal(&struct {
		GameID     string    `json:"gameID"`
		Player1    player    `json:"player1"`
		Player2    player    `json:"player2"`
		NumActions int       `json:"numActions"`
		NewActions []*action `json:"newActions"`
	}{
		GameID:     g.GameID,
		Player1:    *g.player1,
		Player2:    *g.player2,
		NumActions: g.NumActions,
		NewActions: g.calculateActionsToSend(numActionsSeen),
	})
}

func (g *game) calculateActionsToSend(numActionsSeen int) []*action {
	return g.ActionOrder[numActionsSeen:]
}
