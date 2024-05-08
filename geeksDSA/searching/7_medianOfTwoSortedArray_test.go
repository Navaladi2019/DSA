package searching

import "testing"

type MedianOfSortedArrayTc struct {
	arr1 []int
	arr2 []int
	want int
}

func MedianOfSortedArrayTcs() []MedianOfSortedArrayTc {
	return []MedianOfSortedArrayTc{
		{[]int{10, 20, 30, 40, 50}, []int{5, 15, 25, 30, 45}, 27},
		{[]int{1, 2, 3, 4, 5, 6}, []int{10, 20, 30, 40, 50}, 6},
		{[]int{10, 20, 30, 40, 50, 60}, []int{1, 2, 3, 4, 5}, 10},
	}
}

func Test_MeidanOfTwoSortedArray(t *testing.T) {

	for _, tc := range MedianOfSortedArrayTcs() {

		got := FindMedianOfTwoSortedArray(tc.arr1, tc.arr2)

		if got != tc.want {
			t.Error("has error")
		}
	}

}
