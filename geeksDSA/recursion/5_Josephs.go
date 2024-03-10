package recursion

func findJosephsPosition(n int, k int) int {

	if n == 1 {
		return 0
	}

	return findJosephsPosition(n-1, k)
}
