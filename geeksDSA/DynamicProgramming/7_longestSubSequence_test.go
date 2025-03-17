package dynamicprogramming

import "testing"

// conside tyhe array [3,4,2,8,10,5,1] for ths the longes sub sequence is [3,4,8,10]
// because here by having 3 we do not consider numbers less than tha

//[3],[3,4],[3,4,8],[3,4,8,10],[3,5],[3,8,10],[3,8],[3,10] so for each index we will be computing the subsequence

//3,4,5,12,15,17,17

func Test_LongestSubsequence(t *testing.T) {

	got := LongestSubsequence([]int{3, 4, 2, 8, 10, 5, 1})

	if got != 4 {
		t.Errorf("Error in longest subsequence on^2")
	}

	got = LongestIncreasingSubSequence([]int{3, 4, 2, 8, 10, 5, 1})

	if got != 4 {
		t.Errorf("Error in longest subsequence on^2")
	}
}
