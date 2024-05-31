package tree

func isChildrenSum(n *Node[int]) bool {

	if n == nil || (n.left == nil && n.right == nil) {
		return true
	}
	if getDefaultValueofNode(n.left)+getDefaultValueofNode(n.right) != n.data {
		return false
	}
	return isChildrenSum(n.left) && isChildrenSum(n.right)
}

func getDefaultValueofNode(n *Node[int]) int {
	if n == nil {
		return 0
	}

	return n.data
}
