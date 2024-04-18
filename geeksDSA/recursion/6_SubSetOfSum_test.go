package recursion

import "testing"

type SubSetSum struct {
	ip   []int
	sum  int
	want int
}

var SubSetSumTcs = func() []SubSetSum {
	return []SubSetSum{
		{[]int{10, 5, 2, 3, 6}, 8, 2},
		{[]int{1, 2, 3}, 4, 1},
		{[]int{10, 10, 15}, 0, 1},
	}
}

func Test_SubSetSum(t *testing.T) {

	for _, tc := range SubSetSumTcs() {

		sum := 0

		for i := 0; i < len(tc.ip); i++ {
			sum += tc.ip[i]
		}
		got := FindSubSetProblemOptimized(tc.ip, 0, tc.sum)

		if tc.want != got {
			t.Error("Has Error")
		}
	}

}
