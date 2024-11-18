package dynamicprogramming

// here 0_1  means that wither we consider the weight or we do not consider the weight if it overflows
func KnapSackRecursion0_1(v []int, w []int, Sum int, n int) int {

	if n == 0 || Sum == 0 {
		return 0
	}

	if w[n-1] > Sum {
		return KnapSackRecursion0_1(v, w, Sum, n-1)
	} else {
		return max(v[n-1]+KnapSackRecursion0_1(v, w, Sum-w[n-1], n-1), KnapSackRecursion0_1(v, w, Sum, n-1))
	}
}

// here 0_1  means that wither we consider the weight or we do not consider the weight if it overflows
func KnapSackDP0_1(v []int, w []int, Sum int) int {

	dp := make([][]int, len(v)+1)

	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, Sum+1)
		dp[i][0] = 0
	}

	// here we are considering J as sum & i as weight
	for i := 1; i < len(dp); i++ {
		for j := 1; j < len(dp[i]); j++ {
			if w[i-1] > j {
				dp[i][j] = dp[i-1][j]
			} else {
				dp[i][j] = max(v[i-1]+dp[i][j-w[i-1]], dp[i-1][j])
			}
		}
	}
	return dp[len(v)][Sum]
}
