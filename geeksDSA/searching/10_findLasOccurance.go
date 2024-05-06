package searching

func FindLastOccurance(arr []int, k int) int {
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
			if mid != len(arr)-1 && arr[mid+1] == arr[mid] {
				low = mid + 1
			} else {
				return mid
			}
		}
	}

	return index
}
