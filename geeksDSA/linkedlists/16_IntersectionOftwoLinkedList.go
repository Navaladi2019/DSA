package linkedlists

func FindIntersectionOfTwoLinkedList(l1 SinglyLinkedList, l2 SinglyLinkedList) *NodeSingle {

	l1Coun, l2Coun := 0, 0

	for curr := l1.head; curr != nil; curr = curr.next {
		l1Coun++
	}

	for curr := l2.head; curr != nil; curr = curr.next {
		l2Coun++
	}

	targetLL := l2
	diff := l1Coun - l2Coun

	if l2Coun > l1Coun {
		diff = l1Coun - l2Coun
		targetLL = l2
	}

	var intersection *NodeSingle = targetLL.head

	for i := 1; i < diff; i++ {
		intersection = intersection.next
	}

	return intersection
}
