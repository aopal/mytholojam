package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/aopal/mytholojam/server/types"
)

func status() {
	res, err := http.Get(statusEndpoint + currentGame + "/0") // + currentGameData().NumActions)
	if err != nil {
		fmt.Printf("An error occurred: %v\n", err)
		return
	}

	b, _ := ioutil.ReadAll(res.Body)
	defer res.Body.Close()

	var g types.Game

	json.Unmarshal(b, &g)

	gameList[currentGame] = &g
	relink(&g)
}

func relink(g *types.Game) {
	if currentPlayer() != nil {
		relinkPlayer(currentPlayer())
	}

	if opponent() != nil {
		relinkPlayer(opponent())
	}
}

func relinkPlayer(p *types.Player) {
	for _, s := range p.Spirits {
		s.Inhabiting = p.Equipment[s.InhabitingId]
		p.Equipment[s.InhabitingId].InhabitedBy = s
	}
}
