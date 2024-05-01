package hashing

import "testing"

type LongestSubArrayWithGivenSumTc struct {
	arr  []int
	sum  int
	want int
}

var LongestSubArrayWithGivenSumTcs = func() []LongestSubArrayWithGivenSumTc {
	return []LongestSubArrayWithGivenSumTc{
		{[]int{5, 8, -4, -4, 9, -2, 2}, 0, 3},
		{[]int{3, 1, 0, 1, 8, 2, 3, 6}, 5, 4},
		{[]int{8, 3, 7}, 15, 0},
		{[]int{8, 3, 1, 5, -6, 6, 2, 2}, 4, 4},
	}
}

func Test_LongestSubArrayWithGivenSumTcs(t *testing.T) {

	for _, tc := range LongestSubArrayWithGivenSumTcs() {

		got := LongestSubArrayWithGivenSum(tc.arr, tc.sum)

		if got != tc.want {
			t.Error("Has error on longest find problem")
		}
	}
}
