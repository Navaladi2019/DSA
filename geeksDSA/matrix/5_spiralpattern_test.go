package matrix

import "testing"

type SpiralTc struct {
	mat  [][]int
	want []int
}

var SpiralTcs = func() []SpiralTc {
	return []SpiralTc{
		{[][]int{{1, 2, 3, 4},
			{5, 6, 7, 8},
			{9, 10, 11, 12}},
			[]int{1, 2, 3, 4, 8, 12, 11, 10, 9, 5, 6, 7}},
		{[][]int{{1, 2, 3, 4}},
			[]int{1, 2, 3, 4}},
		{[][]int{{1}, {2}, {3}, {4}},
			[]int{1, 2, 3, 4}},

		{[][]int{{1}},
			[]int{1}},
		{[][]int{{1, 2, 3},
			{4, 5, 6},
			{7, 8, 9}},
			[]int{1, 2, 3, 6, 9, 8, 7, 4, 5}},
	}
}

func Test_TestSpiral(t *testing.T) {

	for _, tc := range SpiralTcs() {
		got := FindSpiralPattern(tc.mat)

		for i := 0; i < len(got); i++ {
			if got[i] != tc.want[i] {
				t.Error("has errror in spiral pattern")
			}
		}
	}
}
