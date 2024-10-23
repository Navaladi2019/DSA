package linkedlists

func CloneLinkedListWithDict(h *NodeSingleRandom) *NodeSingleRandom {

	dict := make(map[*NodeSingleRandom]*NodeSingleRandom)

	for curr := h; curr != nil; curr = curr.next {
		dict[curr] = &NodeSingleRandom{data: curr.data}
	}

	res := dict[h]

	for curr := h; curr != nil; curr = curr.next {

		if curr.next != nil {
			dict[curr].next = dict[curr.next]
		}
		if curr.random != nil {
			dict[curr].random = dict[curr.random]
		}
	}

	return res

}
