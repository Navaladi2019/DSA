package arrays

func findIsArraySorted(ar []int) bool {

	var isAscending bool
	var isDescending bool

	for i := 1; i < len(ar); i++ {

		if isAscending == false && isDescending == false {

			if ar[i-1] < ar[i] {
				isAscending = true
			}

			if ar[i-1] > ar[i] {
				isDescending = true
			}
		}

		if isAscending && ar[i-1] > ar[i] {
			return false
		}

		if isDescending && ar[i-1] < ar[i] {
			return false
		}
	}

	return true
}
