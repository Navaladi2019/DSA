package heap

func SortKSorted(arr []int, k int) []int {

	for i := 0; i < len(arr); i = i + k {
		q := MinHeap{}
		maxLength := i + k + 1

		if maxLength > len(arr) {
			maxLength = len(arr)
		}

		newList := make([]int, len(arr[i:maxLength]), len(arr[i:maxLength]))
		copy(newList, arr[i:maxLength])
		q.Init(newList)

		for j := 0; j <= k && i+j < maxLength; j++ {
			_, arr[j+i] = q.Extract()
		}
	}
	return arr
}

func SortKSorted_Efficient(arr []int, k int) []int {
	newList := make([]int, k+1, k+1)
	copy(newList, arr[0:k+1])
	q := MinHeap{}
	q.Init(newList)
	j := k
	for i := 0; i < len(arr); i++ {
		_, val := q.Extract()
		arr[i] = val
		j++
		if j < len(arr) {
			q.Insert(arr[j])
		}
	}

	return arr
}
