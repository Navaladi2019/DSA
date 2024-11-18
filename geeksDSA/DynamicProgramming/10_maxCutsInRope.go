package dynamicprogramming

func FindMaxCutsInRope_Recursion(arr []int, val int) int {
	if val == 0 {
		return 0
	}

	if val < 0 {
		return -1
	}
	res := 0
	for i := 0; i < len(arr); i++ {
		val := FindMaxCutsInRope_Recursion(arr, val-arr[i])
		if val > -1 {
			res = max(res, val+1)
		}
	}
	return res
}
