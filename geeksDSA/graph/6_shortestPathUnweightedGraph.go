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
