package graph

import (
	"testing"
)

func Test_DetectCycleInUndirectedGraphBFS(t *testing.T) {

	test := GetTesrtCasesForDetectCycle()

	for _, tt := range test {
		t.Run("", func(t *testing.T) {
			DFS_DetectCycle(tt, 0)
		})
	}
}

func GetTesrtCasesForDetectCycle() [][][]int {

	parent := [][][]int{}
	arr1 := make([][]int, 4)
	arr1[0] = []int{1}
	arr1[1] = []int{0, 3, 2}
	arr1[2] = []int{1, 3}
	arr1[3] = []int{1, 2}
	parent = append(parent, arr1)

	arr1 = make([][]int, 5)
	arr1[0] = []int{1}
	arr1[1] = []int{0, 4, 2}
	arr1[2] = []int{1, 3}
	arr1[3] = []int{2}
	arr1[4] = []int{1}
	parent = append(parent, arr1)

	arr1 = make([][]int, 4)
	arr1[0] = []int{1, 3}
	arr1[1] = []int{0, 3, 2}
	arr1[2] = []int{1, 3}
	arr1[3] = []int{0, 1, 2}
	parent = append(parent, arr1)

	return parent
}
