package searching

/*
Naive solution would be o n2 without two pointer approach
*/
func FindPairsEqualToSum(arr []int, sum int) bool {

	i := 0
	j := len(arr) - 1

	for i != j {

		if arr[i]+arr[j] == sum {
			return true
		}

		if arr[i]+arr[j] > sum {
			j--
		} else {
			i++
		}
	}

	return false
}
