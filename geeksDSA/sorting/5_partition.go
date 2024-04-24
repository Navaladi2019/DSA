package sorting

/* in lomuto partition we consider the h to be last eleent otherwise we just swap it to be last*/
func Lomuto_partition(arr []int, low int, high int) int {

	pivot := arr[high]
	lastHighIndex := low

	for i := low; i <= high-1; i++ {

		if arr[i] < pivot {

			temp := arr[i]
			arr[i] = arr[lastHighIndex]
			arr[lastHighIndex] = temp
			lastHighIndex++
		}

	}

	arr[high], arr[lastHighIndex] = arr[lastHighIndex], arr[high]

	return lastHighIndex

}

/* In Hoares Partition we consider the pivot element to be on first.
Hoares Partition does not put pivot element in the correct place*/
func Hoares_Partition(arr []int, low int, high int) int {

	pivot := arr[low]

	i := low - 1
	j := high + 1

	for {

		// below is do whle loop in golang
		for next := true; next; next = arr[i] < pivot {
			i++
		}

		for next := true; next; next = arr[j] > pivot {
			j--
		}

		if i >= j {
			return j
		}

		arr[i], arr[j] = arr[j], arr[i]

	}

}
