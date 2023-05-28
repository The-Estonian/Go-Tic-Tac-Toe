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
		} else if userInput == "1" {
			grid.PlacePlayer("x", 2)
			grid.FullGrid()
		} else if userInput == "2" {
			grid.PlacePlayer("x", 1)
			grid.FullGrid()
		} else if userInput == "3" {
			grid.PlacePlayer("x", 3)
			grid.FullGrid()
		}
	}
}

func main() {
	gameLogic()
}
