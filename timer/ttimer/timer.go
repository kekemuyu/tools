package ttimer

import (
	"time"

	"timer/list"

	"github.com/gen2brain/beeep"
	"github.com/mattn/go-runewidth"
	"github.com/nsf/termbox-go"
)

const (
	WIDTH = 30
	HIGHT = 10

	TEXT_WIDTH = WIDTH - 2
	TEXT_HIGHT = HIGHT - 2

	SET_TEXT_WIDTH = WIDTH
	SET_TEXT_HIGHT = 1
)

type Timer struct {
	TimerList *list.List
	Settime   []time.Time
	Timeup    chan time.Time
}

func New() *Timer {
	return &Timer{
		Settime: make([]time.Time, 0),
		Timeup:  make(chan time.Time, 10),
	}
}

func (c *Timer) Run() {
	defer termbox.Close()
	c.Init()
	list := c.TimerList
	go func() {
		for {
			now := time.Now()

			if len(c.Settime) > 0 {
				for k, v := range c.Settime {

					if v.Before(now) || v.Equal(now) {
						if k == 0 {
							c.Settime = c.Settime[1:]
						} else if k == len(c.Settime)-1 {
							c.Settime = c.Settime[:len(c.Settime)-1]
						} else if k > 0 && k < len(c.Settime) {
							c.Settime = append(c.Settime[0:k], c.Settime[k+1:]...)
						}
						msg := make([]string, 0)
						for _, v := range c.Settime {
							msg = append(msg, v.Format("2006-01-02 15:04:05"))
						}
						list.Msg = msg
						list.Clear()
						list.Show()

						c.Timeup <- v
					}
				}
			}
			termbox.Flush()
		}
	}()

	go func() {
		for {
			select {
			case <-c.Timeup:
				Beep(30)
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
			case termbox.KeyArrowLeft, termbox.KeyArrowRight:

				list.TitleSelect()
				termbox.Flush()
			case termbox.KeyArrowUp:
				list.TitleItemAdd()
				termbox.Flush()
			case termbox.KeyArrowDown:
				list.TitleItemDec()
				termbox.Flush()
			case termbox.KeySpace:
				settime := list.TitleNameToTime()
				c.Settime = append(c.Settime, settime)

				msg := make([]string, 0)
				for _, v := range c.Settime {
					msg = append(msg, v.Format("2006-01-02 15:04:05"))
				}
				list.Msg = msg
				list.Clear()
				list.Show()
				termbox.Flush()
			}

		case termbox.EventError:
			panic(ev.Err)
		}

	}
}

func (c *Timer) Init() {
	fg, bg := termbox.ColorDefault, termbox.ColorDefault
	err := termbox.Init()
	if err != nil {
		panic(err)
	}

	termbox.SetInputMode(termbox.InputEsc)

	termbox.Clear(fg, bg)
	w, h := termbox.Size()
	midy := (h - HIGHT) / 2
	midx := (w - WIDTH) / 2

	settime_text_x := midx
	settime_text_y := midy - 2

	c.TimerList = list.New(settime_text_x, settime_text_y, WIDTH, HIGHT, fg, bg)
	c.TimerList.TitleInit()
	termbox.Flush()
}

func Beep(n int) {
	for i := 0; i < n; i++ {
		err := beeep.Beep(beeep.DefaultFreq, beeep.DefaultDuration)
		if err != nil {
			panic(err)
		}
	}
}

func showmsg(x, y int, fg, bg termbox.Attribute, msg string) {
	for _, c := range msg {
		termbox.SetCell(x, y, c, fg, bg)
		x += runewidth.RuneWidth(c)
	}
}
