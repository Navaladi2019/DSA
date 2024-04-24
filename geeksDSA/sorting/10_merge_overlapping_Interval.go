package sorting

import (
	"sort"
)

func MergeOverLapingINterals_Naval(arr []interval) []interval {

	var result = make([]interval, 0, 10)
	sort.Slice(arr, func(i, j int) bool {
		return arr[i].start < arr[j].start
	})

	item := interval{start: arr[0].start, end: arr[0].end}
	isInterval := false

	for i := 1; i < len(arr); i++ {

		if isOverLapping(item, arr[i]) {
			item.end = max(item.end, arr[i].end)
			isInterval = true

		} else {

			if isInterval {
				result = append(result, item)
			}

			item = interval{start: arr[i].start, end: arr[i].end}

			isInterval = false
		}

	}

	if isInterval {
		item := interval{start: item.start, end: item.end}
		result = append(result, item)
	}

	return result

}

func isOverLapping(arr1 interval, arr2 interval) bool {
	if arr2.start < arr1.end {
		return true
	}

	return false
}

type interval struct {
	start int
	end   int
}
