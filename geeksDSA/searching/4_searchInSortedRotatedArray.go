package searching

//[10,20,50,60,5,8]
//[10, 20, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16]
func FindInSortedRoratedArray(arr []int, n int) int {

	low := 0
	high := len(arr) - 1

	for low <= high {
		mid := (low + high) / 2

		if arr[mid] == n {
			return mid
		}

		if arr[low] <= arr[mid] {

			//left half is sorted
			if arr[low] <= n && arr[mid] > n {

				high = mid - 1
			} else {
				low = mid + 1
			}

		} else {

			if arr[mid] <= arr[high] && arr[mid] < n {

				// right half is sorted
				low = mid + 1
			} else {

				high = mid - 1
			}
		}

	}

	return -1
}
