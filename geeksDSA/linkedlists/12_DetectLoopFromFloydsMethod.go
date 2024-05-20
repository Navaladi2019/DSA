package linkedlists

func DetectLoopInFloyds(l *NodeSingle) bool {

	slow := l
	fast := l

	for fast != nil && fast.next != nil {
		slow = slow.next
		fast = fast.next.next
		if slow == fast {
			return true
		}

	}

	return false
}
