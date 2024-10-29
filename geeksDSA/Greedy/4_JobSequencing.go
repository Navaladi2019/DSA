package greedy

import (
	"slices"
)

type JobSequencing struct {
	DeadLine int
	Profit   int
}

func GetMaximumProfitOnJobSequencing(arr []JobSequencing) int {

	maxDeadLineHrs := 0
	slices.SortFunc(arr, func(a JobSequencing, b JobSequencing) int {
		maxDeadLineHrs = max(maxDeadLineHrs, a.DeadLine, b.DeadLine)
		if a.Profit < b.Profit {
			return 1
		}
		return -1
	})
	profit := 0

	availableSlots := make([]int, maxDeadLineHrs)

	for _, v := range arr {

		if availableSlots[v.DeadLine-1] == 0 {
			availableSlots[v.DeadLine-1] = v.Profit
		} else {
			for j := v.DeadLine - 1; j >= 0; j-- {
				if availableSlots[j] == 0 {
					availableSlots[j] = v.Profit
					break
				}
			}
		}
	}

	for i := 0; i < len(availableSlots); i++ {

		profit += availableSlots[i]
	}

	return profit
}
