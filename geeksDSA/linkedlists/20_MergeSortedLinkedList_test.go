package linkedlists

import "testing"

func Test_MergeTwoSortedLinkedList(t *testing.T) {

	a := SinglyLinkedList{}
	a.Insert(10)

	b := SinglyLinkedList{}

	c := MergeTwoSortedLinkedList(a, b)

	RecursiveLinkedList(c.head)
}
