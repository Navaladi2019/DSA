package sorting

import (
	"testing"
)

func TestHoaresPartition(t *testing.T) {
	tests := []struct {
		input    []int
		low      int
		high     int
		expected int
	}{
		{[]int{3, 6, 8, 10, 1, 2, 1}, 0, 6, 2},      // Pivot at index 3
		{[]int{9, 8, 7, 6, 5, 4, 3, 2, 1}, 0, 8, 7}, // Pivot at index 4
		{[]int{}, 0, -1, -1},                        // Empty array
		{[]int{1}, 0, 0, 0},
		//{[]int{4, 5, 3, 4, 4, 2}, 0, 5, 2},
		{[]int{2, 1}, 0, 1, 0}, // Single element array
	}

	for _, test := range tests {
		arr := make([]int, len(test.input))
		copy(arr, test.input)
		pivotIndex := Hoares_Partition_1(arr, test.low, test.high)
		if pivotIndex != test.expected {
			t.Errorf("For input %v and range [%d, %d], expected pivot index %d, but got %d", test.input, test.low, test.high, test.expected, pivotIndex)
		}
	}
}
