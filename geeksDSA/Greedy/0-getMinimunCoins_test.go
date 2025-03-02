package greedy

import "testing"

func Test_GetMinCoins(t *testing.T) {

	coints := []int{10, 5, 2, 1}

	got := GetMinCoins(coints, 52)

	if got != 6 {
		t.Error("error in coins")
	}
}
