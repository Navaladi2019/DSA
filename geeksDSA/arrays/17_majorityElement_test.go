package arrays

import "testing"

type MajEleTc struct {
	arr  []int
	want int
}

var MajEleTcs = func() []MajEleTc {
	return []MajEleTc{
		{[]int{8, 3, 4, 8, 8}, 4},
		{[]int{3, 7, 4, 7, 7, 5}, -1},
		{[]int{20, 30, 40, 50, 50, 50, 50}, 6},
	}
}

func Test_MajEle(t *testing.T) {

	for _, tc := range MajEleTcs() {

		got := FindMajorityElement(tc.arr)

		if tc.want != got {
			t.Error("has error")
		}
	}
}
