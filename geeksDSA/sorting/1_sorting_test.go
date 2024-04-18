package sorting

import "testing"

type SortingTc struct {
	ip   []int
	want []int
}

var SortingTcs = func() []SortingTc {
	return []SortingTc{
		{[]int{1, 35, 2, 8, 9, 20, 0}, []int{0, 1, 2, 8, 9, 20, 35}},
		{[]int{10, 5, 8, 20, 2, 18}, []int{2, 5, 8, 10, 18, 20}},
		{[]int{2, 1}, []int{1, 2}},
		{[]int{1}, []int{1}},
	}
}

type MergeSortTc struct {
	ip   []int
	low  int
	high int
	op   []int
}

var MergeSortTcs = func() []MergeSortTc {

	return []MergeSortTc{
		{[]int{2, 3, 4, 5, 1, 100, 101, 102, 102}, 0, 8, []int{1, 2, 3, 4, 5, 100, 101, 102, 102}},
	}
}

func Test_MergeSort(t *testing.T) {

	for _, tc := range MergeSortTcs() {

		MergeSort(tc.ip, tc.low, tc.high)

		for i, _ := range tc.op {

			if tc.op[i] != tc.ip[i] {
				t.Error("MergeSort not working")
			}
		}

	}
}

func Test_BubbleSort(t *testing.T) {

	for _, tc := range SortingTcs() {

		BubbleSort(tc.ip)

		for i := 0; i < len(tc.want); i++ {
			if tc.want[i] != tc.ip[i] {
				t.Error("sorting not working in bubble sort.")
			}
		}
	}
}
