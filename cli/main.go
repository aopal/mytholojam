package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: " + os.Args[0] + " [SERVER]")
		return
	}

	initialize()

	reader := bufio.NewReader(os.Stdin)

	fmt.Println("****************************************")
	fmt.Println("Type 'exit' or hit Ctrl+C to exit")
	fmt.Println("Type 'help' to see a list of commands")

	exit := false
	for !exit {
		fmt.Print("> ")
		input, _ := reader.ReadString('\n')
		words := strings.Split(strings.TrimSpace(strings.ToLower(input)), " ")

		exit = handleInput(words)
	}
	fmt.Println("Goodbye")
}

func handleInput(words []string) bool {
	switch words[0] {
	case "help", "h":
		help()
	case "exit", "e":
		return true
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
		clearActions()
	case "listact", "la":
		listActions()
	case "submitact", "sa":
		submitActions()
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

	return false
}
