package sorting

/*
has o(n^2) and its a stable sort
*/
func BubbleSort(arr []int) {

	swapped := false

	for i := 0; i < len(arr); i++ {

		for j := 0; j < len(arr)-1-i; j++ {

			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
				swapped = true
			}
		}

		if swapped == false {
			break
		}
	}
}

/*
has o(n^2) and its  not a stable sort but it has very less swap.
Selection Sort works by putting small elements at i ( from left to right)
*/
func SelectionSort(arr []int) {

	for i := 0; i < len(arr); i++ {

		minindex := i
		for j := i + 1; j < len(arr); j++ {
			if arr[j] < arr[minindex] {
				minindex = j
			}
		}
		arr[i], arr[minindex] = arr[minindex], arr[i]
	}
}

/* Insertion sort works better with small items
Here we consider that left side is aleady sorted
its a stable sorting algorithm
*/

func InsertionSort(arr []int) {
	for i := 1; i < len(arr); i++ {

		temp := arr[i]

		j := i - 1

		for j >= 0 && arr[j] > temp {
			arr[j+1] = arr[j]
			j--
		}

		arr[j+1] = temp

	}
}

func MergerTwoSortedSort(arr1 []int, arr2 []int) []int {
	result := make([]int, 0, len(arr1)+len(arr2))

	i := 0
	j := 0

	for i < len(arr1) && j < len(arr2) {
		if arr1[i] <= arr2[j] {
			result = append(result, arr1[i])
			i++
		} else {
			result = append(result, arr2[j])
			j++
		}

	}

	if i == len(arr1) && j < len(arr2) {
		result = append(result, arr2[j:]...)
	}

	if j == len(arr2) && i < len(arr1) {
		result = append(result, arr1[i:]...)
	}

	return result
}

// mid := l + (r-l)/2 this is done to reduce the integer overflow thats it and nothng much

func MergeSort(arr []int, l int, r int) {

	if l < r {
		mid := l + (r-l)/2
		MergeSort(arr, l, mid)
		MergeSort(arr, mid+1, r)
		Merge(arr, l, mid, r)

	}
}

/* it has o(n) and n auxillary space*/
func Merge(arr []int, low int, mid int, high int) {

	left := make([]int, mid-low+1)
	copy(left, arr[low:mid+1])

	right := make([]int, (high - mid))
	copy(right, arr[mid+1:high+1])

	i := 0
	j := 0
	k := low

	for i < len(left) && j < len(right) {
		if left[i] <= right[j] {
			arr[k] = left[i]
			k++
			i++
		} else {
			arr[k] = right[j]
			j++
			k++
		}

	}

	for i < len(left) {
		arr[k] = left[i]
		i++
		k++

	}
	for j < len(right) {
		arr[k] = right[j]
		j++
		k++

	}
}
