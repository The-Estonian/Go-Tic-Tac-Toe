package gameLogic

import (
	"fmt"
	"time"

	"01.kood.tech/git/jsaar/tic_tac_toe/grid"
)

func checkPossibleNumbers(numSlice []int, num int) bool {
	for _, v := range numSlice {
		if v == num {
			return true
		}
	}
	return false
}

func removeIntFromList(numSlice []int, num int) []int {
	returnSlice := []int{}
	for _, v := range numSlice {
		if v != num {
			returnSlice = append(returnSlice, v)
		}
	}
	return returnSlice
}

func winSolution(list []int) bool {
	if len(list) < 3 {
		return false
	}
	winner := false
	winConditionComb := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
		{1, 4, 7},
		{2, 5, 8},
		{3, 6, 9},
		{1, 5, 9},
		{7, 5, 3},
	}
	counter := 0
	for x := 0; x < len(winConditionComb); x++ {
		for y := 0; y < len(winConditionComb[x]); y++ {
			for z := 0; z < len(list); z++ {
				if winConditionComb[x][y] == list[z] {
					counter++
					if counter == 3 {
						return true
					}
				}
			}
		}
		counter = 0
	}
	return winner
}

func Launch() {
	game := true
	for game {
		boxChoices := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
		playerXboxes := []int{}
		playerOboxes := []int{}
		for i := 1; i < 10; i++ {
			grid.PrintGrid()
			if i%2 == 0 {
				fmt.Print("Player 2 choose a number from 1 to 9: ")
				playerOChoice := true
				for playerOChoice {
					var userTwoChoices int
					fmt.Scanln(&userTwoChoices)
					if checkPossibleNumbers(boxChoices, userTwoChoices) {
						playerOboxes = append(playerOboxes, userTwoChoices)
						boxChoices = removeIntFromList(boxChoices, userTwoChoices)
						grid.PlacePlayer("o", userTwoChoices)
						playerOChoice = false
						if winSolution(playerOboxes) {
							grid.PrintGrid()
							fmt.Println("Player 2 WINS!")
							time.Sleep(time.Duration(5) * time.Second)
							game = false
							i = 9
						}
					} else {
						userTwoChoices = 0
						fmt.Print("That number is taken, try another one: ")
					}
				}
			} else {
				fmt.Print("Player 1 choose a number from 1 to 9: ")
				playerXChoice := true
				for playerXChoice {
					var userOneChoices int
					fmt.Scanln(&userOneChoices)
					if checkPossibleNumbers(boxChoices, userOneChoices) {
						playerXboxes = append(playerXboxes, userOneChoices)
						boxChoices = removeIntFromList(boxChoices, userOneChoices)
						grid.PlacePlayer("x", userOneChoices)
						playerXChoice = false
						if winSolution(playerXboxes) {
							grid.PrintGrid()
							fmt.Println("Player 1 WINS!")
							time.Sleep(time.Duration(5) * time.Second)
							game = false
							i = 9
						}
					} else {
						userOneChoices = 0
						fmt.Print("That number is taken, try another one: ")
					}
				}
			}
		}
		game = false
	}
}
