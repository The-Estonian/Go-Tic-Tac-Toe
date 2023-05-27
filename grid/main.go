package grid

import (
	"fmt"
	"os/exec"
	"os"
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
}

func PlacePlayer(player string) {
	if player == "o" {
		b = b[:4] + b1 + b[12:]
		c = c[:4] + b2 + c[12:]
		d = d[:4] + b3 + d[12:]
		e = e[:4] + b4 + e[12:]
	} else {
		b = b[:4] + a1 + b[12:]
		c = c[:4] + a2 + c[12:]
		d = d[:4] + a3 + d[12:]
		e = e[:4] + a4 + e[12:]
	}
}
