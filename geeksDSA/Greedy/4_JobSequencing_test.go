package greedy

import "testing"

func Test_GetMaximumProfitOnJobSequencing(t *testing.T) {

	arr := make([]JobSequencing, 5)
	arr[0] = JobSequencing{DeadLine: 4, Profit: 50}
	arr[1] = JobSequencing{DeadLine: 1, Profit: 5}
	arr[2] = JobSequencing{DeadLine: 1, Profit: 20}
	arr[3] = JobSequencing{DeadLine: 5, Profit: 10}
	arr[4] = JobSequencing{DeadLine: 5, Profit: 80}

	got := GetMaximumProfitOnJobSequencing(arr)

	if got != 160 {
		t.Error("has error in job sequencxe expected value of 120")
	}
}
