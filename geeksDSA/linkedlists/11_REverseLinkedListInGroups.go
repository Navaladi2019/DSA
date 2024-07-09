package linkedlists

/*Reverses SingleyLinkedList in groups

[10,20,30,40,50,60] -> [30,20,10,60,50,40] for k =3

[10,20,30,40,50] -> [30,20,10,50,40] for k =3

[10,20,30] -> [30,20,10] for k =4
*/
func reverseLinkedListInGroups(l SinglyLinkedList, k int) SinglyLinkedList {

	prev := l.head
	head, curr := ReverseLinkedInGroup(l.head, nil, k)
	l.head = head

	for curr != nil {
		tempcurr := curr
		prev.next, curr = ReverseLinkedInGroup(curr, nil, k)
		prev = tempcurr
	}

	return l
}

//[10,20,30,40,50,60] -> [30,20,10,60,50,40] for k =3
func ReverseLinkedInGroup(n *NodeSingle, curperv *NodeSingle, k int) (*NodeSingle, *NodeSingle) {

	curr := n
	prev := curperv

	for i := 0; i < k && curr != nil; i++ {
		tempCurr := curr.next
		curr.next = prev
		prev = curr
		curr = tempCurr
	}

	return prev, curr
}
