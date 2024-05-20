package linkedlists

func FindMiddleOfLinkedList(l SinglyLinkedList) int {

	if l.head == nil {
		return -1
	}

	fast := l.head
	slow := l.head

	for fast != nil && fast.next != nil {

		if fast.next == nil {
			fast = nil
		} else {
			fast = fast.next.next
		}
		slow = slow.next
	}

	return slow.data

}
