package dynamicprogramming

func rob(nums []int) int {
	return dp(len(nums)-1, nums)
}

func dp(i int, nums []int) int {

	if i == 0 || i == 1 {
		return nums[i]
	}

	if val, ok := m[i]; ok {
		return val
	}
	val2 := dp(i-2, nums) + nums[i]
	vals := max(dp(i-1, nums), val2)

	m[i] = vals

	return vals
}

var m map[int]int = make(map[int]int)
