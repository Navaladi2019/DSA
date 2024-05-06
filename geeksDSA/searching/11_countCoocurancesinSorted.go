package searching

func CountOddurancesinSorted(arr []int, k int) int {

	star := FindFirstOccurance(arr, k)

	if star == -1 {
		return 0
	}

	end := FindLastOccurance(arr, k)

	return (end - star) + 1
}
