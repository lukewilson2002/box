package box

import (
	"fmt"
	"github.com/gdamore/tcell/v2"
	"reflect"
	"strings"
)

// Note: the order of the defined functions in the interface is the order that
// a component is recommended to implement the functions. That way the draw
// function, which is typically the longest, will be at the bottom.

type Widget interface {
	// GetChildren returns all children of the Widget. The returned slice is
	// owned by the Widget, so it should NOT be mutated. This is because some
	// Widgets implement this function by returning their internal slice of
	// child Widgets.
	GetChildren() []Widget

	// HandleMouse is called by a parent of the Widget when they receive the
	// event. A parent passes the event down to their child after attempting
	// to handle it, passing their child's accurate currentRect, which is
	// determined differently for every Widget, but is based upon the result of
	// the child's DisplaySize function. HandleMouse will return `true` if the
	// event is handled. Otherwise, `false`, so the event continues to be
	// propagated. If this Widget or any of its child Widgets handle this event
	// successfully (by returning true), then SetFocused(true) should be called.
	//
	// The currentRect is used by the Widget to determine its current position
	// and size on the terminal. The rect is determined by the Widget's parent.
	HandleMouse(currentRect Rect, ev *tcell.EventMouse) bool

	// HandleKey is called by a parent of the Widget when they receive the
	// event. The Widget will only try to handle the event if it is focused.
	// HandleKey will return `true` if the event is handled. Otherwise, `false`,
	// so the event can continue to be propagated. If this Widget or any of its
	// child Widgets handle this event successfully (by returning true), then
	// SetFocused(true) should be called.
	HandleKey(ev *tcell.EventKey) bool

	// SetFocused alerts the Widget that it has received input focus from the
	// user. The value can be kept in the Widget to differ its appearance during
	// Draw. The Widget will call SetFocused(b) on all of its children, also.
	SetFocused(b bool)

	// Bounds returns the exact position and size of the Widget when it will
	// be drawn This is used for containers like Center, and especially for the
	// HandleMouse function to work properly, as a Widget's position and size
	// will be determined by the result of calling its Bounds function.
	//
	// For nearly all Widgets, the resulting Rect should be within the Rect that
	// was provided. Otherwise, the Widget is drawing outside its permitted
	// space, which is probably a bug, unless intended.
	Bounds(space Rect) Rect

	// Draw renders the Widget onto the terminal screen, bounded by the provided
	// Rect. It is a bug if the Widget draws any part of itself outside the rect
	// provided. Draw should not call Sync() on the tcell.Screen or other
	// synchronizing functions, as all synchronization will be done by the event
	// loop.
	Draw(rect Rect, s tcell.Screen)
}

func Tree(w Widget) string {
	sb := new(strings.Builder)
	tree(w, sb, 0)
	return sb.String()
}

func tree(w Widget, sb *strings.Builder, depth int) {
	typeName := reflect.TypeOf(w).String()

	sb.WriteString(fmt.Sprintf("%s- %s\n",
		strings.Repeat("  ", depth), typeName))

	for _, child := range w.GetChildren() {
		tree(child, sb, depth+1)
	}
}
