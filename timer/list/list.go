package list

import (
	"time"

	"github.com/mattn/go-runewidth"
	"github.com/nsf/termbox-go"
)

const (
	BOARDER_DIS = '*'

	SET_YEAR int = iota
	SET_MON
	SET_DAY
	SET_H
	SET_M
	SET_S
)

type Box struct {
	Marginx int
	Marginy int
	Width   int
	Hight   int
	Fg, Bg  termbox.Attribute
}
type BoxTitle struct {
	Name         string
	Marginx      int
	Marginy      int
	Width        int
	Hight        int
	Fg, Bg       termbox.Attribute
	SelFg, SelBg termbox.Attribute
	SetItem      int
}
type List struct {
	Title   BoxTitle
	Boarder Box
	Content Box
	Msg     []string
}

func New(x, y int, w, h int, fg, bg termbox.Attribute) *List {
	for i := 0; i < w; i++ {
		termbox.SetCell(x+i, y, BOARDER_DIS, fg, bg)
		termbox.SetCell(x+i, y+h, BOARDER_DIS, fg, bg)
	}

	for i := 0; i < h; i++ {
		termbox.SetCell(x, y+i, BOARDER_DIS, fg, bg)
		termbox.SetCell(x+w, y+i, BOARDER_DIS, fg, bg)
	}

	var title BoxTitle
	nowtime := time.Now().Format("2006-01-02 15:04:05")
	title = BoxTitle{
		Name:    nowtime,
		Marginx: x + 4,
		Marginy: y - 2,
		Width:   w,
		Hight:   y,
		Fg:      fg,
		Bg:      bg,
		SelFg:   termbox.ColorGreen,
		SelBg:   bg,
	}

	showmsg(title.Marginx, title.Marginy, title.Fg, title.Bg, nowtime)

	var boarder Box
	boarder = Box{
		Marginx: x,
		Marginy: y,
		Width:   w,
		Hight:   y,
		Fg:      fg,
		Bg:      bg,
	}

	content := Box{
		Marginx: x + 4,
		Marginy: y + 1,
		Width:   w - 2,
		Hight:   y - 2,
		Fg:      fg,
		Bg:      bg,
	}
	return &List{
		Title:   title,
		Boarder: boarder,
		Content: content,
	}
}

func (c *List) Clear() {
	x := c.Content.Marginx
	y := c.Content.Marginy
	w := c.Content.Width
	h := c.Content.Hight
	fg := c.Content.Fg
	bg := c.Content.Bg
	for i := 0; i < w; i++ {
		for j := 0; j < h; j++ {
			termbox.SetCell(x+i, y+j, ' ', fg, bg)
		}
	}
}

func (c *List) Show() {
	msg := c.Msg
	if len(msg) <= 0 {
		return
	}
	x := c.Content.Marginx
	y := c.Content.Marginy
	fg := c.Content.Fg
	bg := c.Content.Bg
	for k, v := range msg {
		showmsg(x, y+k, fg, bg, v)
	}
}

func (c *List) TitleInit() {
	c.Title.Name = time.Now().Format("2006-01-02 15:04:05")
	showmsg(c.Title.Marginx, c.Title.Marginy, c.Title.Fg, c.Title.Bg, c.Title.Name)
}

func (c *List) TitleGetSelect() (setitem, item_width int) {
	if c.Title.SetItem > SET_S {
		c.Title.SetItem = SET_YEAR
	}
	switch c.Title.SetItem {
	case SET_YEAR:
		setitem = 0
		item_width = 4
	case SET_MON:
		setitem = 5
		item_width = 2
	case SET_DAY:
		setitem = 8
		item_width = 2
	case SET_H:
		setitem = 11
		item_width = 2
	case SET_M:
		setitem = 14
		item_width = 2
	case SET_S:
		setitem = 17
		item_width = 2
	}
	return
}

func (c *List) TitleSelect() {

	c.Title.SetItem++
	setitem, item_width := c.TitleGetSelect()
	showSetMsg(c.Title.Marginx, c.Title.Marginy, c.Title.Fg, c.Title.Bg, c.Title.SelFg, c.Title.SelBg, c.Title.Name, setitem, item_width)
}

func (c *List) TitleItemAdd() {
	loc, _ := time.LoadLocation("Local")
	settime, err := time.ParseInLocation("2006-01-02 15:04:05", c.Title.Name, loc)
	if err != nil {
		panic(err)
	}
	switch c.Title.SetItem {
	case SET_YEAR:
		settime = settime.AddDate(1, 0, 0)
	case SET_MON:
		settime = settime.AddDate(0, 1, 0)
	case SET_DAY:
		settime = settime.AddDate(0, 0, 1)
	case SET_H:
		settime = settime.Add(time.Hour)
	case SET_M:
		settime = settime.Add(time.Minute)
	case SET_S:
		settime = settime.Add(time.Second)
	}

	c.Title.Name = settime.Format("2006-01-02 15:04:05")
	setitem, item_width := c.TitleGetSelect()
	showSetMsg(c.Title.Marginx, c.Title.Marginy, c.Title.Fg, c.Title.Bg, c.Title.SelFg, c.Title.SelBg, c.Title.Name, setitem, item_width)
}

func (c *List) TitleItemDec() {
	loc, _ := time.LoadLocation("Local")
	settime, err := time.ParseInLocation("2006-01-02 15:04:05", c.Title.Name, loc)
	if err != nil {
		panic(err)
	}
	switch c.Title.SetItem {
	case SET_YEAR:
		settime = settime.AddDate(-1, 0, 0)
	case SET_MON:
		settime = settime.AddDate(0, -1, 0)
	case SET_DAY:
		settime = settime.AddDate(0, 0, -1)
	case SET_H:
		settime = settime.Add(-time.Hour)
	case SET_M:
		settime = settime.Add(-time.Minute)
	case SET_S:
		settime = settime.Add(-time.Second)
	}

	c.Title.Name = settime.Format("2006-01-02 15:04:05")
	setitem, item_width := c.TitleGetSelect()
	showSetMsg(c.Title.Marginx, c.Title.Marginy, c.Title.Fg, c.Title.Bg, c.Title.SelFg, c.Title.SelBg, c.Title.Name, setitem, item_width)
}

func (c *List) TitleNameToTime() time.Time {
	loc, _ := time.LoadLocation("Local")
	settime, err := time.ParseInLocation("2006-01-02 15:04:05", c.Title.Name, loc)
	if err != nil {
		panic(err)
	}
	return settime
}
func showmsg(x, y int, fg, bg termbox.Attribute, msg string) {
	for _, c := range msg {
		termbox.SetCell(x, y, c, fg, bg)
		x += runewidth.RuneWidth(c)
	}
}

func showSetMsg(x, y int, fg, bg, setfg, setbg termbox.Attribute, msg string, setitem, item_width int) {
	for k, c := range msg {
		if k >= setitem && k < setitem+item_width {
			termbox.SetCell(x, y, c, setfg, setbg)
		} else {
			termbox.SetCell(x, y, c, fg, bg)
		}

		x += runewidth.RuneWidth(c)
	}
}
