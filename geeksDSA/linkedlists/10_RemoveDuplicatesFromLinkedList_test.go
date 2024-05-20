package linkedlists

import (
	"fmt"
	"testing"
)

func Test_RemoveDuplicateFromSortedLL(t *testing.T) {

	l := SinglyLinkedList{}
	l.Insert(1, 1, 2, 3, 4, 5, 5)

	RecursiveLinkedList(l.head)

	fmt.Println("****")

	l = RemoveDuplicateFromLL(l)

	RecursiveLinkedList(l.head)

}
