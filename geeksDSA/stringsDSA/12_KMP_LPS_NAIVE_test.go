package stringsDSA

import "testing"

type LPSTc struct {
	str string
	LPS []int
}

type PatternTc struct {
	str                string
	patt               string
	matchingStartIndex int
}

var PatternTcs = func() []PatternTc {
	return []PatternTc{
		{"abcdefg", "bcd", 1},
		{"aaaaab", "aaaa", 0},
		{"abcd", "ba", -1},
	}
}

var LPSTcs = func() []LPSTc {
	return []LPSTc{
		{"ababc", []int{0, 0, 1, 2, 0}},
		{"aaaa", []int{0, 1, 2, 3}},
		{"abcd", []int{0, 0, 0, 0}},
		{"abacabad", []int{0, 0, 1, 0, 1, 2, 3, 0}},
		{"abbabb", []int{0, 0, 0, 1, 2, 3}},
		{"aabbcdebb", []int{0, 1, 0, 0, 0, 0, 0, 0}},
		{"aabbaa", []int{0, 1, 0, 0, 1, 2}},
		{"aabaa", []int{0, 1, 0, 1, 2}},
		{"aabbaa", []int{0, 1, 0, 0, 1, 2}},
	}
}

func Test_KMP_LPS_NAIVE(t *testing.T) {

	for _, tc := range LPSTcs() {
		got := FindLPS_Naive(tc.str)

		for i := 0; i < len(tc.LPS); i++ {

			if got[i] != tc.LPS[i] {
				t.Error("Has error in LPS")
			}
		}
	}

}

func Test_KMP_LPS(t *testing.T) {

	for _, tc := range LPSTcs() {
		got := FindLPS(tc.str)

		for i := 0; i < len(tc.LPS); i++ {

			if got[i] != tc.LPS[i] {
				t.Error("Has error in LPS")
			}
		}
	}

}

func Test_KMP(t *testing.T) {
	for _, tc := range PatternTcs() {

		got := FindPatternKMP(tc.str, tc.patt)

		if got != tc.matchingStartIndex {
			t.Error("string not matching")
		}
	}
}
