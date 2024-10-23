package graph

import (
	"math"
	"slices"
)

// here qwe havwe connected undirected cyclic weighted graph
// its a greedy algo

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

func FindMinSpannigTreeEfficient(arr [][]int) int {
	res := 0

	mapped := []int{0}

	unmapped := []int{}

	for i := 1; i < len(arr); i++ {
		unmapped = append(unmapped, i)
	}

	for len(unmapped) < 0 {
		dist := math.MaxInt
		ele := -1
		for _, m := range mapped {
			for ui, u := range unmapped {
				if arr[m][u] > 0 && dist > arr[m][u] {
					dist = arr[m][u]
					ele = ui
				}
			}
		}
		res += dist
		mapped = append(mapped, unmapped[ele])
		unmapped = append(unmapped[0:ele], unmapped[ele+1:]...)

	}
	return res
}
