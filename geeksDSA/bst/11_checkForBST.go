package bst

func CheckForBst(n *Node[int], maxval int, minval int) bool {
	if n == nil {
		return true
	}

	if n.Value < minval || n.Value > maxval {
		return false
	}

	return CheckForBst(n.Left, n.Value, minval) && CheckForBst(n.Right, maxval, n.Value)
}
