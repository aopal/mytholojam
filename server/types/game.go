package types

import (
	"encoding/json"
	"sync"
)

type Game struct {
	GameID      string             `json:"gameID"`
	Players     map[string]*Player `json:"players"`
	NumActions  int                `json:"numActions"`
	ActionOrder []*Action          `json:"newActions"` // nextactions are only appended here once both have been received by the server
	Player1     *Player            `json:"player1"`
	Player2     *Player            `json:"player2"`
	Lock        sync.RWMutex       `json:"-"`
}

func (g *Game) ToJSON(numActionsSeen int) ([]byte, error) {
	return json.Marshal(&struct {
		GameID     string    `json:"gameID"`
		Player1    *Player   `json:"player1"`
		Player2    *Player   `json:"player2"`
		NumActions int       `json:"numActions"`
		NewActions []*Action `json:"newActions"`
	}{
		GameID:     g.GameID,
		Player1:    g.Player1,
		Player2:    g.Player2,
		NumActions: g.NumActions,
		NewActions: g.calculateActionsToSend(numActionsSeen),
	})
}

func (g *Game) calculateActionsToSend(numActionsSeen int) []*Action {
	return g.ActionOrder[numActionsSeen:]
}
