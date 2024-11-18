package dynamicprogramming

// here for i & j we are p[asing length and not the index
func EditDistanceToGetSameString(s1 string, s2 string, i int, j int) int {

	if i == 0 {
		return j
	}

	if j == 0 {
		return i
	}

	if s1[i-1] == s2[j-1] {
		return EditDistanceToGetSameString(s1, s2, i-1, j-1)
	} else {
		return 1 + min(
			//Replace
			EditDistanceToGetSameString(s1, s2, i-1, j-1),
			//Delete
			EditDistanceToGetSameString(s1, s2, i-1, j),
			//Insert
			EditDistanceToGetSameString(s1, s2, i, j-1),
		)
	}
}

// here for i & j we are p[asing length and not the index
func EditDistanceToGetSameString_DP(s1 string, s2 string) int {

	dp := make([][]int, len(s1)+1)

	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, len(s2)+1)
		dp[i][0] = i
	}
	for i := 0; i < len(dp[0]); i++ {
		dp[0][i] = i
	}

	for i := 1; i < len(dp); i++ {
		for j := 1; j < len(dp[i]); j++ {

			if s1[i-1] == s2[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				dp[i][j] = 1 + min(dp[i-1][j-1], dp[i][j-1], dp[i-1][j])
			}
		}
	}

	return dp[len(s1)][len(s2)]

}
