package arrays

import "testing"

type RotateArrayTc struct {
	ip     []int
	rotate int
	want   []int
}

var RotateArrayTcs = func() []RotateArrayTc {
	return []RotateArrayTc{
		{[]int{1, 2, 3, 4, 5}, 1, []int{2, 3, 4, 5, 1}},
		{[]int{1, 2, 3, 4, 5}, 2, []int{3, 4, 5, 1, 2}},
		{[]int{1, 2, 3, 4, 5}, 5, []int{1, 2, 3, 4, 5}},
		{[]int{1, 2, 3, 4, 5}, 10, []int{1, 2, 3, 4, 5}},
	}
}

func Test_RorateArray(t *testing.T) {

	for _, tc := range RotateArrayTcs() {
		got := RotateArray(tc.ip, tc.rotate)

		for i := 0; i < len(tc.want); i++ {
			if tc.want[i] != got[i] {
				t.Error("Has Error")
				break
			}
		}
	}
}

func Test_RotateArrayLogic(t *testing.T) {

	for _, tc := range RotateArrayTcs() {
		got := RotateArrayLogic(tc.ip, tc.rotate)

		for i := 0; i < len(tc.want); i++ {
			if tc.want[i] != got[i] {
				t.Error("Has Error")
				break
			}
		}
	}
}
