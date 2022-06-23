package geo

import "strconv"

// Point reprensent a point in 2D
type Point struct {
	x, y float64
}

func (p *Point) Sub(p2 Point) Vector {
	return Vector{p.x - p2.x, p.y - p2.y}
}

// Move to a new point base on a vector
func (p *Point) Move(v Vector) Point {
	return Point{p.x + v.x, p.y + v.y}
}

func (p Point) String() string {
	return "(" + strconv.FormatFloat(p.x, 'f', -1, 64) + ", " + strconv.FormatFloat(p.y, 'f', -1, 64) + ")"
}
