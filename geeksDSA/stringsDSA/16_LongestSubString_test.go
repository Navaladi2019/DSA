package stringsDSA

import "testing"

func Test_LongestDistinctSubString(t *testing.T) {

	got := LongestSubStringEfficient("abcdabc")

	if got != 4 {
		t.Error("Has error in longest substring")
	}

	got = LongestSubStringEfficient("abaacdbab")

	if got != 4 {
		t.Error("Has error in longest substring 2")
	}

	got = LongestSubStringEfficient("aaa")

	if got != 1 {
		t.Error("Has error in longest substring aa")
	}

	got = LongestSubStringEfficient("")

	if got != 0 {
		t.Error("Has error in longest substring em")
	}
}
