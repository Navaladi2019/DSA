package linkedlists

import (
	"fmt"
	"testing"
)

func Test_SimpleLinkedList(t *testing.T) {

	LinkedListNaval1s := SinglyLinkedList{}

	LinkedListNaval1s.Insert(10)
	LinkedListNaval1s.Insert(20)
	LinkedListNaval1s.Insert(30)
	LinkedListNaval1s.Insert(40)
	LinkedListNaval1s.Insert(50)

	for current := LinkedListNaval1s.head; current != nil; current = current.next {
		t.Log(current.data)
	}

}

func Test_SimpleLinkedListRec(t *testing.T) {

	LinkedListNaval1s := SinglyLinkedList{}

	LinkedListNaval1s.Insert(10)
	LinkedListNaval1s.Insert(20)
	LinkedListNaval1s.Insert(30)
	LinkedListNaval1s.Insert(40)
	LinkedListNaval1s.Insert(50)
	LinkedListNaval1s.InsertAt(1000, 3)
	RecursiveLinkedList(LinkedListNaval1s.head)

	i := LinkedListNaval1s.SearchIterative(20)

	if i != 1 {
		t.Error("Has Error in search Iterative")
	}

	i = LinkedListNaval1s.SearchIterative(30)

	if i != 2 {
		t.Error("Has Error in search Iterative")
	}

	i = LinkedListNaval1s.SearchRecursive(30, LinkedListNaval1s.head, 0)

	if i != 2 {
		t.Error("Has Error in search Iterative")
	}

	i = LinkedListNaval1s.SearchRecursive(20, LinkedListNaval1s.head, 0)

	if i != 1 {
		t.Error("Has Error in search Iterative")
	}

	i = LinkedListNaval1s.SearchRecursive(200, LinkedListNaval1s.head, 0)

	if i != -1 {
		t.Error("Has Error in search Iterative")
	}
	t.Log("passed all tests in recu")

}

func Test_SortedInsert(t *testing.T) {

	l := SinglyLinkedList{}

	l.InsertSortedWay(10)
	RecursiveLinkedList(l.head)

	fmt.Println("***")

	l.InsertSortedWay(1)
	RecursiveLinkedList(l.head)

	fmt.Println("***")

	fmt.Println("***")
	l.InsertSortedWay(2)
	RecursiveLinkedList(l.head)

	fmt.Println("***")
	l.InsertSortedWay(20)
	RecursiveLinkedList(l.head)

	fmt.Println("***")
	l.InsertSortedWay(30)
	RecursiveLinkedList(l.head)

	fmt.Println("***")

	l.InsertSortedWay(0)
	RecursiveLinkedList(l.head)

	fmt.Println("***")

	l.InsertSortedWay(1)
	RecursiveLinkedList(l.head)
}
