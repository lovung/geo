package geo

type Segment struct {
	vex1, vex2 Point
}

func NewSegment(vex1, vex2 Point) Segment {
	return Segment{vex1, vex2}
}
