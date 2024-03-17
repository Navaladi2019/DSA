package arrays

import "testing"

type MaxSumWindowSlidingLogictc struct {
	ip   []int
	ipk  int
	want int
}

var MaxSumWindowSlidingLogictcs = func() []MaxSumWindowSlidingLogictc {
	return []MaxSumWindowSlidingLogictc{
		{[]int{10, 5, -2, 20, 1}, 2, 21},
		{[]int{1, 8, 30, -5, 20, 7}, 3, 45},
		{[]int{5, -10, 6, 90, 3}, 2, 96},
	}
}

func Test_MaxSumWindowSlidingLogictcs(t *testing.T) {

	for _, tc := range MaxSumWindowSlidingLogictcs() {
		got := GetMaxSumWindowSlidingLogic(tc.ip, tc.ipk)

		if got != tc.want {
			t.Error("erroer")
		}
	}
}
