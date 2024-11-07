package dynamicprogramming

import "testing"

func Test_CoinChangeNaive(t *testing.T) {

	got := CoinChangeNaive([]int{1, 2, 3}, 2, 4)

	if got != 4 {
		t.Error("Has error in naive")
	}

	got = CoinChange_DP([]int{1, 2, 3}, 4)

	if got != 4 {
		t.Error("Has error in dp")
	}
}
