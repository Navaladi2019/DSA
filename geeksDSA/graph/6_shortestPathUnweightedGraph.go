package graph

import (
	"fmt"

	queue "github.com/Navaladi2019/GoRefresher/geeksDSA/Queue"
)

func FindShortestPathUnweightedGraph(arr [][]int, u int) {

	distance := make([]int, len(arr))
	visited := make([]bool, len(arr))

	q := queue.ArrQueue[int]{}
	q.Init(100)
	q.Enqueue(u)
	visited[u] = true

	for !q.IsEmpty() {
		u, _ = q.Dequeue()

		for _, v := range arr[u] {
			if visited[v] == false {
				distance[v] = distance[u] + 1
				visited[v] = true
				q.Enqueue(v)
			}
		}
	}

	fmt.Println(distance)
}

func FindShortestPathINweightedGraph_1(arr [][]int, u int) {

	visited := make([]bool, len(arr))
	distance := make([]int, len(arr))

	q := queue.ArrQueue[int]{}
	q.Enqueue(u)

	distance[u] = 0
	visited[u] = true

	for !q.IsEmpty() {
		u, _ := q.Dequeue()
		for i := 0; i < len(arr[u]); i++ {
			if visited[arr[u][i]] == false {
				distance[arr[u][i]] = distance[u] + 1
				visited[arr[u][i]] = true
				q.Enqueue(arr[u][i])
			}
		}

	}
}
