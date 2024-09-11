package graph

import (
	"fmt"
	"testing"
)

func Test_SortItemInTopologicalOrderInDirectedGraph(t *testing.T) {

	arr := make([][]int, 3)
	arr[0] = []int{1, 2}
	SortItemInTopologicalOrderInDirectedGraph(arr)

	fmt.Println("***")
	arr = make([][]int, 6)
	arr[0] = []int{1, 2}
	arr[2] = []int{3}
	arr[1] = []int{3}
	arr[3] = []int{4, 5}

	SortItemInTopologicalOrderInDirectedGraph(arr)

	fmt.Println("***")

	arr = make([][]int, 5)
	arr[0] = []int{2, 3}
	arr[1] = []int{3, 4}
	SortItemInTopologicalOrderInDirectedGraph(arr)
}
