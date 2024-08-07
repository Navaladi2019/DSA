package bst

var first, second, prev *Node[int]

func FixBSTWithTwoNodesSwapped(root *Node[int]) {

	if root == nil {
		return
	}
	FixBSTWithTwoNodesSwapped(root.Left)

	if prev != nil && prev.Value > root.Value {

		if first == nil {
			first = prev
		}
		second = root
	}

	prev = root
	FixBSTWithTwoNodesSwapped(root.Right)
}

func FixBSTWithTwoNodesSwappedParent(root *Node[int]) {
	FixBSTWithTwoNodesSwapped(root)
	if first != nil && second != nil {
		first.Value, second.Value = second.Value, first.Value
	}
}
