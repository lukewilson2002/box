package box

type Positioning uint8

const (
	// Inherit is the default layout for all Widgets. The Rect property will be
	// ignored. Calling Align's Bounds will return Bounds on the
	// Child.
	Inherit Positioning = iota

	// Absolute positioning causes a Widget to be placed at any X, Y coordinate
	// with any arbitrary width and height as specified. This is useful for
	// drop-down menus or other floating widgets. Calling Align's Bounds
	// will return zero.
	Absolute

	// Relative positioning is similar to Absolute, but causes a Widget to
	// inherit its parent's position. Calling Align's Bounds will return
	// zero.
	Relative
)

type Alignment uint8

const (
	AlignLeft Alignment = iota
	AlignRight
	AlignCenter
	AlignStart = AlignLeft
	AlignEnd   = AlignRight
)
