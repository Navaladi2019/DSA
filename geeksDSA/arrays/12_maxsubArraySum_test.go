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

var MaxSumofSubArrayCirculartcs = func() []MaxSumofSubArraytc {

	return []MaxSumofSubArraytc{
		{[]int{5, -2, 3, 4}, 12},
		{[]int{2, 3, -4}, 5},
		{[]int{8, -4, 3, -5, 4}, 12},
		{[]int{-3, 4, 6, -2}, 10},
	}
}

func Test_GetMaxSumofSubArray(t *testing.T) {

	for _, tc := range MaxSumofSubArraytcs() {
		got := getMaxSumOfSubArray_1(tc.ip)

		if got != tc.want {
			t.Error("failed")
		}
	}
}

func Test_GetMaxSumofSubArrayCircular(t *testing.T) {

	for _, tc := range MaxSumofSubArrayCirculartcs() {
		got := getMaxSumOfSubArray_Circular(tc.ip)

		if got != tc.want {
			t.Error("failed")
		}
	}
}
