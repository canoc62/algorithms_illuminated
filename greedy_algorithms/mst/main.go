package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"bufio"
	"strings"
	"strconv"

	graph "github.com/algorithms_illuminated/greedy_algorithms/mst/graph"
	prim "github.com/algorithms_illuminated/greedy_algorithms/mst/prim"
	// graph "github.com/algorithms_illuminated/greedy_algorithms/mst/graph"
)

func main() {
	filePtr := flag.String("f", "tests/test1.txt", "filename")	
	flag.Parse()

	file, err := os.Open(*filePtr)
	if err != nil {
		log.Fatalln(err)
		os.Exit(1)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	g := graph.NewGraph()

	scanner.Scan()

	for scanner.Scan() {
    vals := strings.Fields(scanner.Text())
    vertex1Id := vals[0]
		vertex2Id := vals[1]
		cost, strToIErr := strconv.Atoi(vals[2])
		if strToIErr != nil {
			panic(strToIErr)
		}
		_, addEdgeErr := g.AddEdge(vertex1Id, vertex2Id, cost)
		if err != nil {
			panic(addEdgeErr)
		}
	}
	
	// grab random vertex
	firstVertex := ""
	for _, v := range g.Vertices() {
		firstVertex = v.Id
		break
	}
	spanningTreeEdges := prim.Prim(g, firstVertex)

	sumPrimCosts := 0
	for _, e := range spanningTreeEdges {
		sumPrimCosts += e.Cost
	}

	fmt.Printf("Prim sum of edge costs in graph: %d\n", sumPrimCosts)
}