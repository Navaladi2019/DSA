package matrix

import "testing"

type SnakePatternTc struct {
	mat  [][]int
	want []int
}

var SnakePatternTcs = func() []SnakePatternTc {

	return []SnakePatternTc{
		{[][]int{
			{1, 2, 3, 4},
			{5, 6, 7, 8},
			{9, 10, 11, 12},
			{13, 14, 15, 16}},
			[]int{1, 2, 3, 4, 8, 7, 6, 5, 9, 10, 11, 12, 16, 15, 14, 13}},
	}
}

func Test_SnakePattern(t *testing.T) {

	for _, tc := range SnakePatternTcs() {

		got := GetArrayInSnakePattern(tc.mat)

		for i := 0; i < len(tc.want); i++ {

			if got[i] != tc.want[i] {
				t.Error("has error in snake pattern")
			}
		}
	}
}
