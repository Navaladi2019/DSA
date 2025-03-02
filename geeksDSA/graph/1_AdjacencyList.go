package graph

import "fmt"

func AddEdgeList(arr [][]int, u int, v int) {

	var ismapped bool = false
	for i := 0; i < len(arr[u]); i++ {

		if (arr[u][i]) == v {
			ismapped = true
			break
		}
	}
	if ismapped {
		return
	}

	arr[u] = append(arr[u], v)
	arr[v] = append(arr[v], u)
}

func CreateGraphAdjacencyList() {

	arr := make([][]int, 4)
	AddEdgeList(arr, 0, 1)
	AddEdgeList(arr, 0, 1)
	AddEdgeList(arr, 0, 2)
	AddEdgeList(arr, 1, 2)
	AddEdgeList(arr, 1, 3)
	PrintGraphAdjacencyList(arr)
}

func PrintGraphAdjacencyList(arr [][]int) {

	for i, v := range arr {
		fmt.Println("[", i, v, "]")
	}
}
