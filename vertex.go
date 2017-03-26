package graph

// Vertex is a vertex
type Vertex interface {
	ID() VertexID
}

// VertexID must be comparable
type VertexID interface {
	String() string
}
