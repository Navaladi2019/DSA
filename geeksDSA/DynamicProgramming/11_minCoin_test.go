package dynamicprogramming

import (
	"testing"
)

func Test_GetMinCoins(t *testing.T) {

	got := GetMinCoinsDP([]int{25, 10, 5}, 30)

	if got != 2 {
		t.Error("as error in coin")
	}

	got = GetMinCoinsDP([]int{9, 6, 5, 1}, 11)

	if got != 2 {
		t.Error("as error in coin")
	}
}
