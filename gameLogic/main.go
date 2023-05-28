package gameLogic

import (
	"01.kood.tech/git/jsaar/tic_tac_toe/grid"
	"fmt"
)

func Launch() {
	game := true
	for game {
		grid.PrintGrid()
		userInput := ""
		fmt.Print(`
	     Options:
	     >Start<
	     >Quit<
	     :`)
		fmt.Scanln(&userInput)
		if userInput == "Quit" || userInput == "quit" {
			game = false
		} else if userInput == "Start" || userInput == "start" {
			playerXboxes := []int{}
			playerOboxes := []int{}
			for i := 1; i < 10; i++ {
				grid.PrintGrid()
				player := "x"
				if i % 2 == 0 {
					player = "o"
				} else {
					player = "x"
				}
				userChoice := 0
				fmt.Print("Choose a number from 1 to 9: ")
				fmt.Scanln(&userChoice)
				if player == "x" {
					playerXboxes = append(playerXboxes, userChoice)
				} else if player == "o" {
					playerOboxes = append(playerOboxes, userChoice)
				}
				grid.PlacePlayer(player, userChoice)
				fmt.Println(playerXboxes)
				fmt.Println(playerOboxes)
			}
		}
	}
}
