package heap

import "container/heap"

type Item []int

func (item Item) Len() int { return len(item) }

func (item Item) Less(i, j int) bool {
	return item[i] > item[j]
}

func (pq Item) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (item *Item) Push(x any) {
	val := x.(int)
	*item = append(*item, val)
}

func (item *Item) Pop() any {
	old := *item
	n := len(old)
	ele := old[n-1]
	*item = old[0 : n-1]
	return ele
}

func getMinHeap(arr []int) {
	h := Item(arr)
	heap.Init(&h)
	heap.Push(&h, 11)
	heap.Push(&h, 12)
	heap.Push(&h, 13)
	heap.Push(&h, 14)
	heap.Push(&h, 15)
	heap.Push(&h, 16)
	heap.Push(&h, 17)
	heap.Push(&h, 18)
	heap.Push(&h, 19)
	heap.Push(&h, 20)
	heap.Push(&h, 21)

}
