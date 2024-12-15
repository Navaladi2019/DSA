package trie

import "testing"

func Test_CountDistinctRowsInMatrix(t *testing.T) {

	arr := [][]int{
		{1, 0, 0},
		{1, 1, 1},
		{1, 0, 0},
		{0, 1, 0},
	}

	got := CountDistinctRowsInMatrix(arr)

	if got != 3 {
		t.Error("has error in got")
	}
}
