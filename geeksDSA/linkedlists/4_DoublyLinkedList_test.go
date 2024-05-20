package linkedlists

import (
	"testing"
)

func Test_Double(t *testing.T) {

	l := DoublyLinkedList{}

	l.Insert(10)
	l.Insert(20)
	l.Insert(30)

	l.DeleteLast()
	RecursiveLinkedListNext(l.head)
}
