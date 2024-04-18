package sorting

/* in lomuto partition we consider the h to be last eleent otherwise we just swap it to be last*/
func Lomuto_partition(arr []int, low int, high int) {

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

	arr[high] = arr[lastHighIndex]
	arr[lastHighIndex] = pivot

}

/* In Hoares Partition we consider the pivot element to be on first.
Hoares Partition does not put pivot element in the correct place*/
func Hoares_Partition(arr []int, low int, high int) {

	pivot := arr[low]

	i := low
	j := high

	for {

		for arr[i] < pivot {
			i++
		}

		for arr[j] > pivot {
			j++
		}

		if i >= j {
			return
		}

		//swap
		temp := arr[i]
		arr[i] = arr[j]
		arr[j] = temp

	}

}
