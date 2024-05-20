package linkedlists

import "fmt"

type CNode struct {
	next *CNode

	data interface{}
}

type CirCulatLinkedList struct {
	head *CNode
	tail *CNode
}

func (l *CirCulatLinkedList) Insert(val interface{}) {

	if l.head == nil {
		l.head = &CNode{data: val}
		l.head.next = l.head
		l.tail = l.head
	} else {

		l.tail.next = &CNode{data: val}
		l.tail = l.tail.next
		l.tail.next = l.head

	}
}

func (l *CirCulatLinkedList) InsertAtBegining(val interface{}) {

	if l.head == nil {
		l.head = &CNode{data: val}
		l.head.next = l.head
		l.tail = l.head
	} else {

		tmpHead := l.head
		l.head = &CNode{data: val}
		l.head.next = tmpHead
		l.tail.next = l.head

	}
}
func (l *CirCulatLinkedList) DeleteHead() {

	if l.head == nil || l.head.next == l.head {
		l.head = nil
	} else {
		l.head = l.head.next
		l.tail.next = l.head
	}
}

func (l *CirCulatLinkedList) DeleteAt(index int) {

	if l.head == nil {
		return
	}

	if index == 0 {
		l.DeleteHead()
	} else {
		curIndex := 1
		prevNode := l.head
		for curNode := l.head.next; curNode != l.head && curIndex <= index; curNode = curNode.next {
			if curIndex == index {
				prevNode.next = curNode.next

				if prevNode.next == l.head {
					l.tail = l.head
				}
			} else {
				prevNode = curNode
			}
			curIndex++
		}

	}
}

func (l *CirCulatLinkedList) Iterate() {
	if l.head == nil {
		return
	}

	fmt.Println(l.head.data)
	for CurNode := l.head.next; CurNode != l.head; CurNode = CurNode.next {

		fmt.Println(CurNode.data)
	}
}
