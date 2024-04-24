package sorting

import "testing"

func Test_MergecYCLEsORTdUPLICATE(t *testing.T) {

	for _, tc := range SortingTcs() {

		CycleSortDuplicate(tc.ip)

		for i, _ := range tc.want {

			if tc.want[i] != tc.ip[i] {
				t.Error("MergeSort not working")
			}
		}

	}
}
