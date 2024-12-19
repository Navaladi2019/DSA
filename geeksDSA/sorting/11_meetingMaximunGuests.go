package sorting

import (
	"slices"
	"sort"
)

// HERE MAKE A TABLE FOR ARRIVAL AND DEPARTURE
func FindMaximunMeetingGuestTime(arrival []int, departure []int) int {

	result := 0

	sort.Slice(arrival, func(i, j int) bool { return arrival[i] < arrival[j] })
	sort.Slice(departure, func(i, j int) bool { return departure[i] < departure[j] })

	current := 0
	i := 0
	j := 0
	for i < len(arrival) && j < len(departure) {

		if arrival[i] < departure[j] {
			current++
			i++
		} else {
			result = max(result, current)
			current--
			j++

		}
	}
	result = max(result, current)
	return result
}

func FindMaximunMeetingGuestTime_1(arrival []int, departure []int) int {

	res := 0

	slices.Sort(arrival)
	slices.Sort(departure)

	//j := 0

	for i := 0; i < len(departure); i++ {

		curr := 0
		for j := 0; j < len(arrival) && arrival[j] < departure[i]; j++ {
			curr++
		}

		res = max(res, curr-i)
	}

	return res
}

func FindMaximunMeetingGuestTime_Efficient(arrival []int, departure []int) int {

	res := 0

	slices.Sort(arrival)
	slices.Sort(departure)

	j := 0

	i := 0

	for i < len(arrival) && j < len(departure) {
		if arrival[i] < departure[j] {
			i++
		} else {
			j++
		}
		res = max(res, i-j)
	}

	return res
}
