package tree

import "math"

// o(n^2)
func IsBalancedTree(n *Node[int]) bool {
	if n == nil {
		return true
	}
	dif := height(n.left) - height(n.right)
	if math.Abs(float64(dif)) <= 1 && IsBalancedTree(n.left) && IsBalancedTree(n.right) {
		return true
	}
	return false
}

func IsBalancedTree_Efficient(n *Node[int]) int {
	if n == nil {
		return 0
	}

	leftHeight := IsBalancedTree_Efficient(n.left)

	if leftHeight == -1 {
		return -1
	}
	rightHeight := IsBalancedTree_Efficient(n.right)

	if rightHeight == -1 {
		return -1
	}

	if math.Abs(float64(leftHeight-rightHeight)) > 1 {
		return -1
	}

	return max(leftHeight, rightHeight)
}
