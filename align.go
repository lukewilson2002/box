package box

import (
	"github.com/gdamore/tcell/v2"
	"log"
)

// TODO: test and improve this component
type Align struct {
	Child       Widget
	Positioning Positioning
	Rect        Rect // Rect of Child if Positioning is Absolute.
}

var _ Widget = (*Align)(nil)

func (a *Align) GetChildBounds(rect Rect) Rect {
	switch a.Positioning {
	case Absolute:
		return a.Rect
	case Relative:
		return a.Rect.Add(Rect{X: rect.X, Y: rect.Y}) // Offset
	default:
		return rect
	}
}

func (a *Align) GetChildren() []Widget {
	return []Widget{a.Child}
}

func (a *Align) HandleMouse(rect Rect, ev *tcell.EventMouse) bool {
	log.Println("Align", rect)
	if a.Child != nil {
		return a.Child.HandleMouse(a.Child.Bounds(rect), ev)
	}
	return false
}

func (a *Align) HandleKey(ev *tcell.EventKey) bool {
	if a.Child != nil {
		return a.Child.HandleKey(ev)
	}
	return false
}

func (a *Align) SetFocused(b bool) {
	if a.Child != nil {
		a.Child.SetFocused(b)
	}
}

func (a *Align) Bounds(space Rect) Rect {
	if a.Child != nil {
		switch a.Positioning {
		case Absolute:
			// FIXME: In the future,
			// when all functions take "space" and not actual rect,
			//make this space.WithSize(0,
			//0) because Absolute Align has no size, but the child DOES
			//return space.WithSize(0, 0)

			return a.Rect
		case Relative:
			fallthrough
		default:
			return a.Child.Bounds(space)
		}
	}
	return space.WithSize(0, 0)
}

func (a *Align) Draw(rect Rect, s tcell.Screen) {
	if a.Child != nil {
		a.Child.Draw(a.GetChildBounds(rect), s)
	}
}
