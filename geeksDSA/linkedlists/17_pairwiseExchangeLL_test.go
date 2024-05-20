package linkedlists

import "testing"

//12345 -> 21435
func Test_PairwiseSwaP(t *testing.T) {

	l := SinglyLinkedList{}
	l.Insert(1, 2, 3, 4, 5, 6, 7)
	o := PairwiseSwap(l.head)
	RecursiveLinkedList(o)
}
