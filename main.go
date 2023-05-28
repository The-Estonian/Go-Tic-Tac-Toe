package main

import (
	"01.kood.tech/git/jsaar/tic_tac_toe/grid"
	"fmt"
)

func gameLogic() {
	game := true
	for game {
		grid.PrintGrid()
		userInput := ""
		fmt.Print(`exit - to exit
o to enter o
x to enter x
: `)
		fmt.Scanln(&userInput)
		if userInput == "exit" {
			game = false
		} else if userInput == "1" {
			grid.PlacePlayer("x", 1)
			grid.PrintGrid()
		} else if userInput == "2" {
			grid.PlacePlayer("o", 2)
			grid.PrintGrid()
		} else if userInput == "3" {
			grid.PlacePlayer("x", 3)
			grid.PrintGrid()
		} else if userInput == "4" {
			grid.PlacePlayer("o", 4)
			grid.PrintGrid()
		} else if userInput == "5" {
			grid.PlacePlayer("x", 5)
			grid.PrintGrid()
		} else if userInput == "6" {
			grid.PlacePlayer("o", 6)
			grid.PrintGrid()
		} else if userInput == "7" {
			grid.PlacePlayer("x", 7)
			grid.PrintGrid()
		} else if userInput == "8" {
			grid.PlacePlayer("o", 8)
			grid.PrintGrid()
		} else if userInput == "9" {
			grid.PlacePlayer("x", 9)
			grid.PrintGrid()
		}
	}
}

func main() {
	gameLogic()
}
