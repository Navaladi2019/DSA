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

func RemoveDuplicateFromLL_1(l SinglyLinkedList) SinglyLinkedList {

	curr := l.head

	for curr != nil {

		if curr.next == nil {
			break
		}

		for curr.next != nil && curr.next.data == curr.data {
			curr.next = curr.next.next
		}

		curr = curr.next
	}

	return l
}

func RemoveDuplicateFromLL_Efficient(curr *NodeSingle, prev *NodeSingle) *NodeSingle {

	if curr == nil {
		return nil
	}

	if curr.data == prev.data {
		prev.next = RemoveDuplicateFromLL_Efficient(curr.next, curr)
	}

	return curr
}
