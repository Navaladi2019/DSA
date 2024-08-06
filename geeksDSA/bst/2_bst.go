package bst

type Numbers interface {
	~int | ~uint8 | ~float64
}

type BST[T Numbers] struct {
	root *Node[T]
}

func (bt *BST[T]) Insert(val T) {
	bt.root = InsertRecursive(bt.root, val)

}

func (bt *BST[T]) Delete(val T) {
	bt.root = DeleteRecursive(bt.root, val)

}

func InsertRecursive[T Numbers](root *Node[T], val T) *Node[T] {
	if root == nil {
		return &Node[T]{Value: val}
	} else if root.Value > val {
		root.Left = InsertRecursive(root.Left, val)

	} else if root.Value < val {
		root.Right = InsertRecursive(root.Right, val)
	}
	return root
}

func DeleteRecursive[T Numbers](root *Node[T], val T) *Node[T] {
	if root == nil {
		return nil
	} else if root.Value > val {
		root.Left = DeleteRecursive(root.Left, val)

	} else if root.Value < val {
		root.Right = DeleteRecursive(root.Right, val)
	} else {
		if root.Right == nil && root.Left != nil {
			return root.Left
		} else if root.Left == nil && root.Right != nil {
			return root.Right
		} else {
			succNode := getSuccessorOfDelete(root)
			root.Value = succNode.Value
			root.Right = DeleteRecursive(root.Right, succNode.Value)
		}
	}
	return root
}

func getSuccessorOfDelete[T Numbers](root *Node[T]) *Node[T] {
	curr := root.Right

	for curr != nil && curr.Left != nil {
		curr = curr.Left
	}
	return curr
}

func InsertIterative[T Numbers](root *Node[T], val T) *Node[T] {

	temp := &Node[T]{Value: val}

	curr := root
	LastCurr := curr
	for curr != nil {
		LastCurr = curr
		if val < curr.Value {
			curr = curr.Left
		} else if val > curr.Value {
			curr = curr.Right
		} else {
			return root
		}
	}
	if LastCurr == nil {
		return temp
	}

	if val < LastCurr.Value {
		LastCurr.Left = temp
	} else {
		LastCurr.Right = temp
	}

	return root
}

func Delete[T Numbers](root *Node[T], val T) *Node[T] {

	curr := root

	var prev *Node[T]
	isPrevLeft := false
	for curr != nil {

		if curr.Value == val {

			if prev == nil {
				return InsertNodeAndDelete(curr)
			} else {
				if isPrevLeft {
					prev.Left = InsertNodeAndDelete(curr)
				} else {
					prev.Right = InsertNodeAndDelete(curr)
				}
				break
			}

		} else {
			prev = curr
			if curr.Value < val {
				curr = curr.Right
				isPrevLeft = false
			} else {
				curr = curr.Left
				isPrevLeft = true
			}
		}
	}

	return root
}

func InsertNodeAndDelete[T Numbers](curr *Node[T]) *Node[T] {

	if curr.Right == nil && curr.Left != nil {
		return curr.Left
	} else if curr.Left == nil && curr.Right != nil {
		return curr.Right
	} else {
		InsertAtLeftLeaf(curr.Right, curr.Left)
	}
	return curr.Right
}

func InsertAtLeftLeaf[T any](root *Node[T], toInsert *Node[T]) {

	curr := root

	for curr != nil {
		if curr.Left == nil {
			curr.Left = toInsert
			break
		}
		curr = curr.Left
	}

}
