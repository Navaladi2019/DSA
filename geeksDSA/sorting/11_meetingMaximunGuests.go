package sorting

import (
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
