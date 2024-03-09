package recursion

import "testing"

func Test_RecursionGet1ToN(t *testing.T) {

	got := new([]int)

	RecursionGet1ToNOther(5, 0, got)

	for index, value := range *got {

		if index+1 != value {
			t.Errorf("failed for value index %d => %d", index, value)
		}
	}
}

func Test_FactorailTailREcursive(T *testing.T) {

	got := factorialTailRecursive(4, 5)

	if got != 120 {
		T.Errorf("error factoial")
	}
}
