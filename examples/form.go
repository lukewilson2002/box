package main

import (
	"fmt"
	"github.com/fivemoreminix/box"
	"github.com/gdamore/tcell/v2"
	"os"
)

func main() {
	screen, err := tcell.NewScreen()
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to create tcell screen: %v", err)
	}
	if err = screen.Init(); err != nil {
		fmt.Fprintf(os.Stderr, "failed to initialize: %v", err)
	}

	widget := &box.Row{
		Children: []box.Widget{
			&box.Column{ // Labels
				Children: []box.Widget{
					&box.Label{Text: "Username: "},
					&box.Label{Text: "Password: "},
					&box.Label{Text: "Favorite number: "},
				},
				HorizontalAlign: box.AlignRight,
			}, // Fields
			&box.Column{
				Children: []box.Widget{
					// TODO: input fields
				},
				OnKeyEvent: func(col *box.Column, ev *tcell.EventKey) bool {
					if ev.Key() == tcell.KeyTab {
						col.FocusNext()
						return true
					}
					return false
				},
			},
		},
		FocusedIndex: 1,
	}

	var app box.App
	app = box.App{
		MainWidget: widget,
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
