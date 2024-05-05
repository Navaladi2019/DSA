package arrays

func findIsArraySorted(ar []int) bool {

	var isAscending bool
	var isDescending bool

	if len(ar) == 1 {
		return true
	}

	if ar[0] < ar[1] {
		isAscending = true
	}

	if ar[0] > ar[1] {
		isDescending = true
	}

	for i := 1; i < len(ar); i++ {

		if isAscending && ar[i-1] > ar[i] {
			return false
		}

		if isDescending && ar[i-1] < ar[i] {
			return false
		}
	}

	return true
}
