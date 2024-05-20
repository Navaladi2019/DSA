package linkedlists

/*
[10,20,30,40]
[5,15,16,18,35]
*/
func MergeTwoSortedLinkedList(a SinglyLinkedList, b SinglyLinkedList) SinglyLinkedList {

	l := SinglyLinkedList{}

	aNext := a.head
	bNext := b.head

	if aNext == nil {
		l.head = bNext
		return l
	}

	if bNext == nil {
		l.head = aNext
		return l
	}

	if a.head.data < b.head.data {
		l.head = a.head
		aNext = a.head.next
	} else {
		l.head = b.head
		bNext = b.head.next
	}

	curr := l.head

	for aNext != nil && bNext != nil {

		aNextTemp := aNext.next
		bNextTemp := bNext.next
		if aNext.data < bNext.data {
			curr.next = aNext
			aNext = aNextTemp
		} else {
			curr.next = bNext
			bNext = bNextTemp
		}
		curr = curr.next
	}

	if aNext != nil {
		curr.next = aNext
	}

	if bNext != nil {
		curr.next = bNext
	}

	return l
}
