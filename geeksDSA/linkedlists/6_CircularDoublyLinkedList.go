package linkedlists

import "fmt"

type CDNode struct {
	next *CDNode
	prev *CDNode
	data interface{}
}

type CirCularDoublyLinkedList struct {
	head *CDNode
}

func (l *CirCularDoublyLinkedList) Insert(val interface{}) {

	if l.head == nil {
		l.head = &CDNode{data: val}
		l.head.next = l.head
		l.head.prev = l.head
	} else {
		tempNode := &CDNode{data: val, next: l.head, prev: l.head.prev}
		l.head.prev.next = tempNode
		l.head.prev = tempNode

	}

}

func (l *CirCularDoublyLinkedList) InsertAtBegining(val interface{}) {

	if l.head == nil {
		l.head = &CDNode{data: val}
		l.head.next = l.head
		l.head.prev = l.head
	} else {
		tempNode := &CDNode{data: val, next: l.head, prev: l.head.prev}
		l.head.prev.next = tempNode
		l.head.prev = tempNode
		l.head = tempNode

	}

}

func (l *CirCularDoublyLinkedList) IterateFromHead() {

	if l.head == nil {
		return
	}

	fmt.Println(l.head.data)

	for curNode := l.head.next; l.head != curNode; curNode = curNode.next {
		fmt.Println(curNode.data)
	}
}

func (l *CirCularDoublyLinkedList) IterateFromTail() {

	if l.head == nil {
		return
	}

	fmt.Println(l.head.prev.data)

	for curNode := l.head.prev.prev; l.head.prev != curNode; curNode = curNode.prev {
		fmt.Println(curNode.data)
	}
}
