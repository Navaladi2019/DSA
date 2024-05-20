package linkedlists

/*Reverses SingleyLinkedList in groups

[10,20,30,40,50,60] -> [30,20,10,60,50,40] for k =3

[10,20,30,40,50] -> [30,20,10,50,40] for k =3

[10,20,30] -> [30,20,10] for k =4
*/
func reverseLinkedListInGroups(l SinglyLinkedList, k int) SinglyLinkedList {

	prev := l.head

	head, tempCurr := ReverseLinkedListByGroup(prev, nil, k)
	l.head = head

	for tempCurr != nil {
		prev.next, tempCurr = ReverseLinkedListByGroup(tempCurr, nil, k)
		prev = tempCurr
	}
	return l
}

func ReverseLinkedListByGroup(node *NodeSingle, prev *NodeSingle, k int) (*NodeSingle, *NodeSingle) {

	curr := node

	i := 0
	for curr != nil && i < k {
		i++
		tempCurr := curr.next
		curr.next = prev
		prev = curr
		curr = tempCurr

	}

	return prev, curr
}
