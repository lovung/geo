package geo

// Polygon has 4 kinds
// - open (excluding its boundary),
// - boundary only (excluding interior),
// - closed (including both boundary and interior),
// - and self-intersecting.
// Empty Polygon has no vertices
// For simple, we are working with closed polygon first
// (includes convex and concave)
type Polygon struct {
	vertices []Point
}

// NumOfVertices ...
func (p *Polygon) NumOfVertices() int {
	return len(p.vertices)
}

// DirectionalEdges returns all Directional Edges of the polygon
// by vertices order
func (p *Polygon) DirectionalEdges() []DirectionalEdge {
	n := p.NumOfVertices()
	if n < 2 {
		return nil
	}
	edges := make([]DirectionalEdge, n)
	for i := 1; i < n; i++ {
		edges[i-1] = DirectionalEdge{
			Start: p.vertices[i-1],
			End:   p.vertices[i],
		}
	}
	edges[n-1] = DirectionalEdge{
		Start: p.vertices[n-1],
		End:   p.vertices[0],
	}
	return edges
}
