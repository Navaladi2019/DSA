package stringsDSA

import "testing"

func Test_isAnagram(t *testing.T) {

	got := IsAnagram_SmallCase("listen", "silent")

	if got != true {
		t.Error("has Error in anagram")
	}

	got = IsAnagram_SmallCase("aaacb", "cabaa")

	if got != true {
		t.Error("has Error in anagram")
	}

	got = IsAnagram_SmallCase("aab", "bab")

	if got != false {
		t.Error("has Error in anagram")
	}
}
