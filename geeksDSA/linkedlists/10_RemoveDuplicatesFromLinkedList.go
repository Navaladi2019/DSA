package linkedlists

func RemoveDuplicateFromLL(l SinglyLinkedList) SinglyLinkedList {

	lastUniqueNode := l.head

	curr := l.head.next

	for curr != nil {
		if lastUniqueNode.data != curr.data {
			lastUniqueNode.next = curr
			lastUniqueNode = curr
		}
		curr = curr.next
	}
	lastUniqueNode.next = nil
	return l
}
