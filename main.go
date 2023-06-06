package main

import (
	"01.kood.tech/git/jsaar/tic_tac_toe/gameLogic"
	"01.kood.tech/git/jsaar/tic_tac_toe/loadingScreen"
	"01.kood.tech/git/jsaar/tic_tac_toe/menu"
)

func main() {
	loadingScreen.Load()
	menu.Load()
	gameLogic.Launch()
}
