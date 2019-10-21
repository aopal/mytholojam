package main

import (
	"fmt"
	"io/ioutil"
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

	tokenList[game] = string(b)
	playerList[game] = 1
	switchGame(game)

	status()

	fmt.Printf("You have successfully created %v\n", game)
}
