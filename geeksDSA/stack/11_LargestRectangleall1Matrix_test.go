package stack

import "testing"

func Test_LargestRectange(t *testing.T) {

	got := LargestRectange([][]int{
		{0, 1, 1, 0},
		{1, 1, 1, 1},
		{1, 1, 1, 1},
		{1, 1, 0, 0},
	})

	if got != 8 {
		t.Error("has error")
	}
}
