package graph

import (
	"testing"
)

func Test_FindMinSpanningTreePrimsAlgo(t *testing.T) {

	arr := [][]int{}

	arr = append(arr, []int{0, 5, 8, 0})
	arr = append(arr, []int{5, 0, 10, 15})
	arr = append(arr, []int{8, 10, 0, 20})
	arr = append(arr, []int{0, 15, 20, 0})

	got := FindMinSpannigTreeEfficient(arr)

	if got != 28 {
		t.Error("ahs error ")
	}

	arr = [][]int{}

	arr = append(arr, []int{0, 2, 0, 6, 0})
	arr = append(arr, []int{2, 0, 3, 8, 5})
	arr = append(arr, []int{0, 3, 0, 0, 7})
	arr = append(arr, []int{6, 8, 0, 0, 9})
	arr = append(arr, []int{0, 5, 7, 9, 0})

	got = FindMinSpannigTreeEfficient(arr)

	if got != 16 {
		t.Error("ahs error ")
	}
}
