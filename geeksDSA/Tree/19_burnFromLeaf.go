package tree

func FindBurnTime(root *Node[int], leaf int) int {
	res := -1
	BurnTree(root, leaf, &res)
	return res
}

func BurnTree(root *Node[int], leaf int, res *int) (distance int, height int) {

	if root == nil {
		return -1, 0
	}

	if root.data == leaf {
		return 0, 0
	}

	leftdist, leftHeight := BurnTree(root.left, leaf, res)
	rightdist, rightHeight := BurnTree(root.right, leaf, res)

	if leftdist != -1 {
		leftdist = leftdist + 1
		*res = max(*res, leftdist+rightHeight)
	}
	if rightdist != -1 {
		rightdist = rightdist + 1
		*res = max(*res, rightdist+leftHeight)
	}

	return max(leftdist, rightdist), max(leftHeight, rightHeight) + 1

}
