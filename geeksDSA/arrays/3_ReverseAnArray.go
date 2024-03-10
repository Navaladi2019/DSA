package arrays

func ReverseAnArray(arr []int) []int {

	for i := 0; i < len(arr)/2; i++ {
		temp := arr[i]
		arr[i] = arr[len(arr)-1-i]
		arr[len(arr)-1-i] = temp
	}

	return arr
}
