package arrays

import "testing"

type isArraySortedTc struct {
	ip   []int
	want bool
}

var isArraySortedTcs = func() []isArraySortedTc {
	return []isArraySortedTc{
		{[]int{8, 12, 15}, true},
		{[]int{8, 10, 10, 12}, true},
		{[]int{100}, true},
		{[]int{100, 20, 200}, false},
	}
}

func Test_isArraySorted(t *testing.T) {

	for _, tc := range isArraySortedTcs() {
		got := findIsArraySorted(tc.ip)
		if tc.want != got {
			t.Error("has Error")
		}
	}
}
