package searching

import "testing"

type PairsEqualToSumtc struct {
	arr  []int
	n    int
	want bool
}

var PairsEqualToSumtcs = func() []PairsEqualToSumtc {
	return []PairsEqualToSumtc{
		{[]int{5, 8, 10, 20, 40, 60}, 60, true},
		{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16}, 3, true},
		{[]int{1, 2, 3, 4, 5, 6, 7, 8, 94}, 98, true},
		{[]int{100, 200, 500, 1000, 2000}, 300, true},
		{[]int{100, 200, 500, 1000, 2000}, 1500, true},
	}
}

func Test_PairsEqualToSumtcs(t *testing.T) {
	for _, tc := range PairsEqualToSumtcs() {

		got := FindPairsEqualToSum(tc.arr, tc.n)

		if got != tc.want {
			t.Error("has error ")
		}
	}
}
