package arrays

import "testing"

type rainwatertc struct {
	ip   []int
	want int
}

var rainwatertcs = func() []rainwatertc {

	return []rainwatertc{
		//{[]int{2, 0, 2}, 2},
		{[]int{3, 0, 1, 2, 5}, 6},
		{[]int{30, 20, 10}, 0},
		{[]int{1, 2, 3}, 0},
		{[]int{5, 4, 0, 1, 2}, 3},
	}
}

func Test_rainwatertcs(t *testing.T) {

	for _, tc := range rainwatertcs() {

		got := FindAmountofTrappingRainWater(tc.ip)
		if got != tc.want {
			t.Error("has error")
		}
	}
}
