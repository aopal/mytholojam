package main

import "github.com/aopal/mytholojam/server/types"

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
