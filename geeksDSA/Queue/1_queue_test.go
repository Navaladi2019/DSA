package queue

import (
	"fmt"
	"testing"
)

func Test_ArrQueue(t *testing.T) {

	q := ArrQueue[int]{}

	q.Init(10)

	q.Enqueue(1)
	q.Enqueue(2)
	fmt.Println(q.GetFront())
	fmt.Println(q.GetRear())
	q.Enqueue(3)
	fmt.Println(q.GetFront())
	fmt.Println(q.GetRear())
	fmt.Println(q.Size())
	fmt.Println(q.Dequeue())
	fmt.Println(q.Dequeue())
	fmt.Println(q.Dequeue())
	fmt.Println(q.Size())
}
