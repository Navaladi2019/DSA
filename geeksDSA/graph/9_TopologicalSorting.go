package graph

import (
	"fmt"
	"slices"

	queue "github.com/Navaladi2019/GoRefresher/geeksDSA/Queue"
)

// this uses Khans BFS based algorithm
// Topological sorting will only work if there is no cycle
func SortItemInTopologicalOrderInDirectedGraph(arr [][]int) []int {

	toposort := make([]int, 0, len(arr))
	edges := make([]int, len(arr))
	for _, u := range arr {
		for _, v := range u {
			edges[v]++
		}
	}
	zeroEdge := FilterSlicesGetIndex(edges, func(e int) bool { return e == 0 })
	q := queue.ArrQueue[int]{}
	q.Init(100)
	for _, val := range zeroEdge {
		q.Enqueue(val)
	}

	for !q.IsEmpty() {
		v, _ := q.Dequeue()
		fmt.Println(v)
		toposort = append(toposort, v)
		for _, val := range arr[v] {
			edges[val] = edges[val] - 1
			if edges[val] == 0 {
				q.Enqueue(val)
			}
		}
	}

	return toposort
}

func FilterSlicesGetIndex[T any](arr []T, filter func(ele T) bool) []int {
	res := make([]int, 0, len(arr))

	for i, v := range arr {
		if filter(v) {
			res = append(res, i)
		}
	}
	res = slices.Clip(res)
	return res

}
