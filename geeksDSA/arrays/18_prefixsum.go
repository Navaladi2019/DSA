package arrays

func PrefixSumPreProcess(arr []int) {

	for i := 1; i < len(arr); i++ {

		arr[i] = arr[i] + arr[i-1]
	}
}

func FindPrefixSum(arr []int, start int, end int) int {

	if start == 0 {
		return arr[end]
	}
	return arr[end] - arr[start-1]
}
