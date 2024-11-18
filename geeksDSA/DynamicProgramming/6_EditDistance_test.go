package dynamicprogramming

import "testing"

func Test_EditDistanceToGetSameString(t *testing.T) {

	got := EditDistanceToGetSameString("CAT", "CUT", 3, 3)

	if got != 1 {
		t.Error("has error in naive")
	}

	got = EditDistanceToGetSameString_DP("CAT", "CUT")

	if got != 1 {
		t.Error("has error in DP")
	}

}
