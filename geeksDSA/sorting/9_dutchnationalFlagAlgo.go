package sorting

// ITS A VARIATION OF HOARES ALGORITHM
//[0,1,2,0,1,1,2]
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

func sortItemsInArray_1(arr []int) {

	low, mid, high := 0, 0, len(arr)-1

	for mid <= high {

		if arr[mid] == 0 {
			arr[low], arr[mid] = arr[mid], arr[low]
			low++
			mid++
		} else if arr[mid] == 1 {
			mid++
		} else {
			arr[mid], arr[high] = arr[high], arr[mid]
			high--
		}
	}

}
