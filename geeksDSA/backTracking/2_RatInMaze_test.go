package backtracking

import "testing"

func Test_CanRatReachCheese(t *testing.T) {

	arr1 := [][]int{
		{1, 0, 0, 0},
		{1, 1, 0, 1},
		{0, 1, 0, 0},
		{1, 1, 1, 1},
	}

	arr2 := [][]int{
		{1, 0, 0},
		{0, 1, 0},
		{1, 1, 1},
	}

	GOT := CanRatReachCheese(arr1, 0, 0)

	if GOT != true {
		t.Error("ahs error in rat reach cheese")
	}

	GOT = CanRatReachCheese(arr2, 0, 0)

	if GOT != false {
		t.Error("ahs error in rat reach cheese")
	}
}
