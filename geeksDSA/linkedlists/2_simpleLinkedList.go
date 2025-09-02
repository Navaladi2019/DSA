package linkedlists

import (
	"errors"
)

type NodeSingle struct {
	data int
	next *NodeSingle
}

type NodeSingleRandom struct {
	data   int
	next   *NodeSingleRandom
	random *NodeSingleRandom
}

func (n *NodeSingle) Init(info int) {
	n.data = info
}

type SinglyLinkedList struct {
	head  *NodeSingle
	last  *NodeSingle
	count int
}

func (l *SinglyLinkedList) Insert(vals ...int) {

	for _, val := range vals {
		if l.head == nil {
			l.head = &NodeSingle{data: val}
			l.last = l.head
		} else {
			l.last.next = &NodeSingle{data: val}
			l.last = l.last.next
		}
		l.count++
	}

}

func (l *SinglyLinkedList) InsertAtBegining(val int) {
	n := &NodeSingle{data: val, next: l.head}
	l.head = n
	l.count++
}

func (l *SinglyLinkedList) InsertAt(val int, index int) error {

	if index+1 > l.count {
		return errors.New("Index greater than count")
	}

	if index == 0 {
		l.InsertAtBegining(val)
		l.last = l.head
		return nil
	}

	for node, count := l.head, 0; node != nil; node = node.next {
		if count == index-1 {
			nextNode := &NodeSingle{data: val, next: node.next}
			node.next = nextNode
			l.count++
			break
		}
		count++
	}

	return nil
}

func (l *SinglyLinkedList) DeleteFirst() {
	if l.count > 0 {
		l.head = l.head.next
		l.count--
	}
}

func (l *SinglyLinkedList) DeleteLast() {

	if l.head == nil {
		return
	}
	if l.head.next == nil {
		l.head = nil
		l.last = nil
		l.count--
		return
	}

	curr := l.head
	for curr.next.next != nil {
		curr = curr.next
	}

	l.last = curr
	curr.next = nil
	l.count--
}

func (l *SinglyLinkedList) DeleteLast_1() {

	if l.head == nil || l.head.next == nil {
		l.head = nil
		return
	}

	curr := l.head

	prev := curr

	for curr != nil {
		if curr.next == nil {
			prev.next = nil
			break
		}
		prev = curr
		curr = curr.next
	}
}

func (l *SinglyLinkedList) SearchIterative(val interface{}) int {

	curr := l.head

	curIndex := 0

	for curr != nil {
		if curr.data == val {
			return curIndex
		}

		curr = curr.next
		curIndex++
	}

	return -1
}

func (l *SinglyLinkedList) SearchRecursive(val interface{}, n *NodeSingle, i int) int {

	if n == nil {
		return -1
	}
	if n.data == val {
		return i
	}
	return l.SearchRecursive(val, n.next, i+1)
}

func (l *SinglyLinkedList) InsertSortedWay(val int) {

	if l.head == nil || l.head.data > val {
		l.InsertAtBegining(val)
	} else {

		prevNode := l.head
		currNode := l.head.next

		for currNode != nil {
			if currNode.data > val {
				break
			}
			prevNode = currNode
			currNode = currNode.next

		}

		prevNode.next = &NodeSingle{data: val, next: currNode}

	}

}

func (l *SinglyLinkedList) Reverse() {

	curr := l.head
	var prev *NodeSingle
	for curr != nil {
		tempnext := curr.next
		curr.next = prev
		prev = curr
		curr = tempnext

	}
	l.head = prev
}

func (l *SinglyLinkedList) Reverse_1() {

	var prev *NodeSingle = nil

	curr := l.head

	for curr != nil {
		tempCurr := curr.next
		curr.next = prev
		prev, curr = curr, tempCurr
	}
}
