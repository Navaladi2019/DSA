package linkedlists

import (
	"math"
)

func FindIntersectionOfTwoLinkedList(l1 SinglyLinkedList, l2 SinglyLinkedList) *NodeSingle {

	l1Coun, l2Coun := 0, 0

	for curr := l1.head; curr != nil; curr = curr.next {
		l1Coun++
	}

	for curr := l2.head; curr != nil; curr = curr.next {
		l2Coun++
	}

	targetLL := l2.head
	passll := l1.head
	diff := l1Coun - l2Coun

	if l2Coun < l1Coun {
		targetLL = l1.head
		passll = l2.head
	}

	for i := 0; i < int(math.Abs(float64(diff))); i++ {
		targetLL = targetLL.next

	}

	var intersection *NodeSingle

	for targetLL != nil && passll != nil {

		if targetLL == passll {
			intersection = targetLL
			break
		}
		targetLL = targetLL.next
		passll = passll.next

	}

	return intersection
}

func FindIntersectionOfTwoLinkedList_1(l1 SinglyLinkedList, l2 SinglyLinkedList) *NodeSingle {

	l1Count := 0
	l2Count := 0

	for curr := l1.head; curr != nil; curr = curr.next {
		l1Count++
	}

	for curr := l2.head; curr != nil; curr = curr.next {
		l2Count++
	}

	var LeadNode, laggingNode *NodeSingle

	var diff int
	if l1Count >= l2Count {
		diff = l1Count - l2Count
		LeadNode = l1.head
		laggingNode = l2.head
	} else {
		diff = l2Count - l1Count
		LeadNode = l2.head
		laggingNode = l1.head
	}

	for i := 0; i < diff; i++ {
		LeadNode = LeadNode.next
	}

	for LeadNode != nil && laggingNode != nil {

		if LeadNode.next == laggingNode.next {
			return LeadNode.next
		}
		LeadNode = LeadNode.next
		laggingNode = laggingNode.next
	}

	return nil

}
