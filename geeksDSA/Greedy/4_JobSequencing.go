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

func GetMaximumProfitOnJobSequencing_Efficient(arr []JobSequencing) int {

	res := 0

	// get max deadline to get the size of result array
	maxDeadLine := 0

	slices.SortFunc(arr, func(a JobSequencing, b JobSequencing) int {
		maxDeadLine = max(maxDeadLine, a.DeadLine, b.DeadLine)
		if a.Profit < b.Profit {
			return 1
		}
		return -1
	})

	availableSlots := make([]int, maxDeadLine)

	for i := 0; i < len(arr); i++ {

		if (availableSlots[arr[i].DeadLine-1]) == 0 {
			availableSlots[arr[i].DeadLine-1] = arr[i].Profit
		} else {
			for j := arr[i].DeadLine - 2; j >= 0; j-- {
				if availableSlots[j] == 0 {
					availableSlots[j] = arr[i].Profit
					break
				}
			}
		}
	}

	for _, val := range availableSlots {
		res += val
	}

	return res
}
