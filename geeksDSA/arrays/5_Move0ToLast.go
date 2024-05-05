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

func MoveZeroToLast_1(arr []int) []int {

	lastNonZeroINdex := -1
	for i := 0; i < len(arr); i++ {

		if lastNonZeroINdex == -1 && arr[i] == 0 {
			lastNonZeroINdex = i
		}
		if lastNonZeroINdex > -1 && arr[i] != 0 {
			arr[lastNonZeroINdex], arr[i] = arr[i], arr[lastNonZeroINdex]
			lastNonZeroINdex++
		}

	}

	return arr
}
