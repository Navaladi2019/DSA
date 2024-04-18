package sorting

import "testing"

type UnionTc struct {
	arr  []int
	arr2 []int
	res  []int
}

var UnionTcs = func() []UnionTc {
	return []UnionTc{{[]int{3, 5, 8}, []int{2, 8, 9, 10, 15}, []int{2, 3, 5, 8, 9, 10, 15}},
		{[]int{2, 3, 3, 3, 4, 4}, []int{4, 4}, []int{2, 3, 4}},
	}
}

func Test_TestUnion(t *testing.T) {
	for _, tc := range UnionTcs() {

		got := finduionOfTwoArrays(tc.arr, tc.arr2)

		for i, _ := range tc.res {
			if tc.res[i] != got[i] {
				t.Error("Union not working")
			}
		}
	}
}
