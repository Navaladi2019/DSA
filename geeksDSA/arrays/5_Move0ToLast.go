package arrays

func MoveZeroToLast(arr []int) []int {

	lastNonZeroINdex := -1
	for i := 0; i < len(arr); i++ {

		if arr[i] != 0 {
			tmp := arr[i]
			arr[i] = arr[lastNonZeroINdex+1]
			arr[lastNonZeroINdex+1] = tmp
			lastNonZeroINdex++
		}
	}

	return arr
}
