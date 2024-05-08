package sorting

import "slices"

func MinimunChoclatyedifferenceInPackets(arr []int, k int) int {

	slices.Sort(arr)

	minDif := arr[k-1] - arr[0]

	for i := 1; i < len(arr)-k; i++ {
		minDif = min(minDif, arr[i+k-1]-arr[i])
	}

	return minDif
}
