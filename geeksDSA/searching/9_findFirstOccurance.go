package searching

func FindFirstOccurance(arr []int, k int) int {
	index := -1

	low := 0
	high := len(arr) - 1

	for low <= high {

		mid := (low + high) / 2

		if arr[mid] > k {

			high = mid - 1
		} else if arr[mid] < k {
			low = mid + 1
		} else {
			if mid != 0 && arr[mid-1] == arr[mid] {
				high = mid - 1
			} else {
				return mid

			}
		}

	}

	return index
}
