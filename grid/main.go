package grid

import (
	"fmt"
	"os"
	"os/exec"
)

var a = ("┌──────────┬──────────┬──────────┐")
var b = ("│          │          │          │")
var c = ("│          │          │          │")
var d = ("│          │          │          │")
var e = ("│          │          │          │")
var f = ("├──────────┼──────────┼──────────┤")
var g = ("│          │          │          │")
var h = ("│          │          │          │")
var i = ("│          │          │          │")
var j = ("│          │          │          │")
var k = ("├──────────┼──────────┼──────────┤")
var l = ("│          │          │          │")
var m = ("│          │          │          │")
var n = ("│          │          │          │")
var o = ("│          │          │          │")
var p = ("└──────────┴──────────┴──────────┘")

var a1 = ("  ╲  ╱  ")
var a2 = ("   ╲╱   ")
var a3 = ("   ╱╲   ")
var a4 = ("  ╱  ╲  ")

var b1 = ("╭──────╮")
var b2 = ("│      │")
var b3 = ("│      │")
var b4 = ("╰──────╯")

func FullGrid() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
	gridList := []string{a, b, c, d, e, f, g, h, i, j, k, l, m, n, o, p}
	for i := 0; i < len(gridList); i++ {
		fmt.Println(gridList[i])
	}

	fmt.Println("----------------")
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)
}
var boxCoordY = [][]int{{4, 12}, {17, 25}, {30, 38}}

func PlacePlayer(player string, box int) {
	selectedBox := []int{}
	switch box{
	case 1:
		selectedBox = boxCoordY[0]
	case 2:
		selectedBox = boxCoordY[1]
	case 3:
		selectedBox = boxCoordY[2]
	}

	if player == "o" {
		b = b[:selectedBox[0]] + b1 + b[selectedBox[1]:]
		c = c[:selectedBox[0]] + b2 + c[selectedBox[1]:]
		d = d[:selectedBox[0]] + b3 + d[selectedBox[1]:]
		e = e[:selectedBox[0]] + b4 + e[selectedBox[1]:]
	} else if player == "x" {
		b = b[:selectedBox[0]] + a1 + b[selectedBox[1]:]
		c = c[:selectedBox[0]] + a2 + c[selectedBox[1]:]
		d = d[:selectedBox[0]] + a3 + d[selectedBox[1]:]
		e = e[:selectedBox[0]] + a4 + e[selectedBox[1]:]
	}
}
