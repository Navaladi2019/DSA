package stack

import "testing"

func Test_FindPreviousGreater(t *testing.T) {
	got := FindPreviousGreater([]int{15, 10, 18, 12, 4, 6, 2, 8}, 0)

	if got != -1 {

		t.Error("haserror")
	}

	got = FindPreviousGreater([]int{15, 10, 18, 12, 4, 6, 2, 8}, 1)

	if got != 15 {

		t.Error("haserror")
	}

	got = FindPreviousGreater([]int{15, 10, 18, 12, 4, 6, 2, 8}, 2)

	if got != -1 {

		t.Error("haserror")
	}

	got = FindPreviousGreater([]int{15, 10, 18, 12, 4, 6, 2, 8}, 3)

	if got != 18 {

		t.Error("haserror")
	}

	got = FindPreviousGreater([]int{15, 10, 18, 12, 4, 6, 2, 8}, 4)

	if got != 12 {

		t.Error("haserror")
	}

	got = FindPreviousGreater([]int{15, 10, 18, 12, 4, 6, 2, 8}, 5)

	if got != 12 {

		t.Error("haserror")
	}

	got = FindPreviousGreater([]int{15, 10, 18, 12, 4, 6, 2, 8}, 6)

	if got != 6 {

		t.Error("haserror")
	}

	got = FindPreviousGreater([]int{15, 10, 18, 12, 4, 6, 2, 8}, 7)

	if got != 12 {

		t.Error("haserror")
	}
}
