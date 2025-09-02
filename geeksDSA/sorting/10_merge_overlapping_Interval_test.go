package sorting

import "testing"

type MergeOverLapingInteralsTc struct {
	arr  []interval
	want []interval
}

func MergeOverLapingInteralsTcs() []MergeOverLapingInteralsTc {
	return []MergeOverLapingInteralsTc{
		{[]interval{{1, 3}, {2, 4}, {5, 7}, {6, 8}}, []interval{{1, 4}, {5, 8}}},
		{[]interval{{7, 9}, {6, 10}, {4, 5}, {1, 3}, {2, 4}}, []interval{{1, 5}, {6, 10}}},
	}
}

func Test_MergeOverLapingInterals_Naval(t *testing.T) {

	for _, tc := range MergeOverLapingInteralsTcs() {

		got := MergeOverlappingINtervals(tc.arr)

		for i := 0; i < len(tc.want); i++ {

			if got[i] != tc.want[i] {
				t.Error("Has error in Merge Overrlapping")
			}
		}

	}
}
