package graph

import (
	"math"
	"slices"
)

func FindMinSpanningTreePrimsAlgo(arr [][]int) int {

	mapped := []int{0}
	unMapped := []int{}

	dist := 0
	for i := 1; i < len(arr); i++ {
		unMapped = append(unMapped, i)
	}

	for i := 0; i < len(arr)-1; i++ {
		minDist := math.MaxInt
		mappedEleIndex := -1
		for _, u := range mapped {
			for i, v := range unMapped {
				if arr[u][v] > 0 && arr[u][v] <= minDist {
					minDist = arr[u][v]
					mappedEleIndex = i
				}
			}
		}

		dist += minDist
		mapped = append(mapped, unMapped[mappedEleIndex])
		unMapped = slices.Delete(unMapped, mappedEleIndex, mappedEleIndex+1)
	}

	return dist
}
