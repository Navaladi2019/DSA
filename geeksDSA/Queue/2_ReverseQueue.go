package queue

// if the size propertyn is not available then we have to use stack to reverse it
func ReverseQueue(q ArrQueue[int]) ArrQueue[int] {

	size := q.Size()

	for size > 0 {
		val, _ := q.Dequeue()
		q.Enqueue(val)
		size--
	}

	return q
}
