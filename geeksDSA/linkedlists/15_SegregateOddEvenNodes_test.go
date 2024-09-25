package linkedlists

import (
	"fmt"
	"testing"
)

func Test_SegregateOddEvenNodes(t *testing.T) {

	l := SinglyLinkedList{}

	l.Insert(10, 20, 1, 8, 5)

	RecursiveLinkedList(l.head)

	fmt.Println("****")
	l = SegregateOddEvenNodes_1(l)

	RecursiveLinkedList(l.head)
}
