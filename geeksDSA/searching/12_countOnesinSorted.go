package searching

func CountOnesInSortedArray(arr []int, k int) int {

	star := FindFirstOccurance(arr, k)

	if star == -1 {
		return 0
	}

	return (len(arr) - star)
}
