package main

import (
	"fmt"
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

	// Here we construct the GUI
	widget := &box.Center{
		Child: &box.Label{
			Text:  "Hello, world!\nTry resizing the screen!",
			Style: tcell.StyleDefault.Foreground(tcell.ColorRed),
		},
	}

	var app box.App
	app = box.App{
		MainWidget: widget,
		OnKeyEvent: func(ev *tcell.EventKey) bool {
			if ev.Key() == tcell.KeyEsc {
				app.Running = false // Stops the app and restores the terminal
				return true         // Prevent the event from being passed onto widgets
			}
			return false
		},
	}
	app.Run(screen)
}
