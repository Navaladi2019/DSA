package linkedlists

import "testing"

func Test_FindMiddleOfLinkedList(t *testing.T) {

	l := SinglyLinkedList{}
	l.Insert(10, 5, 20, 15, 25)

	got := FindMiddleOfLinkedList(l)

	if got != 20 {
		t.Error("has failed for 20")
	}

	l = SinglyLinkedList{}
	l.Insert(10, 5, 20, 15, 25, 30)

	got = FindMiddleOfLinkedList(l)

	if got != 15 {
		t.Error("has failed for 15")
	}

	l = SinglyLinkedList{}
	l.Insert(10)

	got = FindMiddleOfLinkedList(l)

	if got != 10 {
		t.Error("has failed for 10")
	}

	l = SinglyLinkedList{}

	got = FindMiddleOfLinkedList(l)

	if got != -1 {
		t.Error("has failed for nil")
	}

}
