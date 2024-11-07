package dynamicprogramming

func CoinChangeNaive(coins []int, n int, sum int) int {

	if sum == 0 {
		return 1
	}
	if sum < 0 || n < 0 {
		return 0
	}
	return CoinChangeNaive(coins, n, sum-coins[n]) + CoinChangeNaive(coins, n-1, sum)
}

func CoinChange_DP(coins []int, sum int) int {

	dp := make([][]int, len(coins)+1)

	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, sum+1)
		dp[i][0] = 1
	}

	for i := 1; i < len(dp); i++ {

		for j := 1; j < len(dp[i]); j++ {
			dp[i][j] = dp[i-1][j]

			if coins[i-1] <= j {
				dp[i][j] += dp[i][j-coins[i-1]]
			}
		}
	}

	return dp[len(coins)][sum]
}
