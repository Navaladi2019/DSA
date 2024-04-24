package sorting

func CountingSort(arr []int, maxElement int) []int {

	op := make([]int, len(arr))

	countArr := make([]int, maxElement+1)

	for i := 0; i < len(arr); i++ {

		countArr[arr[i]]++
	}

	for i := 1; i < len(countArr); i++ {
		countArr[i] = countArr[i] + countArr[i-1]
	}

	// here we are go from last of array so that we have the stable sorting
	for i := len(arr) - 1; i >= 0; i-- {

		index := countArr[arr[i]] - 1
		countArr[arr[i]]--
		op[index] = arr[i]
	}

	return op
}
