package linkedlists

import (
	"fmt"
	"testing"
)

func Test_CDLL(t *testing.T) {
	l := CirCularDoublyLinkedList{}

	l.Insert(10)

	l.Insert(20)

	l.Insert(30)
	l.InsertAtBegining(1)
	l.IterateFromHead()

	fmt.Println("***")
	l.IterateFromTail()
}
