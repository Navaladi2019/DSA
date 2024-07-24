package tree

// here it takes o(n^2)
func FindDiaOfBinaryTree(n *Node[int]) int {
	if n == nil {
		return 0
	}
	d1 := 1 + height(n.left) + height(n.right)
	d2 := FindDiaOfBinaryTree(n.left)
	d3 := FindDiaOfBinaryTree(n.right)

	return max(d1, d2, d3)
}

var res int

func FindDiaOfBinaryTree_Efficient(n *Node[int]) int {
	if n == nil {
		return 0
	}

	d1 := FindDiaOfBinaryTree_Efficient(n.left)
	d2 := FindDiaOfBinaryTree_Efficient(n.right)
	res = max(res, d1+d2+1)
	return 1 + max(d1, d2)
}
