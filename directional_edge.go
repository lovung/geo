package geo

import "math"

// DirectionalEdge has 2 points and has direction
// from start to end
type DirectionalEdge struct {
	Start, End Point
}

// Len is the distance from Start point to End point
func (de *DirectionalEdge) Len() float64 {
	dX := de.End.X - de.Start.X
	dY := de.End.Y - de.Start.Y
	return math.Sqrt(dX*dX + dY*dY)
}

func (de DirectionalEdge) String() string {
	return de.Start.String() + " -> " + de.End.String()
}
