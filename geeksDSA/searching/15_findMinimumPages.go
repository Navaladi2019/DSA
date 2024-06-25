package searching

import (
	"math"
)

func FindMinimumPages(arr []int, k int) int {

	if len(arr) == 1 {
		return arr[0]
	}
	if len(arr) == 0 {
		return 0
	}
	if k == 1 {
		return SumFromINdex(arr, 0, len(arr)-1)
	}

	res := math.MaxInt

	return res

}

func SumFromINdex(arr []int, s int, e int) int {
	sum := 0

	for i := s; i <= e; i++ {
		sum += arr[i]
	}

	return sum
}

func FindMinimumPages_Loop(arr []int, k int) int {

	if k == 1 {
		return SumFromINdex(arr, 0, len(arr)-1)
	}

	if len(arr) == 0 {
		return 0
	}

	if len(arr) == 1 {
		return arr[0]
	}

	res := math.MaxInt

	for i := 0; i < len(arr); i++ {

		res = min(res, max(FindMinimumPages_Loop(arr[i+1:], k-1), SumFromINdex(arr, 0, i)))
	}

	return res

}
