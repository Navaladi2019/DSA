package tree

import "math"

func CountNodesInCompleteBinaryTree(root *Node[int]) int {

	lh, rh := 0, 0

	curr := root

	for curr != nil {
		lh++
		curr = curr.left
	}

	curr = root

	for curr != nil {
		rh++
		curr = curr.right
	}

	if lh == rh {
		return int(math.Pow(2, float64(lh))) - 1
	}

	return 1 + CountNodesInCompleteBinaryTree(root.left) + CountNodesInCompleteBinaryTree(root.right)
}
