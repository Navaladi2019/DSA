package heap

import (
	"container/heap"
	"fmt"
)

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
	// heap.Push(&h, 12)
	// heap.Push(&h, 13)
	// heap.Push(&h, 14)
	// heap.Push(&h, 15)
	// heap.Push(&h, 16)
	// heap.Push(&h, 17)
	// heap.Push(&h, 18)
	// heap.Push(&h, 19)
	// heap.Push(&h, 20)
	// heap.Push(&h, 21)
	heap.Pop(&h)
	if len(h) > 0 {
		obj := heap.Pop(&h)
		fmt.Println(obj)
	}

}

func topKFrequent(nums []int, k int) []int {

	m := make(map[int]int)

	for i := 0; i < len(nums); i++ {
		m[nums[i]] += 1
	}

	arr := make([]pairs, 0, len(m))
	for k, v := range m {

		arr = append(arr, pairs{value: k, count: v})
	}

	h := pairsMinHeap(arr)

	heap.Init(&h)

	res := make([]int, 0, k)

	for i := 0; i < k; i++ {

		if len(h) > 0 {
			ele := heap.Pop(&h)
			res = append(res, ele.(pairs).value)
		}
	}

	return res
}

type pairs struct {
	value int
	count int
}

type pairsMinHeap []pairs

func (h pairsMinHeap) Len() int           { return len(h) }
func (h pairsMinHeap) Less(i, j int) bool { return h[i].count > h[j].count }
func (h pairsMinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *pairsMinHeap) Push(ele any) {
	*h = append(*h, ele.(pairs))
}

func (h *pairsMinHeap) Pop() any {
	old := *h
	n := len(*h)
	ele := old[n-1]
	*h = old[:n-1]
	return ele
}
