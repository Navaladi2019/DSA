package sorting

func MaxHeap(arr []int, length int, i int) {

	left := 2*i + 1
	right := 2*i + 2

	largest := i

	if left < length && arr[left] > arr[i] {
		largest = left
	}

	if right < length && arr[right] > arr[i] {
		largest = right
	}

	if largest != i {

		arr[i], arr[largest] = arr[largest], arr[i]
		MaxHeap(arr, length, largest)
	}
}

func buildHeap(arr []int) {

	for i := ((len(arr)) - 2) / 2; i >= 0; i-- {
		MaxHeap(arr, len(arr), i)
	}
}

func SortUsingHeap(arr []int) {

	buildHeap(arr)

	for i := len(arr) - 1; i > 0; i-- {

		arr[i], arr[0] = arr[0], arr[i]
		MaxHeap(arr, i, 0)
	}
}
