package stringsDSA

import "testing"

func Test_LongestDistinctSubString(t *testing.T) {

	got := longEstSubstring_1("abcdabc")

	if got != 4 {
		t.Error("Has error in longest substring")
	}

	got = longEstSubstring_1("abaacdbab")

	if got != 4 {
		t.Error("Has error in longest substring 2")
	}

	got = longEstSubstring_1("aaa")

	if got != 1 {
		t.Error("Has error in longest substring aa")
	}

	got = longEstSubstring_1("")

	if got != 0 {
		t.Error("Has error in longest substring em")
	}
}
