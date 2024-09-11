package graph

import (
	"fmt"

	queue "github.com/Navaladi2019/GoRefresher/geeksDSA/Queue"
)

func BFS_Graph(arr [][]int, visited []bool, u int) {
	q := queue.ArrQueue[int]{}
	visited[u] = true
	q.Init(100)
	q.Enqueue(u)
	for !q.IsEmpty() {
		v, _ := q.Dequeue()
		fmt.Println(v)
		for _, val := range arr[v] {
			if visited[val] == false {
				visited[val] = true
				q.Enqueue(val)
			}
		}
	}
}

func BFS_Disconnected(arr [][]int) {
	visited := make([]bool, len(arr))
	for i, _ := range visited {
		if visited[i] == false {
			BFS_Graph(arr, visited, i)
		}
	}
}
