package tree

import "testing"

func Test_isChildrenSum(t *testing.T) {

	got := isChildrenSum(getChildSumNode1())

	if got != true {
		t.Error("has error")
	}

	got = isChildrenSum(getChildSumNode2())

	if got != true {
		t.Error("has error")
	}

	got = isChildrenSum(getChildSumNode3())

	if got != true {
		t.Error("has error")
	}

	got = isChildrenSum(getChildSumNode4())

	if got != false {
		t.Error("has error")
	}
}

func getChildSumNode1() *Node[int] {
	n := &Node[int]{data: 20}
	n.left = &Node[int]{data: 8}
	n.left.left = &Node[int]{data: 3}
	n.left.right = &Node[int]{data: 5}
	n.right = &Node[int]{data: 12}

	return n
}

func getChildSumNode2() *Node[int] {
	n := &Node[int]{data: 5}

	return n
}
func getChildSumNode3() *Node[int] {
	n := &Node[int]{data: 10}
	n.left = &Node[int]{data: 8}
	n.right = &Node[int]{data: 2}
	n.right.right = &Node[int]{data: 2}

	return n
}
func getChildSumNode4() *Node[int] {
	n := &Node[int]{data: 3}
	n.left = &Node[int]{data: 1}
	n.right = &Node[int]{data: 2}
	n.right.left = &Node[int]{data: 1}
	n.right.right = &Node[int]{data: 2}

	return n
}
