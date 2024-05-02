package stringsDSA

import "testing"

func Test_LexographicRank(t *testing.T) {

	got := LexographicRank("BAC")
	if got != 3 {
		t.Error("has error in LexographicRank")
	}

	got = LexographicRank("CBA")
	if got != 6 {
		t.Error("has error in LexographicRank")
	}

	got = LexographicRank("DCBA")
	if got != 24 {
		t.Error("has error in LexographicRank")
	}

	got = LexographicRank("STRING")
	if got != 598 {
		t.Error("has error in LexographicRank")
	}
}
