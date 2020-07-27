package main

import (
	"fmt"
	"tools/timer/timer"

	"github.com/gen2brain/beeep"
	"github.com/nsf/termbox-go"
)

func Beep(n int) {
	for i := 0; i < n; i++ {
		err := beeep.Beep(beeep.DefaultFreq, beeep.DefaultDuration)
		if err != nil {
			panic(err)
		}
	}
}

func main() {
	timer.New()
}
