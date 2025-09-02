package matrix

import "testing"

type SearchInSortedMatyricTc struct {
	mat    [][]int
	target int
	x      int
	y      int
}

var SearchInSortedMatyricTcs = func() []SearchInSortedMatyricTc {

	return []SearchInSortedMatyricTc{
		{[][]int{{10, 20, 30, 40}, {15, 25, 35, 45}, {27, 29, 37, 48}, {32, 33, 39, 50}}, 27, 2, 0},
	}
}

func Test_SearchInSortedMatrix(t *testing.T) {

	for _, tc := range SearchInSortedMatyricTcs() {

		x, y := SearchInRowColumnSortedMatrix_1(tc.mat, tc.target)

		if x != tc.x && y != tc.y {
			t.Error("has error in matrix sorting")
		}
	}
}
