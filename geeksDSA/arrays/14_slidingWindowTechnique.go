package arrays

func GetMaxSumWindowSlidingLogic(arr []int, k int) int {

	WindowValue := 0

	for i := 0; i < k; i++ {
		WindowValue += arr[i]
	}

	res := WindowValue

	for i := k; i < len(arr); i++ {

		WindowValue = WindowValue - arr[i-k] + arr[i]
		res = max(WindowValue, res)

	}

	return res
}
