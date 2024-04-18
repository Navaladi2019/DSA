package searching

/*
JHere we are going to use two pair in combination so it would become on^2 instead of o n^3*/
func FindTripletInSortedArray(arr []int, sum int) bool {

	for l := 0; l < len(arr)-2; l++ {

		pairSum := sum - arr[l]

		i := l
		j := len(arr) - 1

		// here finding other two number by inding pairs
		for i != j {

			for i != j {

				if arr[i]+arr[j] == pairSum {
					return true
				}

				if arr[i]+arr[j] > pairSum {
					j--
				} else {
					i++
				}
			}

		}

	}

	return false
}
