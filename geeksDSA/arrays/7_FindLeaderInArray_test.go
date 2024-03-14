package arrays

import "testing"

type LeaderArrayTc struct {
	ip   []int
	want []int
}

var LeaderArrayTcs = func() []LeaderArrayTc {
	return []LeaderArrayTc{
		{[]int{7, 10, 4, 3, 6, 5, 2}, []int{10, 6, 5, 2}},
		{[]int{10, 20, 30}, []int{30}},
		{[]int{30, 20, 10}, []int{30, 20, 10}},
	}
}

func Test_LeaderArrayTcs(t *testing.T) {

	for _, tc := range LeaderArrayTcs() {
		got := FindLeaderInArrsy(tc.ip)

		if len(tc.want) != len(got) {
			t.Error(("has error"))
		}
	}
}
