package sorting

import "testing"

type SmallElementTc struct {
	arr  []int
	k    int
	want int
}

var SmallElementTcs = func() []SmallElementTc {
	return []SmallElementTc{
		{[]int{1, 2, 3, 4, 5}, 3, 3},
		{[]int{100, 200, 300}, 2, 200},
		{[]int{5, 4, 3, 2, 1}, 3, 3},
		{[]int{5, 4, 3, 2, 1}, 6, -1},
	}
}

func Test_Find_Kth_SmallestNumber(t *testing.T) {

	for _, tc := range SmallElementTcs() {

		got := Find_Kth_SmallestNumber(tc.arr, tc.k)

		if got != tc.want {
			t.Error("kth smallest element not working")
		}
	}
}
