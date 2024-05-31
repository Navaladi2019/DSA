package tree

func SizeOfBinaryTree_Recursive(n *Node[int]) int {

	if n == nil {
		return 0
	}

	return 1 + SizeOfBinaryTree_Recursive(n.left) + SizeOfBinaryTree_Recursive(n.right)
}
