package linkedlists

// considering the node we are goingto delete is not last
func DeleteNode(n *NodeSingle) {

	if n.next == nil {
		return
	}

	n.data = n.next.data
	n.next = n.next.next
}
