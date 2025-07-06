package arrays

import "testing"

type consucutiveOnetc struct {
	ip   []int
	want int
}

var consucutiveOnetcs = func() []consucutiveOnetc {

	return []consucutiveOnetc{
		{[]int{0, 1, 1, 0, 1, 0}, 2},
		{[]int{0, 1, 1, 1, 0, 1, 0}, 3},
		{[]int{0, 0}, 0},
		{[]int{0, 1, 1, 1, 0, 1, 1, 1, 1, 1, 1}, 6},
	}
}

func Test_consucutiveOnetcs(t *testing.T) {

	for _, tc := range consucutiveOnetcs() {

		got := FindMaxConsucitiveones(tc.ip)
		if got != tc.want {
			t.Error("has error")
		}
	}
}
