package graph

import (
	"fmt"
	"math"
)

func bellmanFord_shortestPath(graph Graph, s int) {
	dist := make([]int, len(graph))
	for i := 0; i < len(graph); i++ {
		dist[i] = math.MaxInt
	}
	dist[s] = 0
	for i := 0; i < len(graph)-1; i++ {
		for vi, _ := range graph {
			for wi, w := range graph[vi] {
				if w != 0 && dist[vi] != math.MaxInt {
					dist[wi] = min(dist[wi], dist[vi]+w)
				}
			}
		}
	}

	fmt.Println(dist)
}
