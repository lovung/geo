package geo

import (
	"math"
	"strconv"
)

// Vector is a 2D vector
// https://en.wikipedia.org/wiki/Euclidean_vector
type Vector struct {
	x, y float64
}

// ZeroVector is the vector (0;0)
var (
	ZeroVector = Vector{0, 0}
)

// NewVector with x and y value
func NewVector(x, y float64) Vector {
	return Vector{x, y}
}

// NewVectorOf2Points create the new vector which through 2 points
// vectorMN: M -> N
func NewVectorOf2Points(m, n Point) Vector {
	return n.Sub(m)
}

func (v Vector) String() string {
	return "(" + strconv.FormatFloat(v.x, 'f', -1, 64) +
		", " + strconv.FormatFloat(v.y, 'f', -1, 64) + ")"
}

// Module |AB| = sqrt(x^2 + y^2)
func (v *Vector) Module() float64 {
	return math.Sqrt(v.x*v.x + v.y*v.y)
}

// Mul with a number to scale the vector
func (v *Vector) Mul(f float64) Vector {
	if f == 0 {
		panic("Division by zero")
	}
	return Vector{v.x * f, v.y * f}
}

// Div with a number to scale the vector
func (v *Vector) Div(f float64) Vector {
	return Vector{v.x / f, v.y / f}
}

// Add v + v2 (AB + BC = AC)
func (v *Vector) Add(v2 Vector) Vector {
	return Vector{v.x + v2.x, v.y + v2.y}
}

// Sub v - v2 = v + (-v2)
func (v *Vector) Sub(v2 Vector) Vector {
	return Vector{v.x - v2.x, v.y - v2.y}
}

// Clone create the same vector
func (v *Vector) Clone() Vector {
	return Vector{v.x, v.y}
}

// Alpha returns the angle between v and Ox ray
// 		Alpha = atan2(y, x)
// Alpha = 0 -> v = Ox
// Alpha = pi/2 -> v = Oy
// Alpha = pi -> v = -Ox
// Alpha = -pi/2 -> v = -Oy
// -pi < Alpha <= pi
func (v *Vector) Alpha() Angle {
	return Angle(math.Atan2(v.y, v.x))
}

// Theta returns the angle between 2 vectors
// positive direction is counter-clockwise
// v.Theta(v2) = acos(v.Dot(v2) / (v.Module() * v2.Module()))
// v.Theta(v2) = - v2.Theta(v)
// -pi < Theta <= pi
func (v *Vector) Theta(v2 Vector) Angle {
	// calculate unit-vector of vector
	theta := v.Alpha() - v2.Alpha()
	if theta <= -math.Pi {
		theta += math.Pi * 2
	}
	if theta > math.Pi {
		theta -= math.Pi * 2
	}
	return Angle(theta)
}
