package dos

import "github.com/gdamore/tcell/v2"

func DrawBox(rect Rect, decoration *dos.BoxDecoration, s tcell.Screen) {
	// Draw top left and bottom left corners
	s.SetContent(rect.X, rect.Y, decoration.TL, nil, decoration.Style)
	s.SetContent(rect.X, rect.Y+rect.H-1, decoration.BL, nil, decoration.Style)
	// Draw top right and bottom right corners
	s.SetContent(rect.X+rect.W-1, rect.Y, decoration.TR, nil, decoration.Style)
	s.SetContent(rect.X+rect.W-1, rect.Y+rect.H-1, decoration.BR, nil, decoration.Style)
	// Draw top and bottom sides
	for col := 1; col < rect.W-1; col++ {
		s.SetContent(rect.X+col, rect.Y, decoration.Hor, nil, decoration.Style)
		s.SetContent(rect.X+col, rect.Y+rect.H-1, decoration.Hor, nil, decoration.Style)
	}
	// Draw left and right sides
	for row := 1; row < rect.H-1; row++ {
		s.SetContent(rect.X, rect.Y+row, decoration.Vert, nil, decoration.Style)
		s.SetContent(rect.X+rect.W-1, rect.Y+row, decoration.Vert, nil, decoration.Style)
	}
}
