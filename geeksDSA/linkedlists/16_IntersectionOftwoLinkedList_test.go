package linkedlists

import "testing"

func Test_FindIntersectionOfTwoLinkedList(t *testing.T) {

	h := &NodeSingle{data: 5}

	h2 := &NodeSingle{data: 8}

	h3 := &NodeSingle{data: 7}
	h4 := &NodeSingle{data: 10}
	h5 := &NodeSingle{data: 12}
	h6 := &NodeSingle{data: 15}

	h.next = h2
	h2.next = h3
	h3.next = h4
	h4.next = h5
	h5.next = h6

	nn := &NodeSingle{data: 1, next: &NodeSingle{data: 2, next: h4}}

	node := FindINtersection_Efficient(SinglyLinkedList{head: h}, SinglyLinkedList{head: nn})

	if node != h4 {
		t.Error("has error in finding intersection of LL")
	}

}
