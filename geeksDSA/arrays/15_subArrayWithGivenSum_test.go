package arrays

import "testing"

type isSubArrayWithGivenSumTc struct {
	ip   []int
	sum  int
	want bool
}

var isSubArrayWithGivenSumTcs = func() []isSubArrayWithGivenSumTc {
	return []isSubArrayWithGivenSumTc{
		{[]int{1, 4, 20, 3, 10, 5}, 33, true},
		{[]int{1, 4, 0, 0, 3, 10, 5}, 7, true},
		{[]int{2, 4}, 3, false},
		{[]int{3, 2, 0, 4, 7}, 6, true},
	}
}

func Test_isSubArrayWithGivenSum(t *testing.T) {
	for _, tc := range isSubArrayWithGivenSumTcs() {

		got := isSubArrayWithGivenSum(tc.ip, tc.sum)

		if got != tc.want {
			t.Error("has error")
		}
	}
}
