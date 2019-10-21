package main

import (
	"mytholojam/server/types"
	"os"
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
var actionsInProgress map[string]*types.ActionPayload
var autoSubmit bool

func initialize() {
	gameList = make(map[string]*types.Game)
	tokenList = make(map[string]string)
	playerList = make(map[string]int)
	currentGame = ""
	actionsInProgress = make(map[string]*types.ActionPayload)

	autoSubmit = true

	server = os.Args[1]
	createEndpoint = server + "/create-game/"
	joinEndpoint = server + "/join-game/"
	statusEndpoint = server + "/status/"
	actionEndpoint = server + "/take-action/"
}
