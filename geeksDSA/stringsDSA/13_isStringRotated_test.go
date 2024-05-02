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
		{"abcd", "cdab", true},
		{"abaaa", "baaaa", true},
		{"abab", "abba", false},
	}
}

func Test_IsStringRotated(t *testing.T) {

	for _, tc := range stringRotatedTcs() {

		got := IsRotated(tc.original, tc.str)

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
