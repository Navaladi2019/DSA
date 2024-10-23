package graph

import (
	"testing"
)

func Test_bellmanFord_shortestPath(t *testing.T) {

	arr := make([][]int, 4)
	arr[0] = []int{0, 1, 4, 0}
	arr[1] = []int{0, 0, -3, 2}
	arr[2] = []int{0, 0, 0, 3}
	arr[3] = []int{0, 0, 0, 0}

	bellmanFord_shortestPath(arr, 0)
}
