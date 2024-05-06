package searching

import "testing"

var findIndexFirstOccuranceTcs = func() []findIndexBinaryArrayIterativetc {
	return []findIndexBinaryArrayIterativetc{
		{[]int{1, 10, 10, 10, 20, 20, 40}, 20, 4},
		{[]int{10, 20, 30}, 25, -1},
		{[]int{15, 15, 15}, 15, 0},
	}
}

func Test_findIndexFirstOccurance(t *testing.T) {
	for _, tc := range findIndexFirstOccuranceTcs() {

		got := FindFirstOccurance(tc.ip, tc.ip1)

		if got != tc.want {
			t.Error("has error")
		}
	}
}
