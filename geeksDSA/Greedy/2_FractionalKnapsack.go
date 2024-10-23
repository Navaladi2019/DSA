package greedy

import "slices"

func FractionalKnapSack(arr []Knapsack, cap int) float32 {

	slices.SortFunc(arr, func(a Knapsack, b Knapsack) int {

		if float32(a.value)/float32(a.weight) > float32(b.value)/float32(b.weight) {
			return -1
		}
		return 1
	})

	var currWeight int = 0
	var currValue float32 = 0

	for i := 0; i < len(arr); i++ {

		if currWeight+arr[i].weight < cap {
			currWeight += arr[i].weight
			currValue += float32(arr[i].value)
		} else {
			currValue += (float32(arr[i].value) / float32(arr[i].weight)) * (float32(cap) - float32(currWeight))
			break
		}

	}

	return currValue
}

type Knapsack struct {
	weight int
	value  int
}
