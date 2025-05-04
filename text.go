package box

import (
	"github.com/gdamore/tcell/v2"
	"github.com/mattn/go-runewidth"
	"strings"
)

// ConfineString inserts newlines where a line would run out of the rect,
// and trims the string to have no more lines than rows in the rect. Returns
// the formatted lines, the minimum columns to draw it, and the number of
// lines produced.
func ConfineString(s string, rect Rect, separator string) (lines []string, width, height int) {
	lines = strings.SplitN(s, separator, rect.H)
	for i := 0; i < len(lines); i++ {
		newLines := strings.Split(runewidth.Wrap(lines[i], rect.W), separator)
		if len(newLines) > 1 {
			// Append the number of lines we are adding
			lines = append(lines, make([]string, len(newLines)-1)...)
			// Shift every item down by len(newLines)
			copy(lines[Min(i+len(newLines), len(lines)):], lines[i+1:])
			copy(lines[i:], newLines) // Insert new lines in space
			i += len(newLines) - 1
		}
	}
	for i := 0; i < len(lines); i++ {
		if stringWidth := runewidth.StringWidth(lines[i]); stringWidth > width {
			width = stringWidth
			if width == rect.W { // I'm calling this an optimization...
				break
			}
		}
	}

	limit := Min(len(lines), rect.H)
	if limit >= 0 {
		lines = lines[:limit] // Keep line count only as large as rect height
	}
	return lines, width, len(lines)
}

type Text struct {
	Text      string
	Align     Alignment
	WrapLen   int    // Force the text to wrap after a specified number of terminal cells.
	Separator string // Empty string defaults to Unix linefeed "\n".
	Style     tcell.Style
}

var _ Widget = (*Text)(nil)

func (t *Text) GetSeparator() string {
	if len(t.Separator) == 0 {
		return "\n"
	} else {
		return t.Separator
	}
}

func (t *Text) GetChildren() []Widget {
	return nil
}

func (t *Text) HandleMouse(_ Rect, _ *tcell.EventMouse) bool {
	return false
}

func (t *Text) HandleKey(_ *tcell.EventKey) bool {
	return false
}

func (t *Text) SetFocused(_ bool) {}

func (t *Text) Bounds(space Rect) Rect {
	rect := space
	if t.WrapLen > 0 {
		rect.W = Min(t.WrapLen, rect.W)
	}
	_, w, h := ConfineString(t.Text, rect, t.GetSeparator())
	return space.WithSize(w, h)
}

func (t *Text) Draw(rect Rect, s tcell.Screen) {
	if t.WrapLen > 0 {
		rect.W = Min(t.WrapLen, rect.W)
	}
	lines, _, lineCount := ConfineString(t.Text, rect, t.GetSeparator())
	for i := 0; i < lineCount; i++ {
		switch t.Align {
		case AlignCenter:
			x := rect.X + rect.W/2 - runewidth.StringWidth(lines[i])/2
			DrawString(x, rect.Y+i, lines[i], t.Style, s)
		case AlignRight:
			x := rect.X + rect.W - runewidth.StringWidth(lines[i])
			DrawString(x, rect.Y+i, lines[i], t.Style, s)
		default:
			DrawString(rect.X, rect.Y+i, lines[i], t.Style, s)
		}
	}
}
