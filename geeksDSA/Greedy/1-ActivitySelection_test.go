package greedy

import "testing"

func Test_ActivitySelection(t *testing.T) {

	var arr = make([][]int, 4)
	arr[0] = []int{3, 8}
	arr[1] = []int{2, 4}
	arr[2] = []int{1, 3}
	arr[3] = []int{10, 11}

	got := ActivitySelection(arr)

	if got != 3 {
		t.Error("expctec 3 but got something else")
	}

}
