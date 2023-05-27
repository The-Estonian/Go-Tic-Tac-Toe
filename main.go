package main

import (
	"01.kood.tech/git/jsaar/tic_tac_toe/grid"
	"fmt"
)

func gameLogic() {
	game := true
	for game {
		grid.FullGrid()
		userInput := ""
		fmt.Print(`exit - to exit
o to enter o
x to enter x
: `)
		fmt.Scanln(&userInput)
		if userInput == "exit" {
			game = false
		} else if userInput == "o" {
			grid.PlacePlayer("o")
			grid.FullGrid()
		} else if userInput == "x" {
			grid.PlacePlayer("x")
			grid.FullGrid()
		}
	}
}

func main() {
	gameLogic()
}
