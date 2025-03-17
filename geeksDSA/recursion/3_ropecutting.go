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

func find_MaximunCutsInRope_1(n int, a int, b int, c int) int {

	if n == 0 {
		return 0
	}

	if n < 0 {
		return -1
	}

	aVal := find_MaximunCutsInRope_1(n-a, a, b, c)

	bVal := find_MaximunCutsInRope_1(n-b, a, b, c)

	cVal := find_MaximunCutsInRope_1(n-c, a, b, c)

	res := max(aVal, bVal, cVal)

	if res != -1 {
		res += 1
	}

	return res

}
