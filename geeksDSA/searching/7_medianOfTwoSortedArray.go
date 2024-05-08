package searching

func FindMedianOfTwoSortedArray(arr1 []int, arr2 []int) int {

	finalLength := len(arr1) + len(arr2)

	j := 0
	k := 0

	isOdd := finalLength%2 == 1
	arr := make([]int, 0, finalLength)

	loopLenth := (finalLength / 2) + 1

	for i := 0; i < loopLenth; i++ {

		if k == len(arr2) || arr1[j] < arr2[k] {
			arr = append(arr, arr1[j])
			j++

		} else {
			arr = append(arr, arr2[k])
			k++
		}
	}

	if isOdd {
		return arr[len(arr)-1]
	}

	return (arr[len(arr)-1] + arr[len(arr)-2]) / 2
}
