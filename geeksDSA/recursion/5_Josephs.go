package recursion

func findJosephsPosition(n int, k int) int {

	if n == 1 {
		return 0
	}

	return (findJosephsPosition(n-1, k) + k) % n
}

func findJosephsPosition_Slice(n int, k int) int {

	arr := make([]int, n)

	for i := 0; i < len(arr); i++ {
		arr[i] = i
	}

	indexToCountFrom := 0
	for len(arr) != 1 {

		if indexToCountFrom > len(arr) {
			indexToCountFrom = 0
		}
		indexToDelete := indexToCountFrom + k - 1

		for indexToDelete >= len(arr) {
			indexToDelete = indexToDelete - len(arr)
		}
		arr = append(arr[:indexToDelete], arr[indexToDelete+1:]...)
		indexToCountFrom = indexToDelete
	}

	return arr[0]
}
