package box

import (
	"github.com/gdamore/tcell/v2"
)

type Center struct {
	Child Widget
}

var _ Widget = (*Center)(nil)

func (c *Center) GetChildren() []Widget {
	return []Widget{c.Child}
}

func (c *Center) GetChildRect(space Rect) Rect {
	childRect := c.Child.Bounds(space)
	x, y := space.W/2-childRect.W/2, space.H/2-childRect.H/2
	return Rect{space.X + x, space.Y + y, childRect.W, childRect.H}
}

func (c *Center) HandleMouse(currentRect Rect, ev *tcell.EventMouse) bool {
	if c.Child != nil {
		return c.Child.HandleMouse(c.GetChildRect(currentRect), ev)
	} else {
		return false
	}
}

func (c *Center) HandleKey(ev *tcell.EventKey) bool {
	if c.Child != nil {
		return c.Child.HandleKey(ev)
	} else {
		return false
	}
}

func (c *Center) SetFocused(b bool) {
	if c.Child != nil {
		c.Child.SetFocused(b)
	}
}

func (c *Center) Bounds(space Rect) Rect {
	return space
}

func (c *Center) Draw(rect Rect, s tcell.Screen) {
	if c.Child != nil {
		c.Child.Draw(c.GetChildRect(rect), s)
	}
}
