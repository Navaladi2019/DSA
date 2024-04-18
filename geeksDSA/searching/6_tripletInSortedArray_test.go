package searching

import "testing"

type triplet struct {
	arr  []int
	n    int
	want bool
}

var tripletTcs = func() []triplet {
	return []triplet{
		{[]int{5, 8, 10, 20, 40, 60}, 70, true},
		{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16}, 9, true},
		{[]int{1, 2, 3, 4, 5, 6, 7, 8, 94}, 14, true},
		{[]int{100, 200, 500, 1000, 2000}, 15000, false},
		{[]int{100, 200, 500, 1000, 2000}, 10, false},
	}
}

func Test_FindTripletInSortedArray(t *testing.T) {
	for _, tc := range tripletTcs() {

		got := FindTripletInSortedArray(tc.arr, tc.n)

		if got != tc.want {
			t.Error("has error ")
		}
	}
}
