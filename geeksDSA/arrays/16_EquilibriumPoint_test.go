package arrays

import "testing"

type ArrayHasEquilibriumPointTc struct {
	ip   []int
	want bool
}

var ArrayHasEquilibriumPointTcs = func() []ArrayHasEquilibriumPointTc {

	return []ArrayHasEquilibriumPointTc{
		{[]int{3, 4, 8, -9, 20, 6}, true},
		{[]int{4, 2, -2}, true},
		{[]int{4, 2, 2}, false},
	}
}

func Test_ArrayHasEquilibriumPoint(t *testing.T) {

	for _, tc := range ArrayHasEquilibriumPointTcs() {
		got := IsArrayHasEquilibriumPointEfficientSolution(tc.ip)

		if got != tc.want {
			t.Error("Has Error")
		}

	}
}
