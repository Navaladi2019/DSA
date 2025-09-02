package stringsDSA

import "testing"

type stringRotatedTc struct {
	str       string
	original  string
	IsRotated bool
}

var stringRotatedTcs = func() []stringRotatedTc {
	return []stringRotatedTc{
		{"abcd", "dcba", false},
		{"abcd", "abcd", true},
		{"abcd", "cdab", true},
		{"abaaa", "baaaa", true},
		{"abab", "abba", false},
		{"aba", "aab", true},
	}
}

func Test_IsStringRotated(t *testing.T) {

	for _, tc := range stringRotatedTcs() {

		got := IsRotated_1(tc.original, tc.str)

		if got != tc.IsRotated {
			t.Error("Is rotrted error")
		}
	}
}

func Test_IsStringRotatedEasy(t *testing.T) {

	for _, tc := range stringRotatedTcs() {

		got := IsRotatedEasyApproach(tc.original, tc.str)

		if got != tc.IsRotated {
			t.Error("Is rotrted error")
		}
	}
}
