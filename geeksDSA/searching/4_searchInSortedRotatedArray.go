package searching

//[10,20,50,60,5,8]
//[10, 20, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16]
// Here there is only two possibility either right half is sorted or left half is sorted
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
			// right half is sorted
			if arr[mid] <= arr[high] && arr[mid] < n {

				low = mid + 1
			} else {

				high = mid - 1
			}
		}

	}

	return -1
}
func FindInSortedRoratedArray_1(arr []int, n int) int {

	low := 0
	high := len(arr) - 1

	for low <= high {
		mid := (low + high) / 2
		if arr[mid] == n {
			return mid
		}

		if arr[low] <= arr[mid] {
			// left side is sorted
			if arr[low] <= n && arr[mid] > n {
				high = mid - 1
			} else {
				low = mid + 1
			}
		} else {
			// right side is sorted
			if arr[high] >= n && arr[mid] < n {
				low = mid + 1
			} else {
				high = mid - 1
			}
		}
	}

	return -1
}

func FindInSortedRotatedArray_2(arr []int, n int) int {

	low := 0
	high := len(arr) - 1

	for low <= high {

		mid := (low + high) / 2

		if (arr[mid]) == n {
			return mid
		} else {
			//leftsorted
			if arr[low] <= arr[mid] {
				if arr[low] <= n && arr[mid] > n {
					high = mid - 1
				} else {
					low = mid + 1
				}
			} else {
				// right is sorted

				if arr[mid] < n && arr[high] >= n {
					low = mid + 1
				} else {
					high = mid - 1
				}

			}
		}
	}

	return -1
}

func FindInSortedRotatedArray(arr []int, n int) int {

	l := 0
	h := len(arr) - 1

	for l <= h {

		mid := (l + h) / 2

		if arr[mid] == n {
			return mid
		}

		if arr[l] <= arr[mid] {
			//left half is sorted
			// Now check if n is with in the range of low and mid
			// if it is within the range then do binary with in that

			if n >= arr[l] && n < arr[mid] {
				h = mid - 1
			} else {
				l = mid + 1
			}

		} else {
			//Right half is sorted
			// Now check if n is with in the range of   mid and high
			// if it is within the range then do binary with in that
			if n > arr[mid] && n <= arr[h] {
				l = mid + 1
			} else {
				h = mid - 1
			}

		}
	}

	return -1
}
