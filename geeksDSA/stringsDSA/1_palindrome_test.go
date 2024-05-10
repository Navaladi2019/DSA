package stringsDSA

import "testing"

func Test_IsPalindrome(t *testing.T) {

	got := IsPalindrome("ABCBA")

	if got != true {
		t.Error("is palindrome fail")
	}

}
