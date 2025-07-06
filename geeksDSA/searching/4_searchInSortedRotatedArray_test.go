package searching

import "testing"

type SearchInSOrtedRotatedtc struct {
	arr  []int
	n    int
	want int
}

var SearchInSOrtedRotatedtcs = func() []SearchInSOrtedRotatedtc {
	return []SearchInSOrtedRotatedtc{
		{[]int{10, 20, 40, 60, 5, 8}, 5, 4},
		{[]int{10, 20, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16}, 1, 2},
		{[]int{1, 2, 3, 4, 5, 6, 7, 8, 94}, 8, 7},
		{[]int{10, 20, 40, 60, 5, 8}, 8, 5},
		{[]int{10, 20, 40, 60, 5, 8}, 10, 0},
		{[]int{100, 200, 500, 1000, 2000, 10, 20}, 100, 0},
		{[]int{100, 200, 500, 1000, 2000, 10, 20}, 25, -1},
	}
}

func Test_SearchInSOrtedRotatedtcs(t *testing.T) {
	for _, tc := range SearchInSOrtedRotatedtcs() {

		got := FindInSortedRotatedArray(tc.arr, tc.n)

		if tc.want != got {
			t.Error("has error in sorted rortated array")
		}
	}
}
