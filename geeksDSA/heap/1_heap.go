package heap

import "sort"

type Heap struct {
	sort.Interface
	arr []int
}

func (h *Heap) Left(i int) int {
	return 2*i + 1
}

func (h *Heap) Right(i int) int {
	return 2*i + 2
}

func (h *Heap) parent(i int) int {
	return (i - 1) / 2
}

func (h *Heap) Cap() int {
	return cap(h.arr)
}

func (h *Heap) Len() int {
	return len(h.arr)
}

type iHeap interface {
	Extract() (ok bool, val int)
	BuildHeap()
}
