package linkedlists

//[10,5,4,3,6]
func SegregateOddEvenNodes(l SinglyLinkedList) SinglyLinkedList {

	var oddStart, OddEnd, evenStart, evenEnd *NodeSingle
	curr := l.head
	for curr != nil {

		tmpNext := curr.next
		if curr.data%2 == 0 {
			if evenStart == nil {
				evenStart, evenEnd = curr, curr
			} else {
				evenEnd.next = curr
				evenEnd = evenEnd.next
			}

		} else {
			if oddStart == nil {
				oddStart, OddEnd = curr, curr
			} else {
				OddEnd.next = curr
				OddEnd = OddEnd.next
			}

		}

		curr = tmpNext
	}

	if evenEnd != nil {
		evenEnd.next = nil
	}

	if OddEnd != nil {
		OddEnd.next = nil
	}

	if evenStart == nil {
		l.head = oddStart
	} else {
		evenEnd.next = oddStart
		l.head = evenStart
	}

	if oddStart == nil {
		l.head = evenStart
	}

	return l
}

func SegregateOddEvenNodes_1(l SinglyLinkedList) SinglyLinkedList {

	curr := l.head

	var EvenStart, EvenEnd, oddStart, OddEnd *NodeSingle

	for curr != nil {

		tempCurr := curr.next

		if curr.data%2 == 0 {

			if EvenStart == nil {
				EvenStart, EvenEnd = curr, curr

			} else {
				EvenEnd.next = curr
				EvenEnd = EvenEnd.next
				EvenEnd.next = nil
			}
		} else {
			if oddStart == nil {
				oddStart, OddEnd = curr, curr
			} else {
				OddEnd.next = curr
				OddEnd = OddEnd.next
				OddEnd.next = nil
			}
		}

		curr = tempCurr

	}

	if EvenStart != nil {
		l.head = EvenStart
		EvenEnd.next = oddStart
	} else {
		l.head = oddStart
	}

	return l
}

func SegregateOddEvenNodes_Efficient(l SinglyLinkedList) SinglyLinkedList {

	var os, oe, es, ee *NodeSingle

	curr := l.head

	for curr != nil {
		tempnext := curr.next
		if curr.data%2 == 0 {
			if es == nil {
				es, ee = curr, curr
			} else {
				ee.next = curr
				ee = ee.next
			}

			ee.next = os
		} else {
			if os == nil {
				os = curr
				oe = curr
			} else {
				oe.next = curr
				oe = oe.next
			}

			oe.next = nil
		}

		curr = tempnext

	}

	if es != nil {
		ee.next = os
		l.head = es
	} else {
		l.head = os
	}

	return l
}
