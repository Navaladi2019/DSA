package heap

type PriorityQueueMine struct {
	heap MaxHeap
}

func CreateNewPriorityQueue(arr []int) PriorityQueueMine {
	q := PriorityQueueMine{}
	q.heap.arr = arr
	q.heap.BuildHeap()
	return q
}
func (q *PriorityQueueMine) Pop() (ok bool, value int) {
	return q.heap.Extract()
}

func (q *PriorityQueueMine) Peek() (ok bool, value int) {
	if q.heap.Len() >= 1 {
		ok = true
		value =
			q.heap.arr[0]
	}
	return
}

func (q *PriorityQueueMine) Push(val int) {
	q.heap.Insert(val)
}
func (q *PriorityQueueMine) IsEmpty() bool {
	return q.heap.Len() == 0
}
