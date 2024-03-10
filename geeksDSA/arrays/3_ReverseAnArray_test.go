package arrays

import "testing"

type ReverseArraytc struct {
	ip   []int
	want []int
}

var ReverseArraytcs = func() []ReverseArraytc {
	return []ReverseArraytc{
		{[]int{1, 2, 3, 4, 5}, []int{5, 4, 3, 2, 1}},
		{[]int{1}, []int{1}},
		{[]int{}, []int{}},
		{[]int{3, 2, 1, 0}, []int{0, 1, 2, 3}},
	}
}

func Test_ReverseArray(t *testing.T) {

	for _, tc := range ReverseArraytcs() {
		t.Run("tc", func(t *testing.T) {
			got := ReverseAnArray(tc.ip)

			for j := 0; j < len(tc.want); j++ {
				if got[j] != tc.want[j] {
					t.Error("failed")
				}
			}
		})
	}

}
