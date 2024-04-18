package arrays

func isSubArrayWithGivenSum(arr []int, sum int) bool {

	startingIndex := 0
	Cursum := 0
	for i := 0; i < len(arr); i++ {
		Cursum = Cursum + arr[i]
		for Cursum > sum {
			Cursum -= arr[startingIndex]
			startingIndex++
		}
		if Cursum == sum {
			return true
		}
	}

	return false
}
