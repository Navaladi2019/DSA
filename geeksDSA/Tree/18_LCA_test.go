package tree

import "testing"

func Test_FindLCA_1(t *testing.T) {

	got := FindLCA_1(getLCANode1(), 60, 70)

	if got.data != 30 {
		t.Error("has error")
	}
	got = FindLCA_1(getLCANode1(), 20, 80)

	if got.data != 10 {
		t.Error("has error")
	}
	got = FindLCA_1(getLCANode1(), 80, 30)

	if got.data != 30 {
		t.Error("has error")
	}

	got = FindLCA_1(getLCANode1(), 70, 80)

	if got.data != 50 {
		t.Error("has error")
	}

}

func getLCANode1() *Node[int] {

	n := Node[int]{data: 10}

	n.left = &Node[int]{data: 20}
	n.right = &Node[int]{data: 30}
	n.right.left = &Node[int]{data: 40}
	n.right.left.left = &Node[int]{data: 60}
	n.right.right = &Node[int]{data: 50}
	n.right.right.left = &Node[int]{data: 70}
	n.right.right.right = &Node[int]{data: 80}

	return &n

}
