package types

import (
	"encoding/json"
	"sync"
)

type Game struct {
	GameID      string             `json:"gameID"`
	Players     map[string]*Player `json:"players"`
	TurnCount   int                `json:"numTurns"`
	NumActions  int                `json:"numActions"`
	ActionOrder []*Action          `json:"newActions"` // nextactions are only appended here once both have been received by the server
	Player1     *Player            `json:"player1"`
	Player2     *Player            `json:"player2"`
	Lock        sync.RWMutex       `json:"-"`
}

func NewGame(gameID string) *Game {
	return &Game{
		GameID:      gameID,
		Players:     make(map[string]*Player),
		ActionOrder: make([]*Action, 0),
		TurnCount:   0,
	}
}

func (g *Game) ToJSON(numActionsSeen int) ([]byte, error) {
	return json.Marshal(&struct {
		GameID     string    `json:"gameID"`
		Player1    *Player   `json:"player1"`
		Player2    *Player   `json:"player2"`
		TurnCount  int       `json:"numTurns"`
		NumActions int       `json:"numActions"`
		NewActions []*Action `json:"newActions"`
	}{
		GameID:     g.GameID,
		Player1:    g.Player1,
		Player2:    g.Player2,
		TurnCount:  g.TurnCount,
		NumActions: g.NumActions,
		NewActions: g.calculateActionsToSend(numActionsSeen),
	})
}

func (g *Game) calculateActionsToSend(numActionsSeen int) []*Action {
	return g.ActionOrder[numActionsSeen:]
}
