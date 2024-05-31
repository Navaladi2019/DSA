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

	lv := IsBalancedTree_Efficient(n.left)

	if lv == -1 {
		return -1
	}
	rv := IsBalancedTree_Efficient(n.right)

	if rv == -1 {
		return -1
	}

	if math.Abs(float64(lv-rv)) > 1 {
		return -1
	}

	return max(lv, rv)
}
