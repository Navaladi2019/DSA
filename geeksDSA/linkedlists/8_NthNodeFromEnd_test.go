package linkedlists

import "testing"

func Test_NthNodeFromEnd(t *testing.T) {

	l := SinglyLinkedList{}
	l.Insert(10, 20, 30, 40, 50)
	got, err := FindNthNodeFromEnd(l, 2)

	if got != 40 || err != nil {
		t.Error("has error in 40")
	}

	l = SinglyLinkedList{}
	l.Insert(10, 20, 30)
	got, err = FindNthNodeFromEnd(l, 3)

	if got != 10 || err != nil {
		t.Error("has error in 10")
	}

	l = SinglyLinkedList{}
	l.Insert(10, 20, 30)
	got, err = FindNthNodeFromEnd(l, 2)

	if got != 20 || err != nil {
		t.Error("has error in 10")
	}

	l = SinglyLinkedList{}
	l.Insert(10, 20)
	got, err = FindNthNodeFromEnd(l, 3)

	if got != -1 || err == nil {
		t.Error("has error in err")
	}
}
