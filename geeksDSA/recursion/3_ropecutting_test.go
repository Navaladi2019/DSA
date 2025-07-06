package recursion

import "testing"

func Test_cuttingRope(t *testing.T) {

	got := find_MaximunCutsInRope_2(23, 13, 10, 18)

	if got != 2 {
		t.Error("did not get result")
	}

	got = find_MaximunCutsInRope_2(5, 2, 5, 1)

	if got != 5 {
		t.Error("did not get result")
	}

	got = find_MaximunCutsInRope_2(23, 12, 9, 11)

	if got != 2 {
		t.Error("did not get result")
	}

	got = find_MaximunCutsInRope_2(5, 4, 2, 6)

	if got != -1 {
		t.Error("did not get result")
	}

}
