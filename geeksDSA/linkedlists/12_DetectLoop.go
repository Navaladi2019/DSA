package linkedlists

func DetectLoopInLLUsingDict(l *NodeSingle) bool {

	dict := make(map[*NodeSingle]struct{})

	for curr := l; curr != nil; curr = curr.next {

		if _, ok := dict[curr]; ok {

			return true
		} else {
			dict[curr] = struct{}{}
		}
	}

	return false
}

// here we are modifying the entrire LL once we visit the node we point the nodes next to dummy node,
// so when a loop exists we will have a nodes next pointing to dummy node
func DetectLoopInLLUsingPointerSwapping(l *NodeSingle) bool {

	dummynode := &NodeSingle{}

	curr := l.next
	for curr != nil {

		if curr.next == dummynode || curr == dummynode {
			return true
		} else {

			oldcur := curr

			curr = curr.next

			oldcur.next = dummynode
		}
	}

	return false
}
