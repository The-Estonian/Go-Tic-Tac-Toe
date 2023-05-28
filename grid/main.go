package grid

import (
	"fmt"
	"os"
	"os/exec"
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
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
	for i := 1; i <= 16; i++ {
		fmt.Println(grid[i])
	}
}

var boxCoordY = [][]int{{4, 12}, {17, 25}, {30, 38}}
var boxCoordX = [][]int{{2, 3, 4, 5}, {7, 8, 9, 10}, {12, 13, 14, 15}}

func PlacePlayer(player string, box int) {
	selectedBoxY := []int{}
	selectedBoxX := []int{}
	switch box {
	case 1:
		selectedBoxY = boxCoordY[0]
		selectedBoxX = boxCoordX[0]
	case 2:
		selectedBoxY = boxCoordY[1]
		selectedBoxX = boxCoordX[0]
	case 3:
		selectedBoxY = boxCoordY[2]
		selectedBoxX = boxCoordX[0]
	case 4:
		selectedBoxY = boxCoordY[0]
		selectedBoxX = boxCoordX[1]
	case 5:
		selectedBoxY = boxCoordY[1]
		selectedBoxX = boxCoordX[1]
	case 6:
		selectedBoxY = boxCoordY[2]
		selectedBoxX = boxCoordX[1]
	case 7:
		selectedBoxY = boxCoordY[0]
		selectedBoxX = boxCoordX[2]
	case 8:
		selectedBoxY = boxCoordY[1]
		selectedBoxX = boxCoordX[2]
	case 9:
		selectedBoxY = boxCoordY[2]
		selectedBoxX = boxCoordX[2]
	}

	if player == "o" {
		grid[selectedBoxX[0]] = grid[selectedBoxX[0]][:selectedBoxY[0]] + grid[21] + grid[selectedBoxX[0]][selectedBoxY[1]:]
		grid[selectedBoxX[1]] = grid[selectedBoxX[1]][:selectedBoxY[0]] + grid[22] + grid[selectedBoxX[1]][selectedBoxY[1]:]
		grid[selectedBoxX[2]] = grid[selectedBoxX[2]][:selectedBoxY[0]] + grid[23] + grid[selectedBoxX[2]][selectedBoxY[1]:]
		grid[selectedBoxX[3]] = grid[selectedBoxX[3]][:selectedBoxY[0]] + grid[24] + grid[selectedBoxX[3]][selectedBoxY[1]:]
	} else if player == "x" {
		grid[selectedBoxX[0]] = grid[selectedBoxX[0]][:selectedBoxY[0]] + grid[17] + grid[selectedBoxX[0]][selectedBoxY[1]:]
		grid[selectedBoxX[1]] = grid[selectedBoxX[1]][:selectedBoxY[0]] + grid[18] + grid[selectedBoxX[1]][selectedBoxY[1]:]
		grid[selectedBoxX[2]] = grid[selectedBoxX[2]][:selectedBoxY[0]] + grid[19] + grid[selectedBoxX[2]][selectedBoxY[1]:]
		grid[selectedBoxX[3]] = grid[selectedBoxX[3]][:selectedBoxY[0]] + grid[20] + grid[selectedBoxX[3]][selectedBoxY[1]:]
	}
}
