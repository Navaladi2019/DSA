package sorting

import (
	"math"
	"slices"
)

func FindMinimunDifference(arr []int) int {

	mindif := math.MaxInt32

	slices.Sort(arr)

	for i := 0; i < len(arr)-2; i++ {
		mindif = min(mindif, int(math.Abs(float64(arr[i+1]-arr[i]))))
	}

	return mindif
}
