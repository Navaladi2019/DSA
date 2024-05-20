package linkedlists

func ReverseSingleLLWithMoreMemory(l SinglyLinkedList) SinglyLinkedList {

	sl := SinglyLinkedList{}

	for h := l.head; h != nil; h = h.next {
		sl.InsertAtBegining(h.data)
	}

	return sl
}

// [10,20,30,40,50]
func ReverseSingleLLEfficient(l SinglyLinkedList) SinglyLinkedList {

	var previous *NodeSingle
	current := l.head

	for current != nil {

		tempCurrent := current.next
		current.next = previous
		previous = current
		current = tempCurrent
	}

	l.head = previous

	return l
}

func ReverseLinkedListRecursively1(current *NodeSingle, p *NodeSingle) *NodeSingle {

	if current == nil {
		return p
	}

	tempNext := current.next
	current.next = p
	p = current
	current = tempNext

	return ReverseLinkedListRecursively1(current, p)

}
