package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"mytholojam/server/gameplay"
	"net/http"
)

func create(game string) {
	res, err := http.Get(createEndpoint + game)
	if err != nil {
		fmt.Printf("An error occurred: %v\n", err)
		return
	}

	b, _ := ioutil.ReadAll(res.Body)
	defer res.Body.Close()

	if res.StatusCode != 200 {
		fmt.Printf("%v\n", string(b)) // The server returned an error:
		return
	}

	currentGame = game
	tokenList[game] = string(b)
	playerList[game] = 1

	status()

	// gameList[game].Players[token] = gameList[game].Player1

	fmt.Printf("You have successfully created '%v'\n", game)
}

func join(game string) {
	if _, ok := tokenList[game]; ok {
		fmt.Printf("You have already joined '%v'\n", game)
		return
	}

	res, err := http.Get(joinEndpoint + game)
	if err != nil {
		fmt.Printf("An error occurred: %v\n", err)
		return
	}

	b, _ := ioutil.ReadAll(res.Body)
	defer res.Body.Close()

	if res.StatusCode != 200 {
		fmt.Printf("%v\n", string(b)) // The server returned an error:
		return
	}

	currentGame = game
	tokenList[game] = string(b)
	playerList[game] = 2

	status()

	fmt.Printf("You have successfully joined '%v'\n", game)
}

func status() {
	res, err := http.Get(statusEndpoint + currentGame + "/0")
	if err != nil {
		fmt.Printf("An error occurred: %v\n", err)
		return
	}

	b, _ := ioutil.ReadAll(res.Body)
	defer res.Body.Close()

	var g gameplay.Game

	json.Unmarshal(b, &g)

	// fmt.Printf("%+v\n", g)
	// fmt.Printf("%+v\n", g.Player1)

	gameList[currentGame] = &g
	relink(&g)
}

func relink(g *gameplay.Game) {
	// if g.Player1 != nil && g.Player2 != nil {
	// 	g.Player1.Opponent = g.Player2
	// 	g.Player2.Opponent = g.Player1
	// }

	if currentPlayer() != nil {
		relinkPlayer(currentPlayer())
	}

	if opponent() != nil {
		relinkPlayer(opponent())
	}
}

func relinkPlayer(p *gameplay.Player) {
	for _, s := range p.Spirits {
		s.Inhabiting = p.Equipment[s.InhabitingId]
		p.Equipment[s.InhabitingId].InhabitedBy = s
		// fmt.Printf("%+v\n", s)
		// fmt.Printf("\t%+v\n\n", s.Inhabiting)
	}
}
