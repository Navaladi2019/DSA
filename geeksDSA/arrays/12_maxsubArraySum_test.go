package arrays

import (
	"testing"
)

//[1,2,3,4]

func Test_subArrays(t *testing.T) {

	got := getSubArrays([]int{1, 2, 3}, 0, [][]int{})
	got = getSubArraysForLoop([]int{1, 2, 3})
	t.Log(got)
}

type MaxSumofSubArraytc struct {
	ip   []int
	want int
}

var MaxSumofSubArraytcs = func() []MaxSumofSubArraytc {

	return []MaxSumofSubArraytc{
		{[]int{2, 3, -8, 7, -1, 2, 3}, 11},
		{[]int{5, 8, 3}, 16},
		{[]int{-6, -1, -8}, -1},
	}
}

func Test_GetMaxSumofSubArray(t *testing.T) {

	for _, tc := range MaxSumofSubArraytcs() {
		got := getMaxSumOfSubArray(tc.ip)

		if got != tc.want {
			t.Error("failed")
		}
	}
}
