package searching

func SearchInInfiniteArray(arr []int, n int) int {

	low, high := 0, 2

	for true {

		mid := (low + high) / 2

		if arr[mid] < n {

			low = mid + 1
			high *= 2
		} else {
			return findIndexBinaryArrayFirstIndex1(arr, n, low, high)
		}
	}

	return -1
}

func findIndexBinaryArrayFirstIndex1(arr []int, n int, low int, high int) int {

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

func findPeakElement(arr []int) int {

	if len(arr) == 1 {
		return arr[0]
	}

	if len(arr) > 1 && arr[0] >= arr[1] {
		return arr[0]
	}

	if len(arr) > 1 && arr[len(arr)-1] >= arr[len(arr)-2] {
		return arr[len(arr)-1]
	}

	for i := 1; i < len(arr)-1; i++ {

		if arr[i] >= arr[i-1] && arr[i] >= arr[i+1] {
			return arr[i]
		}
	}

	return -1
}

func findPeakElementBinarySearch(arr []int) int {

	low := 0
	high := len(arr) - 1

	for low <= high {

		mid := (low + high) / 2

		if (mid == 0 || arr[mid-1] <= arr[mid]) && (mid == len(arr)-1 || arr[mid-1] <= arr[mid]) {
			return mid
		}

		if mid > 0 && arr[mid-1] <= arr[mid] {

			low = mid + 1
		} else {
			high = mid - 1
		}

	}

	return -1
}
