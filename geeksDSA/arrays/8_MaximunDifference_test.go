package arrays

import "testing"

type MaxDiffArrayTc struct {
	ip   []int
	want int
}

var MaxDiffArrayTcs = func() []MaxDiffArrayTc {
	return []MaxDiffArrayTc{
		{[]int{2, 3, 10, 6, 4, 8, 1}, 8},
		{[]int{7, 9, 5, 6, 3, 2}, 2},
		{[]int{10, 20, 30}, 20},
		{[]int{30, 10, 8, 2}, -2},
	}
}

func Test_MaxDiffArrayTcs(t *testing.T) {

	for _, tc := range MaxDiffArrayTcs() {

		got := findMaximunDifferenceArray(tc.ip)

		if got != tc.want {
			t.Error("has error")
		}
	}
}
