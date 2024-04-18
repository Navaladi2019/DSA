package searching

import "testing"

type findIndexBinaryArrayIterativetc struct {
	ip   []int
	ip1  int
	want int
}

var findIndexBinaryArrayIterativetcs = func() []findIndexBinaryArrayIterativetc {
	return []findIndexBinaryArrayIterativetc{
		//{[]int{10, 20, 30, 40, 50, 60}, 20, 1},
		{[]int{5, 15, 25}, 25, 2},
		{[]int{5, 10, 15, 21, 22, 30, 100, 200, 300}, 20, -1},
	}
}

var findFirstIndexBinaryArraytcs = func() []findIndexBinaryArrayIterativetc {
	return []findIndexBinaryArrayIterativetc{
		{[]int{5, 15, 25}, 25, 2},
		{[]int{10, 20, 30, 40, 50, 60}, 20, 1},
		{[]int{5, 10, 15, 21, 22, 30, 100, 200, 300}, 20, -1},
		{[]int{5, 15, 15, 15, 16, 17, 18, 19, 19, 19, 20, 20, 20, 20, 21, 22, 22, 22, 22, 22, 23, 23, 23, 23, 23, 22}, 22, 15},
	}
}

func Test_findIndexBinaryArrayIterativetcs(t *testing.T) {
	for _, tc := range findIndexBinaryArrayIterativetcs() {

		got := findIndexBinaryArrayIterative(tc.ip, tc.ip1)

		if got != tc.want {
			t.Error("has error")
		}
	}
}

func Test_findFirstIndexBinaryArrayIterativetcs(t *testing.T) {
	for _, tc := range findFirstIndexBinaryArraytcs() {

		got := findIndexBinaryArrayFirstIndex(tc.ip, tc.ip1)

		if got != tc.want {
			t.Error("has error")
		}
	}
}

func Test_findIndexBinaryArrayRecursivetcs(t *testing.T) {
	for _, tc := range findIndexBinaryArrayIterativetcs() {

		got := findIndexBinaryArrayRecursive(tc.ip, tc.ip1, len(tc.ip)-1, 0)

		if got != tc.want {
			t.Error("has error")
		}
	}
}
