package arrays

func FindLeaderInArrsy(arr []int) []int {

	leaders := make([]int, 0, 10)

	maxLeader := arr[len(arr)-1]

	leaders = append(leaders, maxLeader)

	for i := len(arr) - 2; i >= 0; i-- {

		if arr[i] > maxLeader {
			maxLeader = arr[i]
			leaders = append(leaders, maxLeader)
		}
	}

	return leaders

}
