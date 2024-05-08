package sorting

import "testing"

type Inversiontc struct {
	arr  []int
	low  int
	high int
	want int
}

func Inversiontcs() []Inversiontc {
	return []Inversiontc{
		{[]int{2, 5, 8, 11, 3, 6, 9, 13}, 0, 7, 6},
	}
}

func Test_CountINversion(t *testing.T) {

	for _, tc := range Inversiontcs() {

		got := CountInversion(tc.arr, tc.low, tc.high)

		if got != tc.want {
			t.Error("has error in inevrsion")
		}
	}
}
