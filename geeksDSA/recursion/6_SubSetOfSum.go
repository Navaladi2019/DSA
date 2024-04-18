package recursion

func FindSubSetofSum(arr []int, cur []int, curindex int, targetSum int) int {

	if curindex == len(arr) {

		sum := 0

		for i := 0; i < len(cur); i++ {
			sum += cur[i]
		}
		if sum == targetSum {
			return 1
		} else {

			return 0
		}
	}

	curcopy := make([]int, len(cur))

	copy(curcopy, cur)

	res := FindSubSetofSum(arr, append(curcopy, arr[curindex]), curindex+1, targetSum) + FindSubSetofSum(arr, curcopy, curindex+1, targetSum)
	return res
}

/* here instead of creating */
func FindSubSetProblemOptimized(arr []int, curIndex int, targetSum int) int {
	if curIndex == len(arr) {
		if targetSum == 0 {
			return 1
		} else {
			return 0
		}
	}

	res := FindSubSetProblemOptimized(arr, curIndex+1, targetSum-arr[curIndex]) + FindSubSetProblemOptimized(arr, curIndex+1, targetSum)
	return res
}
