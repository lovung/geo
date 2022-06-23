package geo

import (
	"fmt"
	"math"
)

// Line in 2D is the collection (set) of all points which have
// a*x + b*y + c = 0
// https://en.wikipedia.org/wiki/Line_(geometry)
type Line struct {
	a, b, c float64
}

// NewLine create the new line which through 2 points
func NewLine(m, n Point) Line {
	vector := n.Sub(m)
	return Line{-vector.y, vector.x, (m.x * vector.y) - (m.y * vector.x)}
}

func (l Line) String() string {
	format := "%vx "
	if l.b < 0 {
		format += "- %vy "
	} else {
		format += "+ %vy "
	}
	if l.c < 0 {
		format += "- %v"
	} else {
		format += "+ %v"
	}
	return fmt.Sprintf(format+" = 0", l.a, math.Abs(l.b), math.Abs(l.c))
}
