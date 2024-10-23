package graph

import (
	"fmt"

	"github.com/Navaladi2019/GoRefresher/geeksDSA/stack"
)

// to solve strongly connected components
// if you can reach from one component to any other component then its called as strongly connected component

// in case of undirected graph then all connected components are styronngly connected components

// order the dfs in decreasing order of finish times

func KosarujusAlgoFroStronglyConnectedComponents(arr Graph) {

	sta := stack.ArrStack[int]{}
	visited := make([]bool, len(arr))

	for i := 0; i < len(arr); i++ {
		getVerticesInDecreasingOrder(arr, &sta, visited, i)
	}

	// transpose a matrix to change the direction of graph

	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			arr[i][j], arr[j][i] = arr[j][i], arr[i][j]
		}
	}

	isPrinted := make([]bool, len(arr))

	for !sta.IsEmpty() {

		val, _ := sta.Pop()

		if isPrinted[val] == true {
			continue
		}

		PrintConnectedVertices(arr, isPrinted, val)
		fmt.Println()
	}

}

func PrintConnectedVertices(arr Graph, isPrinted []bool, i int) {

	if isPrinted[i] == true {
		return
	}
	fmt.Print(i, ",")
	isPrinted[i] = true
	for j := 0; j < len(arr); j++ {
		if arr[i][j] != 0 {
			PrintConnectedVertices(arr, isPrinted, j)
		}
	}
}

func getVerticesInDecreasingOrder(arr Graph, sta *stack.ArrStack[int], visited []bool, i int) {

	if visited[i] == true {
		return
	}
	visited[i] = true
	for j := 0; j < len(arr); j++ {
		if arr[i][j] != 0 {
			getVerticesInDecreasingOrder(arr, sta, visited, j)
		}
	}
	sta.Push(i)
}
