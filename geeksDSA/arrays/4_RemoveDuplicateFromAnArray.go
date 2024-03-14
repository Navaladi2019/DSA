package arrays

/*Remove duplicates and align the unique values left and return size from a sorted array*/
func RemoveDuplicateFromAnArray(arr []int) int {

	lastUniqueIndex := 0

	for i := 1; i < len(arr); i++ {
		if arr[i] != arr[lastUniqueIndex] {
			arr[lastUniqueIndex+1] = arr[i]
			lastUniqueIndex++
		}
	}

	return lastUniqueIndex + 1
}
