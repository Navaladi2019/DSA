package tree

import "math"

func FindMax(n *Node[int]) int {

	if n == nil {
		return math.MinInt
	}

	return max(n.data, FindMax(n.left), FindMax(n.right))
}
