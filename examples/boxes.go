package main

import (
	"fmt"
	"github.com/fivemoreminix/box/dos"
	"os"

	"github.com/fivemoreminix/box"
	"github.com/gdamore/tcell/v2"
)

func main() {
	screen, err := tcell.NewScreen()
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to create tcell screen: %v", err)
	}
	if err = screen.Init(); err != nil {
		fmt.Fprintf(os.Stderr, "failed to initialize: %v", err)
	}

	var app box.App
	app = box.App{
		MainWidget: NewMainWidget(),
		OnKeyEvent: func(ev *tcell.EventKey) bool {
			if ev.Key() == tcell.KeyEsc {
				app.Running = false // Stop the app
				return true         // Report we handled the event
			}
			return false
		},
	}
	app.Run(screen)
}

type MainWidget struct {
	align *box.Align
	text  *box.Text
}

var _ box.Widget = (*MainWidget)(nil)

func NewMainWidget() *MainWidget {
	text := &box.Text{
		Text: "Type something!",
	}
	return &MainWidget{
		align: &box.Align{
			Child: &dos.Box{
				Child: text,
			},
			Positioning: box.Absolute,
		},
		text: text,
	}
}

func (m *MainWidget) GetChildren() []box.Widget {
	return []box.Widget{m.align, m.text}
}

func (m *MainWidget) HandleMouse(currentRect box.Rect, ev *tcell.EventMouse) bool {
	curX, curY := ev.Position()
	m.align.Rect.X, m.align.Rect.Y = curX-m.align.Rect.W, curY-m.align.Rect.H
	return true
}

func (m *MainWidget) HandleKey(ev *tcell.EventKey) bool {
	if ev.Key() == tcell.KeyBS || ev.Key() == tcell.KeyDEL {
		// Delete the character at the end
		if len(m.text.Text) > 0 {
			m.text.Text = m.text.Text[:len(m.text.Text)-1]
		}
	} else {
		// Insert the typed character at the end
		m.text.Text = string(append([]rune(m.text.Text), ev.Rune()))
	}
	return true
}

func (m *MainWidget) SetFocused(b bool) {
	m.align.SetFocused(b)
}

func (m *MainWidget) Bounds(space box.Rect) box.Rect {
	return m.align.Bounds(space)
}

func (m *MainWidget) Draw(rect box.Rect, s tcell.Screen) {
	// Because the Align is set to Absolute positioning, we give it a position and size through m.align.Rect
	// I directly access the align's child to know how big it plans to be, because the align will always
	// return 0,0 for an absolute size.
	space := rect.WithSize(rect.W-m.align.Rect.X, rect.H-m.align.Rect.Y)
	m.align.Rect.W, m.align.Rect.H = m.align.Child.Bounds(space).Size()
	m.align.Draw(rect, s)
}
