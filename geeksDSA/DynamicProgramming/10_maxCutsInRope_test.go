package dynamicprogramming

import "testing"

func Test_FindMaxCutsInRope_Recursion(t *testing.T) {

	got := FindMaxCutsInRope_Recursion([]int{1, 2, 3}, 5)

	if got != 5 {
		t.Error("has error in finding max cuts")
	}
}
