package sorting

import "testing"

type SortTwoTypesArraytc struct {
	arr  []int
	want []int
	fn   FnIsArrayType
}

var SortTwoTypesArraytcs = func() []SortTwoTypesArraytc {
	return []SortTwoTypesArraytc{
		{[]int{1, 0, 1, 1, 0, 1, 0}, []int{0, 0, 0, 1, 1, 1, 1}, isArrayTypeZero},
	}
}

func Test_SortTwoTypesArray(t *testing.T) {

	for _, tc := range SortTwoTypesArraytcs() {

		SortTwoTypesArray(tc.arr, tc.fn)

		for i := 0; i < len(tc.want); i++ {

			if tc.arr[i] != tc.want[i] {
				t.Error("Has error")
			}
		}
	}
}
