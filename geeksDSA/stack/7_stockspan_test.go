package stack

import "testing"

func Test_FindStockSpan(t *testing.T) {

	got := FindStockSpan([]int{60, 10, 20, 40, 35, 30, 50, 70, 65}, 0)

	if got != 1 {

		t.Error("haserror")
	}

	got = FindStockSpan([]int{60, 10, 20, 40, 35, 30, 50, 70, 65}, 1)

	if got != 1 {

		t.Error("haserror")
	}

	got = FindStockSpan([]int{60, 10, 20, 40, 35, 30, 50, 70, 65}, 2)

	if got != 2 {

		t.Error("haserror")
	}

	got = FindStockSpan([]int{60, 10, 20, 40, 35, 30, 50, 70, 65}, 3)

	if got != 3 {

		t.Error("haserror")
	}
	got = FindStockSpan([]int{60, 10, 20, 40, 35, 30, 50, 70, 65}, 4)

	if got != 1 {

		t.Error("haserror")
	}

	got = FindStockSpan([]int{60, 10, 20, 40, 35, 30, 50, 70, 65}, 5)

	if got != 1 {

		t.Error("haserror")
	}

	got = FindStockSpan([]int{60, 10, 20, 40, 35, 30, 50, 70, 65}, 6)

	if got != 6 {

		t.Error("haserror")
	}

	got = FindStockSpan([]int{60, 10, 20, 40, 35, 30, 50, 70, 65}, 7)

	if got != 8 {

		t.Error("haserror")
	}

	got = FindStockSpan([]int{60, 10, 20, 40, 35, 30, 50, 70, 65}, 8)

	if got != 1 {

		t.Error("haserror")
	}
}
