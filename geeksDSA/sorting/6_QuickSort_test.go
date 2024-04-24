package sorting

import "testing"

func Test_QuickSortUsingLomutoPartition(t *testing.T) {

	for _, tc := range SortingTcs() {

		QuickSortUsingLomutoPartition(tc.ip, 0, len(tc.ip)-1)

		for i, _ := range tc.want {

			if tc.want[i] != tc.ip[i] {
				t.Error("QuickSortUsingLomutoPartition not working")
			}
		}

	}
}

func Test_QuickSortUsingHoaresPartition(t *testing.T) {

	for _, tc := range SortingTcs() {

		QuickSortUsingHoaresPartition(tc.ip, 0, len(tc.ip)-1)

		for i, _ := range tc.want {

			if tc.want[i] != tc.ip[i] {
				t.Error("QuickSortUsingHoaresPartition not working")
			}
		}

	}
}
