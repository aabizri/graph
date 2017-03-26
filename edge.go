package graph

// Edge represents an edge
type Edge interface {
	ID() EdgeID
	Source() VertexID
	Destination() VertexID
	Weight() float64
}

// EdgeID Must be unique
type EdgeID interface {
	String() string
}
