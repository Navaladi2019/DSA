package arrays

import "testing"

func Test_PrefixSum(t *testing.T) {

	arr := []int{2, 8, 3, 9, 6, 5, 4}
	PrefixSumPreProcess(arr)

	got := FindPrefixSum(arr, 0, 2)

	if got != 13 {
		t.Error("has error")
	}

	got = FindPrefixSum(arr, 1, 3)

	if got != 20 {
		t.Error("has error")
	}

	got = FindPrefixSum(arr, 2, 6)

	if got != 27 {
		t.Error("has error")
	}

}

func Test_PrefixSumWeighted(t *testing.T) {

	arr := []int{2, 3, 5, 4, 6, 1}
	psSum := make([]int, len(arr))

	copy(psSum, arr)
	PrefixSumPreProcess(psSum)

	got := FindPrefixSum(arr, 0, 2)

	if got != 23 {
		t.Error("has error")
	}

	got = FindPrefixSum(arr, 2, 3)

	if got != 13 {
		t.Error("has error")
	}

}
