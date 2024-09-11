package graph

import (
	queue "github.com/Navaladi2019/GoRefresher/geeksDSA/Queue"
)

// It works by the concept that counter will not not be equal to vertices if there is a loop
// because in cycle we cannot find zero indegree vertices to enqueue into
func DetectLoopUsingKhans(arr [][]int) bool {

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
	counter := 0
	for !q.IsEmpty() {
		counter++
		v, _ := q.Dequeue()
		for _, val := range arr[v] {
			edges[val] = edges[val] - 1
			if edges[val] == 0 {
				q.Enqueue(val)
			}
		}
	}

	return counter != len(arr)
}
