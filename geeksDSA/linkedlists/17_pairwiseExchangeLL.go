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

func PairWiseSwap_Easy(n *NodeSingle) *NodeSingle {

	res := n.next

	curr := n

	var prevLink *NodeSingle

	for curr != nil && curr.next != nil {

		Second := curr.next
		nextCurr := Second.next

		Second.next = curr

		if prevLink != nil {
			prevLink.next = Second
		}
		prevLink = curr
		curr = nextCurr
	}

	if curr != nil {
		prevLink.next = curr
	} else {
		prevLink.next = nil
	}

	return res
}

func pairwiseRec(n *NodeSingle) *NodeSingle {

	if n == nil || n.next == nil {
		return n
	}

	curr := n.next

	tempREcNext := curr.next

	curr.next = n

	curr.next.next = pairwiseRec(tempREcNext)

	return curr

}

func pairwiseSwap_Changedata(h *NodeSingle) *NodeSingle {

	curr := h

	for curr != nil && curr.next != nil {

		curr.data, curr.next.data = curr.next.data, curr.data

		curr = curr.next.next
	}

	return h
}

func pairwiseSwap_Rec(h *NodeSingle) *NodeSingle {
	if h == nil || h.next == nil {
		return h
	}

	nextHead := h.next
	tempNext := nextHead.next

	nextHead.next = h

	h.next = pairwiseSwap_Rec(tempNext)

	return nextHead
}
