package sorting

/* in lomuto partition we consider the h to be last element otherwise we just swap it to be last*/
func Lomuto_partition(arr []int, low int, high int) int {

	pivot := arr[high]
	partitionIndex := low

	for i := low; i <= high-1; i++ {
		if arr[i] < pivot {
			arr[i], arr[partitionIndex] = arr[partitionIndex], arr[i]
			partitionIndex++
		}
	}

	arr[high], arr[partitionIndex] = arr[partitionIndex], arr[high]

	return partitionIndex

}

/* In Hoares Partition we consider the pivot element to be on first.
Hoares Partition does not put pivot element in the correct place

Here we use do while loop because after the first swap when we return partition key we need to return next index
and not the index we swapped last

*/
func Hoares_Partition(arr []int, low int, high int) int {

	if len(arr) == 0 {
		return -1
	}
	pivot := arr[low]

	i := low - 1
	j := high + 1

	for {

		// below is do while loop in golang
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

func Hoares_Partition_1(arr []int, low int, high int) int {

	if len(arr) == 0 {
		return -1
	}
	pivot := arr[low]

	i := low - 1
	j := high + 1

	for {

		// below is do while loop in golang
		i++
		for arr[i] < pivot {
			i++
		}

		j--
		for arr[j] > pivot {
			j--
		}

		if i >= j {
			return j
		}

		arr[i], arr[j] = arr[j], arr[i]

	}
}
