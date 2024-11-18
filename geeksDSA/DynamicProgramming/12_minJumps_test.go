package dynamicprogramming

import "testing"

func Test_MinJumps(t *testing.T) {

	got := MinJumps([]int{3, 4, 2, 1, 2, 1}, 0)

	if got != 2 {
		t.Errorf("min jumps error")
	}
}
