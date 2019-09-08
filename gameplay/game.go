package gameplay

import (
	"encoding/json"
	"sync"
)

type game struct {
	player1Token      string                // not shown in status calls
	player2Token      string                // not shown in status calls
	GameID            string                `json:"gameID"`
	lock              sync.Mutex            // not shown in status calls
	Player1Equip      map[string]*equipment `json:"player1Equip"`
	Player2Equip      map[string]*equipment `json:"player2Equip"`
	Player1Spirits    map[string]*spirit    `json:"player1Spirits"`
	Player2Spirits    map[string]*spirit    `json:"player2Spirits"`
	player1NextAction *action               // not shown in status calls
	player2NextAction *action               // not shown in status calls
	NumActions        int                   `json:"numActions"`
	ActionOrder       []*action             `json:"newActions"` // nextactions are only appended here once both have been received by the server
}

func (g *game) ToJSON(numActionsSeen int) ([]byte, error) {

	return json.Marshal(&struct {
		GameID         string                `json:"gameID"`
		Player1Equip   map[string]*equipment `json:"player1Equip"`
		Player2Equip   map[string]*equipment `json:"player2Equip"`
		Player1Spirits map[string]*spirit    `json:"player1Spirits"`
		Player2Spirits map[string]*spirit    `json:"player2Spirits"`
		NumActions     int                   `json:"numActions"`
		NewActions     []*action             `json:"newActions"`
	}{
		GameID:         g.GameID,
		Player1Equip:   g.Player1Equip,
		Player2Equip:   g.Player2Equip,
		Player1Spirits: g.Player1Spirits,
		Player2Spirits: g.Player2Spirits,
		NumActions:     g.NumActions,
		NewActions:     g.calculateActionsToSend(numActionsSeen),
	})
}

func (g *game) calculateActionsToSend(numActionsSeen int) []*action {
	return g.ActionOrder[numActionsSeen:]
}
