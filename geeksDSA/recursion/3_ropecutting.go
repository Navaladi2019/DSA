package recursion

/*given a rope of length n and three number a,b,c.
cut the rope into maximun number of pieces using a,b,c , provided 0 < a,b,c <=n*/
func find_MaximunCutsInRope(n int, a int, b int, c int) int {

	if n < 0 {
		return -1
	}
	if n == 0 {
		return 0
	}

	res := max(find_MaximunCutsInRope(n-a, a, b, c), find_MaximunCutsInRope(n-b, a, b, c), find_MaximunCutsInRope(n-c, a, b, c))

	if res == -1 {
		return -1
	}
	return res + 1
}
