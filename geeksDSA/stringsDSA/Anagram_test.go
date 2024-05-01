package stringsDSA

import "testing"

func Test_isAnagram(t *testing.T) {

	got := isAnagram_Dict("listen", "silent")

	if got != true {
		t.Error("has Error in anagram")
	}

	got = isAnagram_Dict("aaacb", "cabaa")

	if got != true {
		t.Error("has Error in anagram")
	}

	got = isAnagram_Dict("aab", "bab")

	if got != false {
		t.Error("has Error in anagram")
	}
}
