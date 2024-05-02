package stringsDSA

import "testing"

func Test_ISAnagramPattern(t *testing.T) {

	got := IsAnagramPatternPresent("geeksforgeeksaa", "frog")

	if got != 5 {
		t.Error("has error in anagram pattern")
	}

	got = IsAnagramPatternPresent("geeksforgeeks", "rseek")

	if got != -1 {
		t.Error("has error in anagram pattern")
	}
}
