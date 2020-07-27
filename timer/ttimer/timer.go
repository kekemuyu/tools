package ttimer

import (
	"time"

	"github.com/kekemuyu/tools/timer/list"

	// "github.com/gen2brain/beeep"
	"github.com/nsf/termbox-go"
)

const (
	WIDTH = 30
	HIGHT = 60

	TEXT_WIDTH = WIDTH - 2
	TEXT_HIGHT = HIGHT - 2

	SET_TEXT_WIDTH = WIDTH
	SET_TEXT_HIGHT = 1
)

type Timer struct {
	Settime []time.Time
	Timeup  chan time.Time
}

func New() *Timer {
	return &Timer{
		Settime: make([]time.Time, 0),
		Timeup:  make(chan time.Time),
	}
}

func (c *Timer) Run() {
	go func() {
		for {
			if len(c.Settime) > 0 {
				for _, v := range c.Settime {
					if v.Equal(time.Now()) {
						c.Timeup <- v
					}
				}
			}
		}
	}()

mainloop:
	for {
		switch ev := termbox.PollEvent(); ev.Type {
		case termbox.EventKey:
			switch ev.Key {
			case termbox.KeyEsc:
				break mainloop
			case KeyTab:
			case KeyArrowUp:
			case KeyArrowDown:
			case KeySpace:
			}

		case termbox.EventError:
			panic(ev.Err)
		}

	}
}

func (c *Timer) Init(fg, bg termbox.Attribute) {
	err := termbox.Init()
	if err != nil {
		panic(err)
	}
	defer termbox.Close()
	termbox.SetInputMode(termbox.InputEsc)

	termbox.Clear(fg, bg)
	w, h := termbox.Size()
	midy := (h - HIGHT) / 2
	midx := (w - WIDTH) / 2

	settime_text_x := midx
	settime_text_y := midy - 2

	list.New(settime_text_x, settime_text_y, WIDTH, HIGHT, fg, bg)
	termbox.Flush()
}
