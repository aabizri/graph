/*
Package graph is a graph library made in go (golang) and in the go spirit.
As such graph DOES NOT bundle any method of storing a graph, and no function to manipulate one either, it is thus your responsibility to have one !
A graph manipulation library is available at github.com/aabizri/graphutils
*/
package graph

// Graph is an elementary graph type
type Graph interface {
	// AddEdge adds an edge to the Graph
	AddEdge(u, v VertexID) error

	// DelEdge deletes an Edge from the Graph
	DelEdge(e EdgeID) error

	// Edges retrieves the list of edges between two Vertices
	Edges(u, v VertexID) ([]Edge, error)

	// AddVertex adds a Vertex to the Graph
	AddVertex(u Vertex) error

	// DelVertex deletes a Vertex by its ID, also should delete all edges including this vertex
	DelVertex(u VertexID) error

	// Vertex retrieves a Vertex given its ID
	// Returns VertexNonExistentErr if no such vertex exist
	Vertex(u VertexID) (Vertex, error)

	// Vertices retrieves the list of all vertices
	Vertices() ([]VertexID, error)
}

// GraphResetter implements a Reset method which reinitialises a graph
type GraphResetter interface {
	Graph

	// Reset reinitialises a graph
	Reset() error
}

// GraphAllEdgesRetriever implements a AllEdges method which returns the list of all edges.
type GraphAllEdgesRetriever interface {
	Graph

	// Edges retrieves a list of all the edges
	AllEdges() ([]Edge, error)
}

// A GraphCloner implements a Clone method which clones the graph and returns an exact copy
type GraphCloner interface {
	Graph

	// Clone a graph
	Clone() (Graph, error)
}
