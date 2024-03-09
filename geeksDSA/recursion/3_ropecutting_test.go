package recursion

import "testing"

func Test_cuttingRope(t *testing.T) {

	got := find_MaximunCutsInRope(23, 13, 10, 18)

	if got != 2 {
		t.Error("did not get result")
	}

}
