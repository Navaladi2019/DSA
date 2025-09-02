package stringsDSA

import "testing"

func Test_ISSubSequence(t *testing.T) {

	got := ISSubSequence_2("ABCD", "AD")

	if got != true {
		t.Error("Has Error in subsequence")
	}

	got = ISSubSequence_2("ABCDE", "AED")

	if got != false {
		t.Error("Has Error in subsequence")
	}

	got = ISSubSequenceRecursion("ABCD", "AD", 4, 2)

	if got != true {
		t.Error("Has Error in subsequence")
	}

	got = ISSubSequenceRecursion("ABCDE", "AED", 5, 3)

	if got != false {
		t.Error("Has Error in subsequence")
	}
}
