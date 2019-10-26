package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"mytholojam/server/types"
	"net/http"
	"strconv"
	"strings"
)

func clearActions() {
	currentAP().Actions = make([]*types.Action, 0, 2)
}

func listActions() {
	for _, a := range currentAP().Actions {
		targetNames := func(arr []*types.Equipment) []string {
			newA := make([]string, len(arr))
			for i, v := range arr {
				newA[i] = v.Name
			}
			return newA
		}(a.Targets)

		fmt.Println(a.User.Name + " will use " + a.Move.Name + " on target(s) " + strings.Join(targetNames, ", "))
	}
}

func submitActions() {
	if len(currentAP().Actions) != len(currentPlayer().Spirits) {
		fmt.Println("You have not submitted actions for all spirits.")
		fmt.Println("Currently submited: " + strconv.Itoa(len(currentAP().Actions)) + ", need: " + strconv.Itoa(len(currentPlayer().Spirits)))
		return
	}

	body, _ := json.Marshal(currentAP())

	_, err := http.Post(actionEndpoint+currentGame, "application/json", bytes.NewBuffer(body))
	if err != nil {
		fmt.Printf("An error occurred: %v\n", err)
		return
	}

	clearActions()
}

func act(user, move, target string) {
	if currentAP() == nil {
		fmt.Println("You haven't joined a game yet.")
		return
	}

	if len(currentAP().Actions) > len(currentPlayer().Spirits) {
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

	newA := types.Action{User: u, Move: m, Targets: []*types.Equipment{t}, Turn: currentGameData().TurnCount + 1}
	currentAP().Actions = append(currentAP().Actions, &newA)

	if autoSubmit && len(currentAP().Actions) == len(currentPlayer().Spirits) {
		submitActions()
	}
}
