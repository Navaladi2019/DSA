package arrays

func FindLonogestEvenOddSubArray(arr []int) int {

	res := 1
	current := 1
	for i := 1; i < len(arr); i++ {

		if arr[i]%2 != arr[i-1]%2 {
			current++
		} else {
			current = 1
		}
		res = max(current, res)
	}

	return res
}
