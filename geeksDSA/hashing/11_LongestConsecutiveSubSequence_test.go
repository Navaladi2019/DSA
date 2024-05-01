package hashing

import "testing"

type LongestConseTc struct {
	arr  []int
	want int
}

var LongestConseTcs = func() []LongestConseTc {
	return []LongestConseTc{
		{[]int{1, 9, 3, 4, 2, 20}, 4},
		{[]int{8, 20, 7, 30}, 2},
		{[]int{20, 30, 40}, 1},
	}
}

func Test_LongestConsecutive(t *testing.T) {

	for _, tc := range LongestConseTcs() {

		got := LongestSubSequence(tc.arr)

		if tc.want != got {
			t.Error("has errror")
		}
	}
}
