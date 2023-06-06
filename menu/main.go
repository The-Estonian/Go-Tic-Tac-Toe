package menu

import (
	"fmt"
	"os"
	"time"

	"01.kood.tech/git/jsaar/tic_tac_toe/gameLogic"
	"01.kood.tech/git/jsaar/tic_tac_toe/grid"
)

func Load() {
	game := true
	inputTrigger := false
	for game {
		grid.PrintGrid()
		userInput := ""
		if !inputTrigger {
			fmt.Print(`
	     Options:
	     >Start<
	     >Quit<
	     :`)
		}
		if inputTrigger {
			fmt.Print(`
	  Available commands:
	Start to start the game
	Quit to exit the game`)
			time.Sleep(time.Duration(2) * time.Second)
			inputTrigger = false
			continue
		}
		fmt.Scanln(&userInput)
		if userInput == "Quit" || userInput == "quit" {
			game = false
			os.Exit(1)
		} else if userInput == "Start" || userInput == "start" {
			gameLogic.Launch()
		} else {
			inputTrigger = true
			continue
		}
	}
}
