package linkedlists

import "errors"

//[1,2,3,4,5,6,7]
func FindNthNodeFromEnd(l SinglyLinkedList, k int) (int, error) {

	iNode := l.head
	jNode := l.head

	for i := 0; i < k; i++ {

		if iNode == nil {
			return -1, errors.New("Has insufficient length")
		}

		iNode = iNode.next

	}

	for iNode != nil {
		iNode = iNode.next
		jNode = jNode.next
	}

	return jNode.data, nil
}
