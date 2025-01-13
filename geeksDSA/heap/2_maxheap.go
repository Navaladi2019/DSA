package heap

import "math"

type MaxHeap struct {
	Heap
}

func (h *MaxHeap) Insert(val int) {
	h.arr = append(h.arr, val)
	for i := len(h.arr) - 1; i > 0; {

		if h.arr[h.parent(i)] < h.arr[i] {
			h.arr[h.parent(i)], h.arr[i] = h.arr[i], h.arr[h.parent(i)]
		}

		i = h.parent(i)
	}
}

func (h *MaxHeap) Heapify(i int) {

	left := h.Left(i)
	right := h.Right(i)
	large := i

	if h.Len() > left && h.arr[left] > h.arr[large] {
		large = left
	}

	if h.Len() > right && h.arr[right] > h.arr[large] {
		large = right
	}

	if large != i {
		h.arr[large], h.arr[i] = h.arr[i], h.arr[large]
		h.Heapify(large)
	}
}

func (h *MaxHeap) Extract() (ok bool, val int) {
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

func (h *MaxHeap) IncreaseKey(i int, val int) {
	if i > h.Len()-1 {
		return
	}
	h.arr[i] = val
	for i > 0 && h.arr[h.parent(i)] < h.arr[i] {
		h.arr[h.parent(i)], h.arr[i] = h.arr[i], h.arr[h.parent(i)]
		i = h.parent(i)
	}
}

// when we call increase key minint will move to the root
//and when we call extract max it deletes root node
func (h *MaxHeap) DeleteKey(i int) {
	h.IncreaseKey(i, math.MaxInt)
	h.Extract()
}

// time complexity is o(n)
func (h *MaxHeap) BuildHeap() {
	for i := h.parent(h.Len() - 1); i >= 0; i-- {
		h.Heapify(i)
	}
}

func (h *MaxHeap) Init(arr []int) {
	h.arr = arr
	h.BuildHeap()
}
