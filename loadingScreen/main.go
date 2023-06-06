package loadingScreen

import (
	"fmt"
	"math/rand"
	"time"

	"01.kood.tech/git/jsaar/tic_tac_toe/screenClear"
)

var loadingVisuals map[int]string = map[int]string{
	1: "┌────────────────────┐",
	2: "│                    │",
	3: "└────────────────────┘",
	4: "|",
}

func Load() {
	for i := 1; i < 21; i++ {
		screenClear.Now()
		loadingVisuals[2] = loadingVisuals[2][:i+2] + loadingVisuals[4] + loadingVisuals[2][i+3:]
		fmt.Println(loadingVisuals[1])
		fmt.Println(loadingVisuals[2])
		fmt.Println(loadingVisuals[3])
		time.Sleep(time.Duration(rand.Intn(200) * int(time.Millisecond)))
	}
}
