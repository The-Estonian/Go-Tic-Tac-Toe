package grid

import (
	"fmt"
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

var a1 = ("╲  ╱")
var a2 = (" ╲╱ ")
var a3 = (" ╱╲ ")
var a4 = ("╱  ╲")

var b1 = ("┌─────┐")
var b2 = ("│     │")
var b3 = ("│     │")
var b4 = ("└─────┘")

func FullGrid() {
	fmt.Println("Grid package Full Grid func")
	gridList := []string{a, b, c, d, e, f, g, h, i, j, k, l, m, n, o, p}
	for i := 0; i < len(gridList); i++ {
		fmt.Println(gridList[i])
	}
	fmt.Println(a1)
	fmt.Println(a2)
	fmt.Println(a3)
	fmt.Println(a4)

	fmt.Println(b1)
	fmt.Println(b2)
	fmt.Println(b3)
	fmt.Println(b4)
}
