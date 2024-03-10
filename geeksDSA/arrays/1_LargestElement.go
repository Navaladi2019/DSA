package arrays

func findLargestElementInArray(arr []int) int {
	max := arr[0]
	for i := 0; i < len(arr)/2; i++ {
		if arr[i] > max {
			max = arr[i]
		}

		if arr[len(arr)-1-i] > max {
			max = arr[len(arr)-1-i]
		}
	}
	return max
}

func FindSecondLargestElementIndex(arr []int) int {
	maxIndex := 0
	secondindex := -1
	for i := 0; i < len(arr); i++ {
		if arr[i] > arr[maxIndex] {
			secondindex = maxIndex
			maxIndex = i
		} else if secondindex != -1 && arr[i] > arr[secondindex] && arr[i] != arr[maxIndex] {
			secondindex = i
		}
	}
	return secondindex
}
