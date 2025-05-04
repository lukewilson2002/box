package box

import (
	"math"
	"testing"
)

func TestMin(t *testing.T) {
	if Min(1, 5) != 1 {
		t.Fail()
	}
	if Min(math.MinInt, math.MaxInt) != math.MinInt {
		t.Fail()
	}
	if Min(math.MaxInt, math.MinInt) != math.MinInt {
		t.Fail()
	}
}

func TestMax(t *testing.T) {
	if Max(1, 5) != 5 {
		t.Fail()
	}
	if Max(math.MinInt, math.MaxInt) != math.MaxInt {
		t.Fail()
	}
	if Max(math.MaxInt, math.MinInt) != math.MaxInt {
		t.Fail()
	}
}

func TestClamp(t *testing.T) {
	if Clamp(3, 1, 5) != 3 {
		t.Fail()
	}
	if Clamp(-1, 0, 1) != 0 {
		t.Fail()
	}
	if Clamp(math.MaxInt, 0, 15) != 15 {
		t.Fail()
	}
	if Clamp(math.MaxInt, math.MaxInt, math.MaxInt) != math.MaxInt {
		t.Fail()
	}
	if Clamp(math.MinInt, math.MinInt, math.MinInt) != math.MinInt {
		t.Fail()
	}
}

func TestRect_Equals(t *testing.T) {
	a := Rect{0, 0, 0, 0}
	b := Rect{0, 0, 0, 0}
	if !a.Equals(b) {
		t.Fail()
	}
	if !b.Equals(a) {
		t.Fail()
	}
	b.W = 1
	if a.Equals(b) {
		t.Fail()
	}
}

func TestRect_HasPoint(t *testing.T) {
	if !(Rect{0, 0, 1, 1}.HasPoint(0, 0)) {
		t.Fail()
	}
	if (Rect{1, 1, 1, 1}.HasPoint(0, 0)) {
		t.Fail()
	}
	if (Rect{1, 1, 0, 0}.HasPoint(1, 1)) {
		t.Fail()
	}
	wide := Rect{3, 2, 15, 4}
	if wide.HasPoint(5, 1) {
		t.Fail()
	}
	if wide.HasPoint(8, 7) {
		t.Fail()
	}
	if !wide.HasPoint(5, 3) {
		t.Fail()
	}
}

func TestRect_Add(t *testing.T) {
	a := Rect{0, 0, 0, 1}
	b := Rect{0, 0, 0, 1}
	c := a.Add(b)
	if c.X != 0 || c.Y != 0 || c.W != 0 || c.H != 2 {
		t.Fail()
	}

	b = Rect{2, -3, math.MaxInt, math.MinInt}
	c = a.Add(b)
	if c.X != 2 || c.Y != -3 || c.W != math.MaxInt || c.H != math.MinInt+1 {
		t.Fail()
	}
}

func TestRect_Sub(t *testing.T) {
	a := Rect{0, 0, 0, -1}
	b := Rect{0, 0, 0, 1}
	c := a.Sub(b)
	if c.X != 0 || c.Y != 0 || c.W != 0 || c.H != -2 {
		t.Fail()
	}
	// Test ordinality
	c = b.Sub(a)
	if c.X != 0 || c.Y != 0 || c.W != 0 || c.H != 2 {
		t.Fail()
	}

	a = Rect{4, 0, 5, 0}
	b = Rect{2, -3, 100, 0}
	c = a.Sub(b)
	if c.X != 2 || c.Y != 3 || c.W != -95 || c.H != 0 {
		t.Fail()
	}
}

func TestRect_WithPos(t *testing.T) {
	a := Rect{0, 3, 0, 1}
	b := a.WithPos(5, 2)
	if b.X != 5 || b.Y != 2 || b.W != 0 || b.H != 1 {
		t.Fail()
	}
	if a.X != 0 || a.Y != 3 || a.W != 0 || a.H != 1 {
		t.Fail()
	}
}

func TestRect_WithSize(t *testing.T) {
	a := Rect{0, 3, 0, 1}
	b := a.WithSize(6, 10)
	if b.X != 0 || b.Y != 3 || b.W != 6 || b.H != 10 {
		t.Fail()
	}
	if a.X != 0 || a.Y != 3 || a.W != 0 || a.H != 1 {
		t.Fail()
	}
}

func TestRect_Contract(t *testing.T) {
	// Keep in mind the top left cell is (0, 0)
	a := Rect{1, 1, 4, 4}
	expected := Rect{2, 2, 2, 2}
	if !a.Contract(1).Equals(expected) {
		t.Fail()
	}
	expected = Rect{3, 3, 0, 0}
	if !a.Contract(2).Equals(expected) {
		t.Fail()
	}
	// Cannot contract a rect without a size
	expected2 := expected.Contract(2)
	if !expected2.Equals(expected) {
		t.Fail()
	}
}

func TestRect_Expand(t *testing.T) {
	// Keep in mind the top left cell is (0, 0)
	a := Rect{1, 1, 4, 4}
	expected := Rect{0, 0, 6, 6}
	if !a.Expand(1).Equals(expected) {
		t.Fail()
	}
	expected = Rect{-2, -2, 10, 10}
	if !a.Expand(3).Equals(expected) {
		t.Fail()
	}
}
