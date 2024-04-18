package searching

import (
	"testing"
)

type squareRoottc struct {
	ip int
	op int
}

var squareRoottcs = func() []squareRoottc {
	return []squareRoottc{{16, 4}, {14, 3},
		{256, 16}, {1, 1}, {81, 9}, {99, 9}}
}

func Test_SquareRoot(t *testing.T) {

	for _, tc := range squareRoottcs() {

		got := findSquareRoor(tc.ip)

		if tc.op != got {
			t.Error("has error")
		}
	}
}
