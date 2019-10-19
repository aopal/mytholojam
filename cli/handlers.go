package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"mytholojam/server/types"
	"net/http"
	"strings"
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

	tokenList[game] = string(b)
	playerList[game] = 1
	switchGame(game)

	status()

	// gameList[game].Players[token] = gameList[game].Player1

	fmt.Printf("You have successfully created %v\n", game)
}

func join(game string) {
	if _, ok := tokenList[game]; ok {
		fmt.Printf("You have already joined %v\n", game)
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

	tokenList[game] = string(b)
	playerList[game] = 2
	switchGame(game)

	status()

	fmt.Printf("You have successfully joined %v\n", game)
}

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

func act(user, move, target string) {
	if currentAP() == nil {
		fmt.Println("You haven't joined a game yet.")
		return
	}

	if len(currentAP().Actions) >= 2 {
		fmt.Println("Some weird shit happened. Clear your queued actions and try again")
		return
	}

	status()

	if opponent() == nil {
		fmt.Println("An opponent has not joined the game yet")
		return
	}

	u, m, t, err := findEntities(user, move, target)

	if err != nil {
		fmt.Println(err)
		return
	}

	newA := types.Action{User: u, Move: m, Targets: []*types.Equipment{t}}
	currentAP().Actions = append(currentAP().Actions, &newA)

	if len(currentAP().Actions) == 2 {
		body, _ := json.Marshal(currentAP())

		_, err := http.Post(actionEndpoint+currentGame, "application/json", bytes.NewBuffer(body))

		currentAP().Actions = make([]*types.Action, 0, 2)
		if err != nil {
			fmt.Printf("An error occurred: %v\n", err)
			return
		}
	}
}

func findEntities(user, move, target string) (*types.Spirit, *types.Move, *types.Equipment, error) {
	var u *types.Spirit = nil
	var m *types.Move = nil
	var t *types.Equipment = nil

	for _, v := range currentPlayer().Spirits {
		if strings.ToLower(v.Name) == strings.ToLower(user) {
			u = v
			break
		}
	}

	if u == nil {
		return nil, nil, nil, fmt.Errorf("Invalid user, could not find spirit " + user)
	}

	for _, v := range u.Moves {
		if strings.ToLower(v.Name) == strings.ToLower(move) {
			m = v
			break
		}
	}

	if m == nil {
		return nil, nil, nil, fmt.Errorf("Invalid move, could not find " + move + " for user " + user)
	}

	var teamTargeted map[string]*types.Equipment

	if m.TeamTargetable == "self" {
		teamTargeted = currentPlayer().Equipment
	} else if m.TeamTargetable == "other" {
		teamTargeted = opponent().Equipment
	}

	for _, v := range teamTargeted {
		if strings.ToLower(v.Name) == strings.ToLower(target) {
			t = v
			break
		}
	}

	if t == nil {
		return nil, nil, nil, fmt.Errorf("Invalid target, could not find equipment " + target)
	}

	return u, m, t, nil
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
