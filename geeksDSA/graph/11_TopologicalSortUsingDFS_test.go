package graph

import (
	"fmt"
	"testing"
)

func Test_TopologicalSOrtDFSREC(t *testing.T) {
	arr := make([][]int, 3)
	arr[0] = []int{1, 2}
	TopologicalSOrtDFS(arr)

	fmt.Println("***")
	arr = make([][]int, 6)
	arr[0] = []int{1, 2}
	arr[2] = []int{3}
	arr[1] = []int{3}
	arr[3] = []int{4, 5}

	TopologicalSOrtDFS(arr)

	fmt.Println("***")

	arr = make([][]int, 5)
	arr[0] = []int{2, 3}
	arr[1] = []int{3, 4}
	TopologicalSOrtDFS(arr)
}
