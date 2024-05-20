package linkedlists

import "testing"

func Test_DetectLoopInLLUsingDict(t *testing.T) {

	h := new(NodeSingle)
	h.data = 1

	n2 := new(NodeSingle)
	n2.data = 2

	n3 := new(NodeSingle)
	n3.data = 3
	n4 := new(NodeSingle)
	n4.data = 4

	n5 := new(NodeSingle)
	n5.data = 5
	n6 := new(NodeSingle)
	n6.data = 6
	n7 := new(NodeSingle)
	n7.data = 6

	h.next = n2
	n2.next = n3
	n3.next = n4
	n4.next = n5
	n5.next = n6
	n6.next = n7
	n7.next = n4
	got := DetectLoopInLLUsingDict(h)

	if got != true {
		t.Error("ahs error")
	}
}

func Test_DetectLoopInLLUsingNextPointerSwapping(t *testing.T) {

	h := new(NodeSingle)
	h.data = 1

	n2 := new(NodeSingle)
	n2.data = 2

	n3 := new(NodeSingle)
	n3.data = 3
	n4 := new(NodeSingle)
	n4.data = 4

	n5 := new(NodeSingle)
	n5.data = 5
	n6 := new(NodeSingle)
	n6.data = 6
	n7 := new(NodeSingle)
	n7.data = 6

	h.next = n2
	n2.next = n3
	n3.next = n4
	n4.next = n5
	n5.next = n6
	n6.next = n7
	n7.next = n4
	got := DetectLoopInLLUsingPointerSwapping(h)

	if got != true {
		t.Error("ahs error")
	}
}
