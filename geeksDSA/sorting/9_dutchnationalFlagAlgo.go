package sorting

func SortThreeItemsInArray(arr []int) {

	low, mid, high := 0, 0, len(arr)-1

	for mid <= high {

		if arr[mid] == 0 {
			arr[low], arr[mid] = arr[mid], arr[low]
			low++
			mid++
		} else if arr[mid] == 1 {
			mid++
		} else {
			arr[high], arr[mid] = arr[mid], arr[high]
			high--
		}
	}
}
