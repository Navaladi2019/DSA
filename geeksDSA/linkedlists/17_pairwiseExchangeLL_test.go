package linkedlists

import "testing"

//12345 -> 21435
func Test_PairwiseSwaP(t *testing.T) {

	l := SinglyLinkedList{}
	l.Insert(1, 2, 3, 4, 5)
	//o := PairWiseSwap_Easy(l.head)
	//o := pairwiseSwap_Changedata(l.head)
	o := pairwiseSwap_Rec(l.head)
	RecursiveLinkedList(o)
}
