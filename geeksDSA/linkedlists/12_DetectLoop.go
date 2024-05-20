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

func DetectLoopInLLUsingPointerSwapping(l *NodeSingle) bool {

	temp := &NodeSingle{}

	curr := l.next
	for curr != nil {

		if curr.next == temp || curr == temp {
			return true
		} else {

			oldcur := curr

			curr = curr.next

			oldcur.next = temp
		}
	}

	return false
}
