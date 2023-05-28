package main

import (
	"01.kood.tech/git/jsaar/tic_tac_toe/gameLogic"
	"01.kood.tech/git/jsaar/tic_tac_toe/loadingScreen"
	// "fmt"
)


func main() {
	loadingScreen.Load()
	gameLogic.Launch()
}
