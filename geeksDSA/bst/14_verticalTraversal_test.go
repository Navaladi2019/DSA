package bst

import "testing"

func Test_VerticalTraversal(t *testing.T) {
	root := &Node[int]{Value: 10}
	root.Left = &Node[int]{Value: 15}
	root.Left.Left = &Node[int]{Value: 35}
	root.Left.Left.Left = &Node[int]{Value: 40}
	root.Left.Right = &Node[int]{Value: 20}
	root.Left.Right.Right = &Node[int]{Value: 75}
	root.Left.Right.Right.Right = &Node[int]{Value: 80}
	root.Right = &Node[int]{Value: 25}

	FindVerticalTraversal(root)
}
