package geo

import "strconv"

// Point reprensent a point in 2D
type Point struct {
	X, Y float64
}

func (p *Point) Sub(p2 Point) Vector {
	return Vector{p.X - p2.X, p.Y - p2.Y, false}
}

func (p Point) String() string {
	return "(" + strconv.FormatFloat(p.X, 'f', -1, 64) + ", " + strconv.FormatFloat(p.Y, 'f', -1, 64) + ")"
}
