package queue

import (
	"fmt"
	"testing"
)

func Test_ReverseQueue(t *testing.T) {

	q := ArrQueue[int]{}

	q.Init(10)

	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)
	q.Enqueue(4)
	q.Enqueue(5)
	q.Enqueue(6)

	q = ReverseQueue(q)

	for !q.IsEmpty() {
		fmt.Println(q.Dequeue())
	}
}
