package matrix

import "testing"

type RotateMatrixTc struct {
	mat  [][]int
	want [][]int
}

func RotateMatrixTcs() []RotateMatrixTc {
	return []RotateMatrixTc{
		{[][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}, [][]int{{3, 6, 9}, {2, 5, 8}, {1, 4, 7}}},
	}
}

func Test_RotateArayby90Naval(t *testing.T) {

	for _, tc := range RotateMatrixTcs() {

		got := Rotateby90Naive(tc.mat)

		for i := 0; i < len(tc.mat); i++ {
			for j := 0; j < len(tc.mat[i]); j++ {

				if tc.want[i][j] != got[i][j] {
					t.Error("has error in transpose")
				}
			}
		}
	}
}

func Test_RotateArayby90(t *testing.T) {

	for _, tc := range RotateMatrixTcs() {

		RotateBy90(tc.mat)
		got := tc.mat
		for i := 0; i < len(tc.mat); i++ {
			for j := 0; j < len(tc.mat[i]); j++ {

				if tc.want[i][j] != got[i][j] {
					t.Error("has error in transpose")
				}
			}
		}
	}
}
