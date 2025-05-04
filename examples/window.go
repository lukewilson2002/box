package main

import (
	"fmt"
	"github.com/fivemoreminix/box/dos"
	"os"

	"github.com/fivemoreminix/box"
	"github.com/gdamore/tcell/v2"
)

var (
	defaultStyle = tcell.Style{}.Background(tcell.ColorBlue).Foreground(tcell.ColorGrey)
	windowStyle  = tcell.Style{}.Background(tcell.ColorLightBlue).Foreground(tcell.ColorBlack)
	quote        = `Are you quite sure that all those bells and whistles,
all those wonderful facilities of your so called powerful programming languages,
belong to the solution set rather than the problem set?

 — Edsger W. Dijkstra`
)

func main() {
	screen, err := tcell.NewScreen()
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to create tcell screen: %v", err)
	}
	if err = screen.Init(); err != nil {
		fmt.Fprintf(os.Stderr, "failed to initialize: %v", err)
	}

	label := &box.Text{
		Text:  quote,
		Align: box.AlignLeft,
		Style: defaultStyle,
	}

	primary := &dos.Scaffold{
		MenuBar:    nil,
		MainWidget: &box.Center{Child: label},
		Floating:   []box.Widget{NewMainWindow(label)},
	}

	var app box.App
	app = box.App{
		ClearStyle: defaultStyle,
		MainWidget: primary,
		OnKeyEvent: func(ev *tcell.EventKey) bool {
			if ev.Key() == tcell.KeyEsc {
				app.Running = false
				return true
			}
			return false
		},
	}
	app.Run(screen)
}

func MakeDialog(title string, rect box.Rect, child box.Widget) box.Widget {
	align := &box.Align{
		Child:       nil,
		Positioning: box.Absolute,
		Rect:        rect,
	}
	align.Child = &dos.Shadow{
		Child: &dos.Window{
			Title:            title,
			Child:            child,
			HideClose:        true,
			OnClosed:         nil,
			DisableMoving:    false,
			OnMove:           func(posX, posY int) { align.Rect.X = posX; align.Rect.Y = posY },
			CloseButtonStyle: tcell.Style{}.Background(tcell.ColorRed).Foreground(tcell.ColorBlack),
			TitleBarStyle:    windowStyle.Background(tcell.ColorWhite),
			WindowStyle:      windowStyle,
		},
		Style:     tcell.Style{}.Background(tcell.ColorDarkCyan).Foreground(tcell.ColorBlack),
		MakeSmall: false,
	}
	return align
}

func makeButton(title string, action func()) box.Widget {
	return &box.Padding{
		Child: &dos.Shadow{
			Child: &dos.Button{
				Text:         title,
				NormalStyle:  windowStyle.Background(tcell.ColorWhite),
				FocusedStyle: windowStyle.Background(tcell.ColorWhite),
				OnPressed:    action,
			},
			Style:     tcell.Style{}.Background(tcell.ColorTeal).Foreground(tcell.ColorLightBlue),
			MakeSmall: true,
		},
		Top:    1,
		Right:  1,
		Bottom: 0,
		Left:   1,
	}
}

func NewMainWindow(label *box.Text) box.Widget {
	row := &box.Row{
		Children: []box.Widget{
			makeButton("Left align", func() { label.Align = box.AlignLeft }),
			makeButton("Center align", func() { label.Align = box.AlignCenter }),
			makeButton("Right align", func() { label.Align = box.AlignRight }),
		},
		OnKeyEvent: func(row *box.Row, ev *tcell.EventKey) bool {
			switch ev.Key() {
			case tcell.KeyRight:
				fallthrough
			case tcell.KeyTab:
				row.FocusNext()
			case tcell.KeyLeft:
				row.FocusPrevious()
			default:
				return false
			}
			return true
		},
	}
	return MakeDialog(
		"Text Alignment",
		box.Rect{5, 3, 53, 6},
		&box.Center{
			Child: &box.Column{
				Children: []box.Widget{
					&box.Padding{
						Child: &box.Text{
							Text:  "Choose a text alignment",
							Style: windowStyle,
						},
						Top: 1,
					},
					row,
				},
				HorizontalAlign: box.AlignCenter,
				FocusedIndex:    1, // Focus row only
			},
		},
	)
}
