package tree

var resBurn int = -1

func FindBurnTime(root *Node[int], leaf int) int {

	BurnTree(root, leaf)
	return res
}

func BurnTree(root *Node[int], leaf int) (distance int, height int) {

	if root == nil {
		return -1, 0
	}

	if root.data == leaf {
		return 0, 0
	}

	leftdist, leftHeight := BurnTree(root.left, leaf)
	rightdist, rightHeight := BurnTree(root.right, leaf)

	if leftdist != -1 {
		leftdist = leftdist + 1
		resBurn = max(resBurn, leftdist+rightHeight)
	}
	if rightdist != -1 {
		rightdist = rightdist + 1
		resBurn = max(resBurn, rightdist+leftHeight)
	}

	return max(leftdist, rightdist), max(leftHeight, rightHeight) + 1

}
