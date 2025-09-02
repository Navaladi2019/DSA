package matrix

import "testing"

type transposeTc struct {
	mat  [][]int
	want [][]int
}

var transposeTcs = func() []transposeTc {
	return []transposeTc{
		{[][]int{{1, 2, 3, 4},
			{5, 6, 7, 8},
			{9, 10, 11, 12}, {13, 14, 15, 16}},
			[][]int{{1, 5, 9, 13},
				{2, 6, 10, 14},
				{3, 7, 11, 15},
				{4, 8, 12, 16}},
		},
	}
}

func Test_transposeMatrix(t *testing.T) {

	for _, tc := range transposeTcs() {

		FindTranspose_1(tc.mat)

		for i := 0; i < len(tc.mat); i++ {
			for j := 0; j < len(tc.mat[i]); j++ {

				if tc.want[i][j] != tc.mat[i][j] {
					t.Error("has error in transpose")
				}
			}
		}
	}
}
