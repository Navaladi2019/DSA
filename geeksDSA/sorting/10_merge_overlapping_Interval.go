package sorting

import (
	"slices"
)

func MergeOverLapingINterals_Naval(arr []interval) []interval {

	var result = make([]interval, 0, 10)
	// sort.Slice(arr, func(i, j int) bool {
	// 	return arr[i].start < arr[j].start
	// })

	slices.SortFunc(arr, func(a interval, b interval) int {
		if a.start < b.start {
			return -1
		} else if a.start > b.start {
			return 1
		}
		return 0
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

func MergeOverlappingINtervals(arr []interval) []interval {

	slices.SortFunc(arr, func(a interval, b interval) int {

		if a.start < b.start {
			return -1
		} else if a.start == b.start {
			return 0
		}

		return 1
	})
	res := make([]interval, 0, 1)

	curr := interval{start: arr[0].start, end: arr[0].end}

	islastappended := false
	for i := 1; i < len(arr); i++ {

		if isOverLapping(curr, arr[i]) {
			curr.end = max(curr.end, arr[i].end)
		} else {
			res = append(res, curr)

			curr = interval{start: arr[i].start, end: arr[i].end}

			if i == len(arr)-1 {
				islastappended = true
			}
		}
	}

	if !islastappended {
		res = append(res, curr)
	}

	return res
}

func isOverLapping(arr1 interval, arr2 interval) bool {
	if arr2.start <= arr1.end {
		return true
	}

	return false
}

type interval struct {
	start int
	end   int
}
