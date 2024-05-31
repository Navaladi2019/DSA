package queue

type ArrQueue[T any] struct {
	arr []T
	cap int
}

func (q *ArrQueue[T]) Init(cap int) {
	q.arr = make([]T, 0, cap)
	q.cap = cap
}

func (q *ArrQueue[T]) Enqueue(val T) bool {

	if len(q.arr) == q.cap {
		return false
	}
	q.arr = append(q.arr, val)

	return true

}

func (q *ArrQueue[T]) Dequeue() (val T, ok bool) {

	if len(q.arr) == 0 {
		return
	}
	val = q.arr[0]
	ok = true
	q.arr = q.arr[1:]
	return
}

func (q *ArrQueue[T]) Size() int {
	return len(q.arr)
}

func (q *ArrQueue[T]) Cap() int {
	return q.cap
}

func (q *ArrQueue[T]) GetFront() (val T, ok bool) {
	if len(q.arr) == 0 {
		return
	}
	val = q.arr[0]
	ok = true
	return
}

func (q *ArrQueue[T]) GetRear() (val T, ok bool) {
	if len(q.arr) == 0 {
		return
	}
	val = q.arr[len(q.arr)-1]
	ok = true
	return
}

func (q *ArrQueue[T]) IsFull() bool {
	return len(q.arr) == q.cap
}

func (q *ArrQueue[T]) IsEmpty() bool {
	return len(q.arr) == 0
}
