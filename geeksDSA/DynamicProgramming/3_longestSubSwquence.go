package dynamicprogramming

func LCS(s1 string, s2 string, i int, j int) int {

	if i == -1 || j == -1 {
		return 0
	}

	if s1[i] == s2[j] {
		return 1 + LCS(s1, s2, i-1, j-1)
	} else {
		return max(LCS(s1, s2, i-1, j), LCS(s1, s2, i, j-1))
	}
}

func LCS_Memo(s1 string, s2 string, i int, j int, memo [][]int) int {

	if i == -1 || j == -1 {
		return 0
	} else if memo[i][j] != -1 {
		return memo[i][j]
	} else if s1[i] == s2[j] {
		memo[i][j] = 1 + LCS_Memo(s1, s2, i-1, j-1, memo)
	} else {
		memo[i][j] = max(LCS_Memo(s1, s2, i-1, j, memo), LCS_Memo(s1, s2, i, j-1, memo))
	}
	return memo[i][j]
}

func FindLCSUsinMemoization(s1 string, s2 string) int {
	memo := make([][]int, len(s1))

	for i := 0; i < len(memo); i++ {
		memo[i] = make([]int, len(s2))
		for index := range memo[i] {
			memo[i][index] = -1
		}
	}

	return LCS_Memo(s1, s2, len(s1)-1, len(s2)-1, memo)
}

func FindLCSUsinTabulation(s1 string, s2 string) int {
	table := make([][]int, len(s1)+1)

	for i := 0; i < len(table); i++ {
		table[i] = make([]int, len(s2)+1)
	}

	for i := 1; i < len(table); i++ {
		for j := 1; j < len(table[i]); j++ {
			if s1[i-1] == s2[j-1] {
				table[i][j] = 1 + table[i-1][j-1]
			} else {
				table[i][j] = max(table[i-1][j], table[i][j-1])
			}
		}
	}

	return table[len(s1)][len(s2)]
}
