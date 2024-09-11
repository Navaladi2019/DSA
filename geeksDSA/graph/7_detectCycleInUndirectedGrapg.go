package graph

import (
	"fmt"

	queue "github.com/Navaladi2019/GoRefresher/geeksDSA/Queue"
)

func DetectCycleInUndirectedGraphBFS(arr [][]int) {
	visited := make([]bool, len(arr))
	parent := make([]int, len(arr))

	for I, _ := range parent {
		parent[I] = -1
	}

	q := queue.ArrQueue[int]{}
	q.Init(100)
	q.Enqueue(0)
	visited[0] = true
	hasllop := false
	for !q.IsEmpty() && hasllop == false {
		u, _ := q.Dequeue()
		for _, val := range arr[u] {
			if visited[val] == false {
				parent[val] = u
				q.Enqueue(val)
				visited[val] = true
			} else {
				if parent[u] != val {
					hasllop = true
					fmt.Println("has loop found at", u, val, "has already parent of", parent[u])
					break
				}
			}

		}
	}

	if hasllop == false {
		fmt.Println("No loop found")
	}

}
func DFSDetectCycleRec(arr [][]int, u int, visited []bool, parent int) bool {

	visited[u] = true
	for _, v := range arr[u] {
		if visited[v] == false {
			if DFSDetectCycleRec(arr, v, visited, u) == true {
				return true
			}
		} else if v != parent {
			fmt.Println("has loop found at", u, v, "has already parent of", parent)
			return true
		}
	}

	return false
}

func DFS_DetectCycle(arr [][]int, u int) {

	visited := make([]bool, len(arr))
	DFSDetectCycleRec(arr, 0, visited, -1)
}
