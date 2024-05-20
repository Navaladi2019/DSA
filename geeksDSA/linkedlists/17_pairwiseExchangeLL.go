package linkedlists

//12345 -> 21435
func PairwiseSwaP(N *NodeSingle) *NodeSingle {

	if N.next == nil {
		return N
	}

	n1 := N

	var res *NodeSingle
	var resEnd *NodeSingle

	for n1 != nil && n1.next != nil {

		n2 := n1.next
		tempn1 := n2.next
		n2.next = n1

		if res == nil {
			res = n2
		} else {
			resEnd.next = n2
		}

		resEnd = n2.next

		n1 = tempn1

	}

	if n1 != nil && n1.next == nil {
		resEnd.next = n1
		resEnd = resEnd.next
	}

	resEnd.next = nil

	return res

}

//12345 -> 21435
func PairwiseSwap(N *NodeSingle) *NodeSingle {

	if N.next == nil {
		return N
	}

	prev := N
	N = N.next

	curr := N.next
	N.next = prev

	for curr != nil && curr.next != nil {

		tempCurr := curr

		tempNextCurr := curr.next.next

		curr = curr.next
		curr.next = tempCurr

		prev.next = curr

		prev = curr.next

		curr = tempNextCurr

	}

	prev.next = curr

	return N

}
