package tree

import "testing"

func GetNodeInOrderTraversal() Node[int] {
	root := Node[int]{data: 10}
	root.left = &Node[int]{data: 20}
	root.right = &Node[int]{data: 30}
	root.right.left = &Node[int]{data: 40}
	root.right.right = &Node[int]{data: 50}
	return root
}

func Test_InOrderTraversal(t *testing.T) {
	n := GetNodeInOrderTraversal()
	TraversalInOrder(&n)
}

func Test_PreOrderTraversal(t *testing.T) {
	n := GetNodeInOrderTraversal()
	TraversalPreOrder(&n)
}

func Test_PostOrderTraversal(t *testing.T) {
	n := GetNodeInOrderTraversal()
	TraversalPostOrder(&n)
}
