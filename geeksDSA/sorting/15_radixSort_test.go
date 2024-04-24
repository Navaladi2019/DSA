package sorting

import "testing"

func Test_RadixSort(t *testing.T) {

	for _, tc := range SortingTcs() {
		RadixSort(tc.ip)

		for i := 0; i < len(tc.ip); i++ {

			if tc.want[i] != tc.ip[i] {

				t.Error("Radix Sort not working")
			}
		}

	}
}
