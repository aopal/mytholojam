package main

import (
	"fmt"
	"mytholojam/server/types"
	"strings"
)

func help() {
	fmt.Println("Available commands:")
	fmt.Println("    help/h                            Show this help text")
	fmt.Println("    exit/e                            Exit the CLI")
	fmt.Println("    make/m [GAME]                     Make a new game")
	fmt.Println("    join/j [GAME]                     Join an existing game")
	fmt.Println("    switch/s [GAME]                   Switch to another game")
	fmt.Println("    current/c                         Print currently active game")
	fmt.Println("    list/l                            List joined games")
	fmt.Println("    team/t                            Get the status of your current team")
	fmt.Println("    op/o                              Get the status of your opponent's team")
	fmt.Println("    act/a [USER] [MOVE] [TARGET]      Submit an action for one of your spirits")
	fmt.Println("    clearact/ca                       Clear all queued actions")
	fmt.Println("    listact/la                        List all queued actions")
	fmt.Println("    submitact/sa                      Submit queued actions")
	fmt.Println("    state                             (Debug) Print current state")
}

func switchGame(game string) {
	currentGame = game
	if _, ok := actionsInProgress[game]; !ok {
		ap := types.ActionPayload{Token: tokenList[currentGame], Actions: make([]*types.Action, 0, 2)}
		actionsInProgress[game] = &ap
	}
	current()
}

func current() {
	fmt.Printf("Current game: %v\n", currentGame)
}

func list() {
	keys := make([]string, len(tokenList))

	i := 0
	for k := range tokenList {
		keys[i] = k
		i++
	}

	fmt.Printf("Joined games: %v\n", strings.Join(keys, ", "))
}

func state() {
	fmt.Printf("current game: %v\n", currentGame)
	fmt.Printf("game list: %v\n", tokenList)
}

func printCurrentTeam() {
	if currentPlayer() == nil {
		fmt.Println("You are not currently in a game")
		return
	}

	printTeam(currentPlayer())
}

func printOpponentTeam() {
	if opponent() == nil {
		fmt.Println("An opponent has not joined the game yet")
		return
	}

	printTeam(opponent())
}

func printTeam(p *types.Player) {
	for _, e := range p.Equipment {
		printEquipment(e)
	}
}

func printEquipment(e *types.Equipment) {
	fmt.Printf("%v HP: %v/%v ATK: %v WGHT: %v\n", e.Name, e.HP, e.MaxHP, e.ATK, e.Weight)
	if e.Inhabited {
		printSpirit(e.InhabitedBy)
	}
}

func printSpirit(s *types.Spirit) {
	fmt.Printf("    %v HP: %v/%v ATK: %v SPD: %v\n", s.Name, s.HP, s.MaxHP, s.ATK, s.Speed)
	for _, m := range s.Moves {
		printMove(m)
	}
	for _, m := range s.Inhabiting.Moves {
		printMove(m)
	}
}

func printMove(m *types.Move) {
	fmt.Printf("        %v PWR: %v TYPE: %v PRI: %v\n", m.Name, m.Power, m.Type, m.Priority)
}
