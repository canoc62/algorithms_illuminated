package prim

import (
	"container/heap"
	"fmt"
	"math"

	vertexHeap "github.com/algorithms_illuminated/greedy_algorithms/mst/heap"
	graph "github.com/algorithms_illuminated/greedy_algorithms/mst/graph"
)

func Prim(g *graph.Graph, firstVertexId string) []*graph.Edge {
	x := make(map[string]struct{})
	x[firstVertexId] = struct{}{}

	t := []*graph.Edge{}

	h := &vertexHeap.VertexHeap{}
	heap.Init(h)
 
	vertices := g.Vertices()
	firstVertex, ok := vertices[firstVertexId]
	if !ok {
		panic(fmt.Sprintf("first vertex with id: %s does not exist in graph\n", firstVertexId))
	}

	for k, v := range vertices {
		if k != firstVertexId {
			e, ok := firstVertex.Edges[k]
			if ok {
				v.KeyScore = e.Cost
				v.Winner = e
			} else {
				v.KeyScore = int(math.MaxInt32)
				v.Winner = nil
			}
			heap.Push(h, v)
		}
	}

	for len(*h) > 0 {
		w := heap.Pop(h).(*graph.Vertex) // type assertion
		x[w.Id] = struct{}{}
		t = append(t, w.Winner)

		for edgeEndId, e := range w.Edges {
			_, ok := x[edgeEndId]
			if !ok {
				// remove vertex with edgeEndId From h
				oldEndV, ok := g.Vertices()[edgeEndId]
				if !ok {
					panic(fmt.Sprintf("vertex %s doesn't exist in graph\n", edgeEndId))
				}
				if e.Cost < oldEndV.KeyScore {
					oldEndV = heap.Remove(h, oldEndV.Position).(*graph.Vertex)
					// set keyScore of vertex to cost of e
					oldEndV.KeyScore = e.Cost
					// set winner of vertex to e
					oldEndV.Winner = e
					// insert vertex back into h
					heap.Push(h, oldEndV)
				}
			}
		}
	}

	return t
}