package greedy

import (
	"testing"
)

func Test_FractionalKnapSack(t *testing.T) {

	arr1 := make([]Knapsack, 4)
	arr1[0] = Knapsack{weight: 22, value: 19}
	arr1[1] = Knapsack{weight: 10, value: 9}
	arr1[2] = Knapsack{weight: 9, value: 9}
	arr1[3] = Knapsack{weight: 7, value: 6}

	got := FractionalKnapSack(arr1, 25)

	if got != 23.181818 {
		t.Error("knapsacl error")
	}
}

func Test_findClosestNumber(t *testing.T) {

	nums := []int{-10, -12, -54, -12, -544, -10000}
	lastMin := 0

	for i := 1; i < len(nums); i++ {

		last := uint(nums[lastMin])

		curr := uint(nums[i])
		if last > curr {
			lastMin = i
		}

	}
}
