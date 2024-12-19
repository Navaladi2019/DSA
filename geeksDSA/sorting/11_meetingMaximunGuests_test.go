package sorting

import (
	"testing"
)

func Test_FindMaximunMeetingGuestTime(t *testing.T) {

	arrival := []int{900, 940, 950, 1100, 1500, 1800}
	dep := []int{910, 1200, 1120, 1130, 1900, 2000}

	got := FindMaximunMeetingGuestTime_Efficient(arrival, dep)

	if got != 3 {
		t.Error("Has error in maximun guest meet")
	}
}
