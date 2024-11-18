package dynamicprogramming

import "math"

func GetMinCoins(arr []int, sum int) int {

	if sum == 0 {
		return 0
	}

	if sum < 0 {
		return math.MaxInt
	}
	res := math.MaxInt

	for i := 0; i < len(arr); i++ {
		val := GetMinCoins(arr, sum-arr[i])
		if val < math.MaxInt {
			res = min(res, val+1)
		}
	}
	return res
}

func GetMinCoinsDP(Coins []int, sum int) int {

	dp := make([]int, sum+1)

	for i := 1; i < len(dp); i++ {
		val := math.MaxInt

		for _, c := range Coins {
			if c <= i {
				val = min(val, dp[i-c]+1)
			}
		}

		dp[i] = val
	}
	return dp[sum]

}
