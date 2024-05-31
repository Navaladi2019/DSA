package tree

func height(n *Node[int]) int {
	if n == nil {
		return 0
	}
	return max(height(n.left), height(n.right)) + 1
}
