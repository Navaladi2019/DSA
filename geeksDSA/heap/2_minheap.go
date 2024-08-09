package heap

import "math"

type MinHeap struct {
	Heap
}

func (h *MinHeap) Insert(val int) {
	h.arr = append(h.arr, val)
	for i := len(h.arr) - 1; i > 0; {

		if h.arr[h.parent(i)] > h.arr[i] {
			h.arr[h.parent(i)], h.arr[i] = h.arr[i], h.arr[h.parent(i)]
		}
		i = h.parent(i)
	}
}

func (h *MinHeap) Heapify(i int) {

	left := h.Left(i)
	right := h.Right(i)
	small := i

	if h.Len() > left && h.arr[left] < h.arr[small] {
		small = left
	}

	if h.Len() > right && h.arr[right] < h.arr[small] {
		small = right
	}

	if small != i {
		h.arr[small], h.arr[i] = h.arr[i], h.arr[small]
		h.Heapify(small)
	}
}

func (h *MinHeap) Extract() (ok bool, val int) {
	if h.Len() > 0 {
		val = h.arr[0]
		ok = true
	}
	if h.Len() > 1 {
		h.arr[0] = h.arr[len(h.arr)-1]
		h.arr = h.arr[:h.Len()-1]
		h.Heapify(0)
	} else {
		h.arr = h.arr[:0]
	}

	return
}

func (h *MinHeap) DecreaseKey(i int, val int) {
	if i > h.Len()-1 {
		return
	}
	h.arr[i] = val
	for i > 0 && h.arr[h.parent(i)] > h.arr[i] {
		h.arr[h.parent(i)], h.arr[i] = h.arr[i], h.arr[h.parent(i)]
		i = h.parent(i)
	}
}

// when we call decerese key minint will move to the root
//and when we call extract min it deletes root node
func (h *MinHeap) DeleteKey(i int) {
	h.DecreaseKey(i, math.MinInt)
	h.Extract()
}

// time complexity is o(n)
func (h *MinHeap) BuildHeap() {
	for i := h.parent(h.Len() - 1); i >= 0; i-- {
		h.Heapify(i)
	}
}

func (h *MinHeap) Init(arr []int) {
	h.arr = arr
	h.BuildHeap()
}
