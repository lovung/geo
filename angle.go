package geo

import "math"

type Angle float64

// Angle units.
const (
	Radian Angle = 1
	Degree       = (math.Pi / 180) * Radian
)
