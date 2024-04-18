package searching

func findIndexBinaryArrayIterative(arr []int, n int) int {

	low := 0
	high := len(arr) - 1

	for low <= high {

		mid := (low + high) / 2
		if arr[mid] == n {
			return mid
		} else if arr[mid] > n {

			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return -1
}

func findIndexBinaryArrayRecursive(arr []int, n int, high int, low int) int {

	if low > high {
		return -1
	}

	mid := (low + high) / 2
	if arr[mid] == n {
		return mid
	} else if arr[mid] > n {

		high = mid - 1
	} else {
		low = mid + 1
	}

	return findIndexBinaryArrayRecursive(arr, n, high, low)
}

// [5,10,10,20,20,20,20,20,20,20,20,20,20,20,20,20]
func findIndexBinaryArrayFirstIndex(arr []int, n int) int {

	low := 0
	high := len(arr) - 1

	for low <= high {

		mid := (low + high) / 2
		if arr[mid] > n {
			high = mid - 1
		} else if arr[mid] < n {
			low = mid + 1
		} else {

			if mid == 0 || arr[mid-1] != n {
				return mid
			} else {
				high = mid - 1
			}
		}
	}
	return -1
}

func findIndexBinaryArrayLastIndex(arr []int, n int) int {

	low := 0
	high := len(arr) - 1

	for low <= high {

		mid := (low + high) / 2
		if arr[mid] > n {
			high = mid - 1
		} else if arr[mid] < n {
			low = mid + 1
		} else {

			if mid == (len(arr)-1) || arr[mid+1] != n {
				return mid
			} else {
				low = mid + 1
			}
		}
	}
	return -1
}
