package searching

func findSquareRoor(n int) int {

	low := 0
	high := n

	for low <= high {

		mid := (low + high) / 2

		if mid*mid == n {
			return mid
		} else if mid*mid < n {
			low = mid + 1
		} else {

			if (mid-1)*(mid-1) < n && mid*mid > n {

				return mid - 1
			} else {
				high = mid - 1
			}

		}
	}

	return 0
}
