package matrix

import "testing"

type BoundaryTraversalTc struct {
	mat  [][]int
	want []int
}

var BoundaryTraversalTcs = func() []SnakePatternTc {

	return []SnakePatternTc{
		{[][]int{
			{1, 2, 3, 4},
			{5, 6, 7, 8},
			{9, 10, 11, 12},
			{13, 14, 15, 16}},
			[]int{1, 2, 3, 4, 8, 12, 16, 15, 14, 13, 9, 5}},
		{[][]int{{1}, {5}, {9}, {13}}, []int{1, 5, 9, 13}},
		{[][]int{{1, 2, 3, 4, 5}}, []int{1, 2, 3, 4, 5}},
	}
}

func Test_BoundaryTraversalTcs(t *testing.T) {

	for _, tc := range BoundaryTraversalTcs() {

		got := BoundaryTraversal_1(tc.mat)

		for i := 0; i < len(got); i++ {
			if tc.want[i] != got[i] {
				t.Error("has boundary traversal Error")
			}
		}
	}
}
