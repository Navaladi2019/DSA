package arrays

import "testing"

type MoveSeroToLastTc struct {
	arr1 []int
	arr2 []int
}

var MoveSeroToLastTcs = func() []MoveSeroToLastTc {
	return []MoveSeroToLastTc{
		{[]int{8, 5, 0, 10, 20}, []int{8, 5, 10, 20, 0}},
		{[]int{8, 5, 0, 10, 0, 0, 20, 30, 0}, []int{8, 5, 10, 20, 30, 0, 0, 0, 0}},
		{[]int{0, 0, 0, 10, 0}, []int{10, 0, 0, 0, 0}},
		{[]int{10, 20, 30}, []int{10, 20, 30}},
	}
}

func Test_MoveSeroToLastTcs(t *testing.T) {

	for _, tc := range MoveSeroToLastTcs() {

		MoveZeroToLast_1(tc.arr1)

		for i := 0; i < len(tc.arr2); i++ {
			if tc.arr1[i] != tc.arr2[i] {
				t.Error("has error in moving zero to last")
				break
			}
		}
	}
}
