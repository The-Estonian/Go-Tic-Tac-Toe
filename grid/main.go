package grid

import (
	"fmt"
	"01.kood.tech/git/jsaar/tic_tac_toe/screenClear"
)

var grid map[int]string = map[int]string{
	1:  "┌──────────┬──────────┬──────────┐",
	2:  "│          │          │          │",
	3:  "│          │          │          │",
	4:  "│          │          │          │",
	5:  "│          │          │          │",
	6:  "├──────────┼──────────┼──────────┤",
	7:  "│          │          │          │",
	8:  "│          │          │          │",
	9:  "│          │          │          │",
	10: "│          │          │          │",
	11: "├──────────┼──────────┼──────────┤",
	12: "│          │          │          │",
	13: "│          │          │          │",
	14: "│          │          │          │",
	15: "│          │          │          │",
	16: "└──────────┴──────────┴──────────┘",
	17: "  \\  /  ",
	18: "   \\/   ",
	19: "   /\\   ",
	20: "  /  \\  ",
	21: "+------+",
	22: "|      |",
	23: "|      |",
	24: "+------+",
}

func PrintGrid() {
	screenClear.Now()
	for i := 1; i <= 16; i++ {
		fmt.Println(grid[i])
	}
}

var boxCoordY = [][]int{{4, 12}, {17, 25}, {30, 38}}
var boxCoordX = [][]int{{2, 3, 4, 5}, {7, 8, 9, 10}, {12, 13, 14, 15}}

func PlacePlayer(player string, box int) {
	boxYaxis := []int{}
	boxXaxis := []int{}
	switch box {
	case 1:
		boxYaxis = boxCoordY[0]
		boxXaxis = boxCoordX[0]
	case 2:
		boxYaxis = boxCoordY[1]
		boxXaxis = boxCoordX[0]
	case 3:
		boxYaxis = boxCoordY[2]
		boxXaxis = boxCoordX[0]
	case 4:
		boxYaxis = boxCoordY[0]
		boxXaxis = boxCoordX[1]
	case 5:
		boxYaxis = boxCoordY[1]
		boxXaxis = boxCoordX[1]
	case 6:
		boxYaxis = boxCoordY[2]
		boxXaxis = boxCoordX[1]
	case 7:
		boxYaxis = boxCoordY[0]
		boxXaxis = boxCoordX[2]
	case 8:
		boxYaxis = boxCoordY[1]
		boxXaxis = boxCoordX[2]
	case 9:
		boxYaxis = boxCoordY[2]
		boxXaxis = boxCoordX[2]
	}

	if player == "o" {
		grid[boxXaxis[0]] = grid[boxXaxis[0]][:boxYaxis[0]] + grid[21] + grid[boxXaxis[0]][boxYaxis[1]:]
		grid[boxXaxis[1]] = grid[boxXaxis[1]][:boxYaxis[0]] + grid[22] + grid[boxXaxis[1]][boxYaxis[1]:]
		grid[boxXaxis[2]] = grid[boxXaxis[2]][:boxYaxis[0]] + grid[23] + grid[boxXaxis[2]][boxYaxis[1]:]
		grid[boxXaxis[3]] = grid[boxXaxis[3]][:boxYaxis[0]] + grid[24] + grid[boxXaxis[3]][boxYaxis[1]:]
	} else if player == "x" {
		grid[boxXaxis[0]] = grid[boxXaxis[0]][:boxYaxis[0]] + grid[17] + grid[boxXaxis[0]][boxYaxis[1]:]
		grid[boxXaxis[1]] = grid[boxXaxis[1]][:boxYaxis[0]] + grid[18] + grid[boxXaxis[1]][boxYaxis[1]:]
		grid[boxXaxis[2]] = grid[boxXaxis[2]][:boxYaxis[0]] + grid[19] + grid[boxXaxis[2]][boxYaxis[1]:]
		grid[boxXaxis[3]] = grid[boxXaxis[3]][:boxYaxis[0]] + grid[20] + grid[boxXaxis[3]][boxYaxis[1]:]
	}
}
