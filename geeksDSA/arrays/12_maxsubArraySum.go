package arrays

func getSubArrays(arr []int, index int, result [][]int) [][]int {

	if index == len(arr) {
		return result
	}

	subarry := [][]int{{arr[index]}}

	for _, s := range result {
		s = append(s, arr[index])
		subarry = append(subarry, [][]int{s}...)
	}

	result = append(result, subarry...)

	return getSubArrays(arr, index+1, result)
}

func getSubArraysForLoop(arr []int) [][]int {

	result := make([][]int, 0, 10)

	for i := 0; i < len(arr); i++ {

		subsets := []int{arr[i]}

		for _, s := range result {
			result = append(result, append(s, arr[i]))
		}
		result = append(result, subsets)
	}

	return result
}

/// sub array of [1,2,3] => [1}, [2],[3],[1,2],[2,3],[1,2,3] and not [1,3] because 1 &3 are not consecutive

func findMaximunSubarrayNavive(arr []int) int {

	maxi := arr[0]

	for i := 1; i < len(arr); i++ {
		current := 0

		for j := i; j < len(arr); j++ {
			current += arr[i]
			maxi = max(maxi, current)
		}
	}

	return maxi
}

// max sum of sub array formula max[i] = max(max(i-1)+arr[i],arr[i])

func getMaxSumOfSubArray(arr []int) int {

	maxSum := arr[0]
	lastMax := arr[0]

	for i := 1; i < len(arr); i++ {
		lastMax = max(lastMax+arr[i], arr[i])
		maxSum = max(maxSum, lastMax)
	}

	return maxSum
}
