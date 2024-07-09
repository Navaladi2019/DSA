package linkedlists

func DetectLoopAndRemove(h *NodeSingle) {

	slow := h
	fast := h

	for fast != nil && fast.next != nil {
		slow = slow.next
		fast = fast.next.next
		if slow == fast {
			break
		}

	}

	if fast != slow {
		return
	}

	slow = h

	for slow.next != fast.next {
		slow = slow.next
		fast = fast.next
	}
	fast.next = nil

}
