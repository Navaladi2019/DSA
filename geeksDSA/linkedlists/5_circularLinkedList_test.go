package linkedlists

import "testing"

func Test_CircularLinkedList(t *testing.T) {

	l := CirCulatLinkedList{}

	l.InsertAtBegining(10)
	l.Insert(20)

	l.Insert(30)
	l.Insert(40)
	l.DeleteAt(0)
	l.Insert(50)

	l.Iterate()
}
