package heap

import "container/heap"

func GetMedianOfStream(arr []int) {

	item := Item(arr)
	heap.Init(&item)

}
