package hashing

import "testing"

type COuntDistinctWindowTc struct {
	arr    []int
	window int
	want   []int
}

var COuntDistinctWindowTcs = func() []COuntDistinctWindowTc {
	return []COuntDistinctWindowTc{
		{[]int{10, 20, 20, 10, 30, 40, 10}, 4, []int{2, 3, 4, 3}},
		{[]int{10, 10, 10, 10}, 3, []int{1, 1}},
		{[]int{10, 20, 30, 40}, 3, []int{3, 3}},
	}
}

func Test_COuntDistinctWindowTc(t *testing.T) {
	for _, tc := range COuntDistinctWindowTcs() {
		got := CountDistinctElementINEveryWindow(tc.arr, tc.window)

		for i := 0; i < len(tc.want); i++ {
			if tc.want[i] != got[i] {
				t.Error("has error")
				break
			}
		}
	}
}
