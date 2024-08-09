package heap

import "fmt"

func KLargestElement(arr []int, k int) {

	h := MinHeap{}
	h.Init(arr[0:k])

	for i := k; i < len(arr); i++ {
		if h.arr[0] < arr[i] {
			h.Extract()
			h.Insert(arr[i])
		}
	}
	for h.Len() > 0 {
		fmt.Println(h.Extract())
	}
}
