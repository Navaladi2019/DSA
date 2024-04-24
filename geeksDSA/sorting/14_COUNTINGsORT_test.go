package sorting

import "testing"

type countSortTc struct {
	arr []int
	max int
	op  []int
}

func countSortTcs() []countSortTc {
	return []countSortTc{
		{[]int{1, 4, 4, 1, 0, 1}, 4, []int{0, 1, 1, 1, 4, 4}},
	}
}

func Test_CountSort(t *testing.T) {

	for _, tc := range countSortTcs() {

		got := CountingSort(tc.arr, tc.max)

		for i := 0; i < len(tc.arr); i++ {

			if got[i] != tc.op[i] {
				t.Error(("Counting sort not working"))
			}
		}
	}
}
