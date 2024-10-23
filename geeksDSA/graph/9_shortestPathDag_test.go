package graph

import (
	"testing"
)

func Test_GetShortestPathDag(t *testing.T) {

	arr := make([][]int, 9)
	arr[0] = []int{0, 2, 0, 0, 0, 0, 0, 0, 5}
	arr[1] = []int{0, 0, 3, 0, 0, 0, 0, 0, 0}
	arr[2] = []int{0, 0, 0, 6, 0, 0, 0, 3, 0}
	arr[3] = []int{0, 0, 0, 0, 0, 0, 2, 0, 0}
	arr[4] = []int{0, 2, 2, 0, 0, 4, 0, 0, 0}
	arr[5] = []int{0, 0, 0, 1, 0, 0, 0, 0, 0}
	arr[6] = []int{0, 0, 0, 0, 0, 0, 0, 0, 0}
	arr[7] = []int{0, 0, 0, 0, 0, 0, 0, 0, 0}
	arr[8] = []int{0, 0, 0, 0, 1, 0, 0, 0, 0}
	GetShortestPathDag(arr, 0)
	GetShortestPathDag(arr, 1)
	GetShortestPathDag(arr, 2)
	GetShortestPathDag(arr, 3)
	GetShortestPathDag(arr, 4)
	GetShortestPathDag(arr, 5)
	GetShortestPathDag(arr, 8)
}
