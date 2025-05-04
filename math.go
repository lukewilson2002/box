package box

// Min returns the smaller of the two values.
func Min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

// Max returns the larger of the two values.
func Max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

// Clamp keeps the input value within a range of [min, max].
func Clamp(value, min, max int) int {
	return Max(min, Min(value, max))
}

type Rect struct {
	X, Y int
	W, H int
}

func (r Rect) Equals(other Rect) bool {
	return r.X == other.X && r.Y == other.Y && r.W == other.W && r.H == other.H
}

func (r Rect) HasPoint(x, y int) bool {
	return x >= r.X && y >= r.Y && x < r.X+r.W && y < r.Y+r.H
}

func (r Rect) Add(other Rect) Rect {
	return Rect{r.X + other.X, r.Y + other.Y, r.W + other.W, r.H + other.H}
}

func (r Rect) Sub(other Rect) Rect {
	return Rect{r.X - other.X, r.Y - other.Y, r.W - other.W, r.H - other.H}
}

func (r Rect) Pos() (x, y int) {
	return r.X, r.Y
}

func (r Rect) Size() (w, h int) {
	return r.W, r.H
}

func (r Rect) WithPos(x, y int) Rect {
	return Rect{x, y, r.W, r.H}
}

func (r Rect) WithSize(w, h int) Rect {
	return Rect{r.X, r.Y, w, h}
}

func (r Rect) Contract(amt int) Rect {
	w := r.W - amt*2
	h := r.H - amt*2
	if w < 0 || h < 0 {
		return r
	}
	return Rect{r.X + amt, r.Y + amt, w, h}
}

func (r Rect) Expand(amt int) Rect {
	return Rect{r.X - amt, r.Y - amt, r.W + amt*2, r.H + amt*2}
}
