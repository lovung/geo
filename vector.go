package geo

import (
	"math"
	"strconv"
)

// Vector is a 2D vector
// https://en.wikipedia.org/wiki/Euclidean_vector
type Vector struct {
	x, y    float64
	isEmpty bool
}

func NewVector(x, y float64) Vector {
	return Vector{x, y, false}
}

// NewVectorOf2Points create the new vector which through 2 points
// vectorMN: M -> N
func NewVectorOf2Points(m, n Point) Vector {
	return n.Sub(m)
}

func (v Vector) String() string {
	return "(" + strconv.FormatFloat(v.x, 'f', -1, 64) + ", " + strconv.FormatFloat(v.y, 'f', -1, 64) + ")"
}

func (v *Vector) Module() float64 {
	return math.Sqrt(v.x*v.x + v.y*v.y)
}

func (v *Vector) Mul(f float64) Vector {
	return Vector{v.x * f, v.y * f, false}
}
