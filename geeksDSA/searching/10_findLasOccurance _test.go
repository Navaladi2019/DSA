package searching

import "testing"

var findIndexLastOccuranceTcs = func() []findIndexBinaryArrayIterativetc {
	return []findIndexBinaryArrayIterativetc{
		{[]int{1, 10, 10, 10, 20, 20, 40}, 20, 5},
		{[]int{10, 20, 30}, 25, -1},
		{[]int{15, 15, 15}, 15, 2},
	}
}

func Test_findIndexLastOccuranceTcs(t *testing.T) {
	for _, tc := range findIndexLastOccuranceTcs() {

		got := FindLastOccurance(tc.ip, tc.ip1)

		if got != tc.want {
			t.Error("has error")
		}
	}
}
