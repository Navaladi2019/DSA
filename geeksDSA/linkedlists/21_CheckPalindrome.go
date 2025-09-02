package linkedlists

func IsPalindrome(l SinglyLinkedList) bool {

	if l.head == nil || l.head.next == nil {
		return true
	}

	length := 0

	for curr := l.head; curr != nil; curr = curr.next {
		length++
	}

	isOdd := false

	if length%2 != 0 {
		isOdd = true
	}

	i := 0

	// Finding Node to reverse From
	curr := l.head
	for curr != nil {

		if isOdd == true && length/2 == i {
			curr = curr.next
			break
		}
		if isOdd == false && length/2 == i {
			break
		}
		curr = curr.next
		i++
	}

	//Reverse From Current

	//[1,2,3,4,5]

	var previous *NodeSingle

	for curr != nil {

		tempNext := curr.next
		curr.next = previous
		previous = curr
		curr = tempNext

	}

	start := l.head
	for previous != nil {

		if previous.data != start.data {
			return false
		}

		previous = previous.next
		start = start.next
	}

	return true
}

func IsPalindrome_Efficient(l SinglyLinkedList) bool {

	h := l.head

	slow, fast := h, h

	for fast.next != nil && fast.next.next != nil {

		fast = fast.next.next
		slow = slow.next
	}

	curr := slow.next

	// here we are reversing linked list
	var prev *NodeSingle
	for curr != nil {
		tempNext := curr.next
		curr.next = prev
		prev = curr
		curr = tempNext

	}

	slow = l.head

	for prev != nil {

		if prev.data != slow.data {
			return false
		}

		slow = slow.next
		prev = prev.next
	}

	return true
}
