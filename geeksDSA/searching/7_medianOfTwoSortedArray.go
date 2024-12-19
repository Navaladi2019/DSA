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

func FindMedianOfTwoSortedArray_Efficient(arr1 []int, arr2 []int) int {

	if len(arr1) > len(arr2) {
		arr1, arr2 = arr2, arr1
	}

	low := 0
	high := len(arr1)

	isOdd := (len(arr1)+len(arr2))%2 == 1

	for {

		x := (low + high) / 2
		y := ((len(arr1) + len(arr2)) / 2) - x

		if arr2[y] <= arr1[x+1] && arr1[x] <= arr2[y+1] {
			if isOdd {
				return min(arr1[x+1], arr2[y+1])
			}

			return (max(arr1[x], arr2[y]) + min(arr1[x+1], arr2[y+1])) / 2
		}

		if arr1[x+1] < arr2[y] {
			low = x + 1
		} else {
			high = x
		}

	}

	return 0

}
