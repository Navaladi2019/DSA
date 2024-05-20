package linkedlists

import "fmt"

type Node struct {
	next *Node
	prev *Node
	data interface{}
}

type DoublyLinkedList struct {
	head  *Node
	tail  *Node
	count int
}

func (l *DoublyLinkedList) Insert(val interface{}) {

	if l.head == nil {
		l.head = &Node{data: val}
		l.tail = l.head
	} else {
		l.tail.next = &Node{data: val, prev: l.tail}
		l.tail = l.tail.next
	}
	l.count++
}

func (l *DoublyLinkedList) InsertAtBegining(val interface{}) {

	if l.head == nil {
		l.head = &Node{data: val}
		l.tail = l.head
	} else {
		n := &Node{data: val, prev: nil, next: l.head}
		l.head.prev = n
		l.head = n
	}
	l.count++
}

func (l *DoublyLinkedList) InsertAtEnd(val interface{}) {

	if l.head == nil {
		l.head = &Node{data: val}
		l.tail = l.head
	} else {
		n := &Node{data: val, prev: l.tail, next: nil}
		l.tail.next = n
		l.tail = n
	}
	l.count++
}

func (l *DoublyLinkedList) Reverse() {

	currNode := l.head

	for currNode != nil {
		tempNext := currNode.next
		currNode.next = currNode.prev
		currNode.prev = tempNext
		currNode = currNode.prev
	}

	l.head, l.tail = l.tail, l.head

}

func RecursiveLinkedListNext(n *Node) {

	if n == nil {
		return
	}
	fmt.Println(n.data)
	RecursiveLinkedListNext(n.next)
}

func RecursiveLinkedListPrev(n *Node) {

	if n == nil {
		return
	}
	fmt.Println(n.data)
	RecursiveLinkedListPrev(n.prev)
}

func (l *DoublyLinkedList) DeleteHead() {

	if l.head == nil || l.head.next == nil {
		l.head = nil
		l.tail = nil
	} else {
		l.head = l.head.next
		l.head.prev = nil
	}
}

func (l *DoublyLinkedList) DeleteLast() {

	if l.head == nil || l.head.next == nil {
		l.head = nil
		l.tail = nil
	} else {
		l.tail = l.tail.prev
		l.tail.next = nil
	}

}
