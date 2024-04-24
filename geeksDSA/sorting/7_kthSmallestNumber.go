package sorting

func Find_Kth_SmallestNumber(arr []int, k int) int {

	low, high := 0, len(arr)-1

	for low <= high {
		p := Lomuto_partition(arr, low, high)

		// here p gives the index so we should add it witn p+1 == k
		if p+1 == k {
			return arr[p]
		} else if p+1 < k {

			low = p + 1
		} else {
			high = p - 1
		}
	}
	return -1
}
