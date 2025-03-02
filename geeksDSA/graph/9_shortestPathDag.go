package graph

import (
	"fmt"
	"math"

	queue "github.com/Navaladi2019/GoRefresher/geeksDSA/Queue"
)

// it will now work for negative weights
// shortest path will change when we add weight to all edges

func GetShortestPathDag(arr [][]int, s int) {

	q := queue.ArrQueue[int]{}
	q.Init(100)

	distance := make([]int, len(arr))

	for i := range distance {
		distance[i] = math.MaxInt
	}

	distance[s] = 0
	toposort := SortItemInTopologicalOrderSortedWeightedGraph(arr)

	for _, v := range toposort {
		if distance[v] < math.MaxInt {
			for i := 0; i < len(arr[v]); i++ {
				if arr[v][i] > 0 {
					distance[i] = min(distance[i], distance[v]+arr[v][i])
				}
			}
		}
	}

	fmt.Println(distance)
}

func SortItemInTopologicalOrderSortedWeightedGraph(arr [][]int) (res []int) {

	edges := make([]int, len(arr))
	for i := range arr {
		for j := range arr[i] {
			if arr[i][j] > 0 {
				edges[j]++
			}
		}
	}

	q := queue.ArrQueue[int]{}
	q.Init(20)

	for i := range edges {
		if edges[i] == 0 {
			q.Enqueue(i)
		}
	}

	for !q.IsEmpty() {
		val, _ := q.Dequeue()
		res = append(res, val)
		for j := range arr[val] {
			if arr[val][j] > 0 && j != val {
				edges[j]--
				if edges[j] == 0 {
					q.Enqueue(j)
				}
			}
		}
	}
	return res
}
