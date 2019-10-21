package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

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
