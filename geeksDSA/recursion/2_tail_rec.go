package recursion

func RecursionGet1ToN(n int, decrement int, result *[]int) {

	if decrement < 0 {
		return
	}

	*result = append(*result, n-decrement)
	RecursionGet1ToN(n, decrement-1, result)

}

func RecursionGet1ToNOther(n int, increment int, result *[]int) {

	if n == 0 {
		return
	}

	increment++

	*result = append(*result, increment)
	RecursionGet1ToNOther(n-1, increment, result)

}

func factorialTailRecursive(n int, old int) int {

	if old == 0 {
		old = 1
	}
	if n == 0 || n == 1 {
		return old
	}

	return factorialTailRecursive(n-1, old*n)
}
