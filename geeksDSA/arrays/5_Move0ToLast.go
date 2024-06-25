package arrays

func MoveZeroToLast(arr []int) []int {

	lastNonZeroINdex := -1
	for i := 0; i < len(arr); i++ {

		if arr[i] != 0 {
			arr[i], arr[lastNonZeroINdex+1] = arr[lastNonZeroINdex+1], arr[i]
			lastNonZeroINdex++
		}
	}

	return arr
}

func MoveZeroToLast_1(arr []int) []int {

	lastZeroINdex := -1
	for i := 0; i < len(arr); i++ {

		if lastZeroINdex == -1 && arr[i] == 0 {
			lastZeroINdex = i
		}
		if lastZeroINdex > -1 && arr[i] != 0 {
			arr[lastZeroINdex], arr[i] = arr[i], arr[lastZeroINdex]
			lastZeroINdex++
		}

	}

	return arr
}
