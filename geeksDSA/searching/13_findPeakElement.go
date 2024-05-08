package searching

func FindPeakElement(arr []int) int {

	low := 0
	high := len(arr) - 1

	for low <= high {

		mid := (low + high) / 2

		if (mid == 0 || arr[mid] > arr[mid-1]) && (mid == len(arr)-1 || arr[mid] > arr[mid+1]) {
			return arr[mid]
		}

		if mid > 0 && arr[mid-1] >= arr[mid] {

			high = mid - 1
		} else {
			low = mid + 1
		}

	}

	return -1
}
