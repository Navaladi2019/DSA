package arrays

import "testing"

type MaxLengthOddEvenArraytc struct {
	ip   []int
	want int
}

var MaxLengthOddEvenArraytcs = func() []MaxLengthOddEvenArraytc {
	return []MaxLengthOddEvenArraytc{
		{[]int{10, 12, 14, 7, 8}, 3},
		{[]int{7, 10, 13, 14}, 4},
		{[]int{10, 12, 8, 4}, 1},
	}
}

func Test_MaxLengthOddEvenArray(t *testing.T) {
	for _, tc := range MaxLengthOddEvenArraytcs() {
		got := FindLonogestEvenOddSubArray(tc.ip)

		if got != tc.want {
			t.Error("has error")
		}
	}
}
