package graph

import "fmt"

func DFSRec(arr [][]int, u int, visited []bool) {
	fmt.Println(u)
	visited[u] = true
	for _, val := range arr[u] {
		if visited[val] == false {
			DFSRec(arr, val, visited)
		}
	}
}

func DFS(arr [][]int, u int) {

	visited := make([]bool, len(arr))
	DFSRec(arr, 0, visited)
}
