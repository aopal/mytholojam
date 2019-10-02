package main

import (
	"bufio"
	"fmt"
	"mytholojam/server/gameplay"
	"os"
	"strings"
)

var server string
var createEndpoint string
var joinEndpoint string
var statusEndpoint string
var actionEndpoint string
var currentGame string
var gameList map[string]*gameplay.Game
var tokenList map[string]string
var playerList map[string]int
var currentGameData *gameplay.Game

func main() {
	gameList = make(map[string]*gameplay.Game)
	tokenList = make(map[string]string)
	playerList = make(map[string]int)
	currentGame = ""

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
	}

	exit := false

	for !exit {
		fmt.Print("> ")
		input, _ := reader.ReadString('\n')
		words := strings.Split(strings.TrimSpace(input), " ")

		switch words[0] {
		case "help":
			help()
		case "exit":
			exit = true
		case "create":
			create(words[1])
		case "join":
			join(words[1])
		case "my-team":
			status()
			printCurrentTeam()
		case "opponent-team":
			status()
			printOpponentTeam()
		case "act":
			status()
		case "switch":
			switchGame(words[1])
		case "current":
			fmt.Printf("Currently playing in '%v'\n", currentGame)
		case "list":
			list()
		case "state":
			state()
		case "":
		case "\n":
		default:
			fmt.Printf("Unknown command '%v'\n", words[0])
		}
	}
	fmt.Println("Goodbye")
}

func help() {
	fmt.Println("Available commands:")
	fmt.Println("    help                 Show this help text")
	fmt.Println("    exit                 Exit the CLI")
	fmt.Println("    create [GAME]        Create a new game")
	fmt.Println("    join [GAME]          Join an existing game")
	fmt.Println("    my-team              Get the status of your current team")
	fmt.Println("    opponent-team        Get the status of your opponent's team")
	fmt.Println("    switch [GAME]        Switch to another game")
	fmt.Println("    current              Print currently active game")
	fmt.Println("    list                 List joined games")
	fmt.Println("    state                (Debug) Print current state")
}

func switchGame(game string) {
	currentGame = game
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

func currentPlayer() *gameplay.Player {
	if playerList[currentGame] == 1 {
		return gameList[currentGame].Player1
	} else {
		return gameList[currentGame].Player2
	}
}

func opponent() *gameplay.Player {
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

func printTeam(p *gameplay.Player) {
	for _, e := range p.Equipment {
		printEquipment(e)
	}
}

func printEquipment(e *gameplay.Equipment) {
	fmt.Printf("%v HP: %v/%v ATK: %v WGHT: %v\n", e.Name, e.HP, e.MaxHP, e.ATK, e.Weight)
	if e.Inhabited {
		printSpirit(e.InhabitedBy)
	}
}

func printSpirit(s *gameplay.Spirit) {
	fmt.Printf("    %v HP: %v/%v ATK: %v SPD: %v\n", s.Name, s.HP, s.MaxHP, s.ATK, s.Speed)
	for _, m := range s.Moves {
		printMove(m)
	}
	for _, m := range s.Inhabiting.Moves {
		printMove(m)
	}
}

func printMove(m *gameplay.Move) {
	fmt.Printf("        %v PWR: %v TYPE: %v PRI: %v\n", m.Name, m.Power, m.Type, m.Priority)
}
