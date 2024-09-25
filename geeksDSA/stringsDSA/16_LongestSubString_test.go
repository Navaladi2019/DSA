package stringsDSA

import "testing"

func Test_LongestDistinctSubString(t *testing.T) {

	got := LongestSubstringPractice("abcdabc")

	if got != 4 {
		t.Error("Has error in longest substring")
	}

	got = LongestSubstringPractice("abaacdbab")

	if got != 4 {
		t.Error("Has error in longest substring 2")
	}

	got = LongestSubstringPractice("aaa")

	if got != 1 {
		t.Error("Has error in longest substring aa")
	}

	got = LongestSubstringPractice("")

	if got != 0 {
		t.Error("Has error in longest substring em")
	}
}
