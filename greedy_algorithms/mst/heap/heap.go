package heap

import (
  graph "github.com/algorithms_illuminated/greedy_algorithms/mst/graph"
)

type VertexHeap []*graph.Vertex

func (vh VertexHeap) Len() int {
  return len(vh)
}

func (vh VertexHeap) Less(i, j int) bool {
  return vh[i].KeyScore < vh[j].KeyScore
}

func (vh VertexHeap) Swap(i, j int) {
  vh[i], vh[j] = vh[j], vh[i]
  vh[i].Position = i
  vh[j].Position = j
}

func (vh *VertexHeap) Push(x interface{}) {
	// Push and Pop use pointer receivers because they modify the slice's length,
  // not just its contents.
  position := len(*vh)
  v := x.(*graph.Vertex)
  v.Position = position
  *vh = append(*vh, v)
}

func (vh *VertexHeap) Pop() interface{} {
	old := *vh
	n := len(old)
  v := old[n-1]
  old[n-1] = nil // avoid memory leak
  v.Position = -1 // for safety according to golang docs example
	*vh = old[0 : n-1]
	return v
}