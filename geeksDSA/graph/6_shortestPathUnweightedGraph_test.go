package graph

import "testing"

func Test_FindShortestPathUnweightedGraph(t *testing.T) {

	arr := make([][]int, 4)
	arr[0] = []int{1, 2}
	arr[1] = []int{0, 2, 3}
	arr[2] = []int{0, 1, 3}
	arr[3] = []int{1, 2}
	FindShortestPathUnweightedGraph(arr, 0)

	arr = make([][]int, 6)
	arr[0] = []int{1, 2, 4}
	arr[1] = []int{0, 3}
	arr[2] = []int{0, 3, 4}
	arr[3] = []int{1, 2, 5}
	arr[4] = []int{0, 2, 5}
	arr[5] = []int{4, 3}
	FindShortestPathUnweightedGraph(arr, 0)

}
