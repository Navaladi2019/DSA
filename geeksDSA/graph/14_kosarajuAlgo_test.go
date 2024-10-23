package graph

import "testing"

func Test_KosaRajus(t *testing.T) {

	var graph = make(Graph, 6)
	graph[0] = []int{0, 1, 0, 0, 0, 0}
	graph[1] = []int{0, 0, 1, 0, 0, 0}
	graph[2] = []int{0, 0, 0, 1, 0, 0}
	graph[3] = []int{1, 0, 0, 0, 1, 0}
	graph[4] = []int{0, 0, 0, 0, 0, 1}
	graph[5] = []int{0, 0, 0, 0, 1, 0}
	KosarujusAlgoFroStronglyConnectedComponents(graph)
}
