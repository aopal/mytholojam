package main

import (
	"bufio"
	"fmt"
	"mytholojam/server/types"
	"os"
	"strings"
)

var server string
var createEndpoint string
var joinEndpoint string
var statusEndpoint string
var actionEndpoint string
var currentGame string
var gameList map[string]*types.Game
var tokenList map[string]string
var playerList map[string]int

// var currentGameData *types.Game
var actionsInProgress map[string]*types.ActionPayload

func main() {
	gameList = make(map[string]*types.Game)
	tokenList = make(map[string]string)
	playerList = make(map[string]int)
	currentGame = ""
	actionsInProgress = make(map[string]*types.ActionPayload)

	reader := bufio.NewReader(os.Stdin)

	fmt.Println("****************************************")
	fmt.Println("Type 'exit' or hit Ctrl+C to exit")
	fmt.Println("Type 'help' to see a list of commands")

	if len(os.Args) < 2 {
		fmt.Println("Usage: " + os.Args[0] + " [SERVER]")
		return
	} else {
		server = os.Args[1]
		createEndpoint = server + "/create-game/"
		joinEndpoint = server + "/join-game/"
		statusEndpoint = server + "/status/"
		actionEndpoint = server + "/take-action/"
	}

	exit := false

	for !exit {
		fmt.Print("> ")
		input, _ := reader.ReadString('\n')
		words := strings.Split(strings.TrimSpace(input), " ")

		switch words[0] {
		case "help", "h":
			help()
		case "exit", "e":
			exit = true
		case "make", "m":
			if len(words) != 2 {
				fmt.Printf("Error: Command requires 1 input\n\n")
				help()
				break
			}
			create(words[1])
		case "join", "j":
			if len(words) != 2 {
				fmt.Printf("Error: Command requires 1 input\n\n")
				help()
				break
			}
			join(words[1])
		case "team", "t":
			status()
			printCurrentTeam()
		case "opponent", "o":
			status()
			printOpponentTeam()
		case "act", "a":
			if len(words) != 4 {
				fmt.Printf("Error: Command requires 4 inputs\n\n")
				help()
				break
			}
			act(words[1], words[2], words[3])
		case "clearact", "ca":
			currentAP().Actions = make([]*types.Action, 0, 2)
		case "switch", "s":
			if len(words) != 2 {
				fmt.Printf("Error: Command requires 1 input\n\n")
				help()
				break
			}
			switchGame(words[1])
		case "current", "c":
			current()
		case "list", "l":
			list()
		case "state":
			state()
		case "":
		case "\n":
		default:
			fmt.Printf("Unknown command %v\n", words[0])
		}
	}
	fmt.Println("Goodbye")
}

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

func currentToken() string {
	return tokenList[currentGame]
}

func currentPlayer() *types.Player {
	if playerList[currentGame] == 1 {
		return gameList[currentGame].Player1
	} else {
		return gameList[currentGame].Player2
	}
}

func currentAP() *types.ActionPayload {
	return actionsInProgress[currentGame]
}

func currentGameData() *types.Game {
	return gameList[currentGame]
}

func opponent() *types.Player {
	if playerList[currentGame] == 1 {
		return gameList[currentGame].Player2
	} else {
		return gameList[currentGame].Player1
	}
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
