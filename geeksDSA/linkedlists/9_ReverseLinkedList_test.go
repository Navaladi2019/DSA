package linkedlists

import (
	"fmt"
	"testing"
)

func Test_CheckReverse(t *testing.T) {
	l := SinglyLinkedList{}
	l.Insert(1, 2, 3, 4, 5)

	RecursiveLinkedList(l.head)

	fmt.Println("***")

	res := ReverseLinkedListRecursively1(l.head, nil)

	RecursiveLinkedList(res)
	fmt.Println("***")
	l = SinglyLinkedList{}
	l.Insert(1)

	RecursiveLinkedList(l.head)

	fmt.Println("***")

	res = ReverseLinkedListRecursively1(l.head, nil)

	RecursiveLinkedList(res)
}
