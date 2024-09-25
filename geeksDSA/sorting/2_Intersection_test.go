package sorting

import "testing"

type InsertionTc struct {
	arr  []int
	arr2 []int
	res  []int
}

var InsertionTcs = func() []InsertionTc {

	return []InsertionTc{
		{[]int{3, 5, 10, 10, 10, 15, 15, 20}, []int{5, 10, 10, 15, 30}, []int{5, 10, 15}},
		{[]int{1, 1, 3, 3, 3}, []int{1, 1, 1, 1, 3, 5, 7}, []int{1, 3}},
	}
}

func Test_Insertion(t *testing.T) {
	for _, tc := range InsertionTcs() {

		got := FindInterSectionOdTwoSortedArray_1(tc.arr, tc.arr2)

		for i, _ := range tc.res {
			if tc.res[i] != got[i] {
				t.Error("Failed")
			}
		}
	}
}
