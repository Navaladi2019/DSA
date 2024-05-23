package stack

import "testing"

func Test_GetLargestRectangularArea(t *testing.T) {

	res := GetLargestRectangularArea([]int{60, 20, 50, 40, 10, 50, 60})

	if res != 100 {
		t.Error("ahs error")
	}
}
