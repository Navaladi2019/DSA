package greedy

import "slices"

func ActivitySelection(arr [][]int) int {

	slices.SortFunc(arr, func(a []int, b []int) int { return a[1] - b[1] })
	prev := 0

	processToExecute := 1
	for i := 1; i < len(arr); i++ {

		if arr[prev][1] > arr[i][0] {
			prev = 1
			processToExecute++
		}
	}

	return processToExecute
}
