package graph

// import (
	// "fmt"
	// "errors"
// )

type Vertex struct {
	Id string
	Edges map[string]*Edge
	KeyScore int
	Winner *Edge
	Position int
}

type Edge struct {
	Vertex1Id string
	Vertex2Id string
	Cost int
}

type Graph struct {
	vertices map[string]*Vertex
}

func NewGraph() *Graph {
	g := Graph{ vertices: make(map[string]*Vertex) }

	return &g
}

func (g *Graph) Vertices() map[string]*Vertex {
	return g.vertices
}

func (g *Graph) AddVertex(vertexId string) *Vertex {
	if _, ok := g.vertices[vertexId]; !ok {
		g.vertices[vertexId] = &Vertex{Id: vertexId, Edges: make(map[string]*Edge)}
	}

	return g.vertices[vertexId]
}

func (g *Graph) AddEdge(vertex1Id, vertex2Id string, cost int) (bool, error) {
	vertex1, v1ok := g.vertices[vertex1Id]
	// if vertex does not exist, create it
	if !v1ok {
		vertex1 = g.AddVertex(vertex1Id)
	}

	vertex2, v2ok := g.vertices[vertex2Id]
	if !v2ok {
		vertex2 = g.AddVertex(vertex2Id)
	}

	newEdge := &Edge{Vertex1Id: vertex1Id, Vertex2Id: vertex2Id, Cost: cost}

	if _, ok := vertex1.Edges[vertex2Id]; !ok {
		vertex1.Edges[vertex2Id] = newEdge
	}

	if _, ok := vertex2.Edges[vertex1Id]; !ok {
		vertex2.Edges[vertex1Id] = newEdge
	}

	return true, nil
}