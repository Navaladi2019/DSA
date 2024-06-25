package searching

import "testing"

func Test_FindMinimumPages(t *testing.T) {

	got := FindMinimumPages_Loop([]int{10, 20, 30, 40}, 2)

	if got != 60 {
		t.Error("expected 60")
	}

	got = FindMinimumPages_Loop([]int{10, 20, 30}, 1)

	if got != 60 {
		t.Error("expected 60")
	}

	got = FindMinimumPages_Loop([]int{10, 5, 30, 1, 2, 5, 10, 10}, 3)

	if got != 30 {
		t.Error("expected 30")
	}

}
