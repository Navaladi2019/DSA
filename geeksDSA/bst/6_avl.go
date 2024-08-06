package bst

//https://www.geeksforgeeks.org/insertion-in-an-avl-tree/
// The idea is to use recursive BST insert, after insertion,
// we get pointers to all ancestors one by one in a bottom-up manner.
//  So we don’t need a parent pointer to travel up.
//  The recursive code itself travels up and visits all the ancestors of the newly inserted node.

//Balance Factor = left subtree height – right subtree height

type AVL[T Numbers] struct {
	BST[T]
}

func (a *AVL[T]) Insert(val T) {
	a.root = InsertRecursiveAVL(a.root, val)
}
func InsertRecursiveAVL[T Numbers](root *Node[T], val T) *Node[T] {
	if root == nil {
		return &Node[T]{Value: val}
	} else if root.Value > val {
		root.Left = InsertRecursiveAVL(root.Left, val)

	} else if root.Value < val {
		root.Right = InsertRecursiveAVL(root.Right, val)
	}

	root.height = 1 + max(height(root.Left), height(root.Right))

	bf := getBalanceFactor(root)

	// left left
	if bf > 1 && val < root.Left.Value {
		return RotateRight(root)
	}

	// right right
	if bf < -1 && val > root.Right.Value {
		return RotateLeft(root)
	}

	// left right
	if bf > 1 && val > root.Right.Value {
		root.Left = RotateLeft(root.Left)
		return RotateRight(root)
	}

	// right left
	if bf < -1 && val < root.Right.Value {
		root.Right = RotateRight(root.Right)
		return RotateLeft(root)
	}

	return root
}

func RotateLeft[T Numbers](root *Node[T]) *Node[T] {
	if root == nil || root.Right == nil {
		return root
	}
	nextRoot := root.Right
	temp := nextRoot.Left
	nextRoot.Left = root
	root.Right = temp

	root.height = max(height(root.Left), height(root.Right)) + 1
	nextRoot.height = max(height(root.Left), height(root.Right)) + 1

	return nextRoot
}

func RotateRight[T Numbers](root *Node[T]) *Node[T] {
	if root == nil || root.Left == nil {
		return root
	}
	nextRoot := root.Left
	temp := nextRoot.Right
	nextRoot.Right = root
	root.Left = temp

	root.height = max(height(root.Left), height(root.Right)) + 1
	nextRoot.height = max(height(root.Left), height(root.Right)) + 1
	return nextRoot
}

func height[T Numbers](root *Node[T]) int {
	if root == nil {
		return 0
	}
	return root.height
}

func getBalanceFactor[T Numbers](root *Node[T]) int {

	if root == nil {
		return 0
	}
	return height(root.Left) - height(root.Right)
}
func DeleteAVL[T Numbers](root *Node[T], val T) *Node[T] {
	if root == nil {
		return nil
	} else if root.Value > val {
		root.Left = DeleteAVL(root.Left, val)

	} else if root.Value < val {
		root.Right = DeleteAVL(root.Right, val)
	} else {
		if root.Right == nil && root.Left != nil {
			return root.Left
		} else if root.Left == nil && root.Right != nil {
			return root.Right
		} else {
			succNode := getSuccessorOfDelete(root)
			root.Value = succNode.Value
			root.Right = DeleteAVL(root.Right, succNode.Value)
		}
	}
	root.height = 1 + max(height(root.Left), height(root.Right))

	bf := getBalanceFactor(root)

	//Left Left
	if bf > 1 && getBalanceFactor(root.Left) >= 0 {
		return RotateRight(root)
	}

	//Left Right
	if bf > 1 && getBalanceFactor(root.Left) < 0 {
		root.Left = RotateLeft(root.Left)
		return RotateRight(root)

	}

	//Right Right
	if bf < -1 && getBalanceFactor(root.Right) <= 0 {
		return RotateLeft(root)
	}
	//Right Left
	if bf < -1 && getBalanceFactor(root.Right) > 0 {
		root.Right = RotateRight(root.Right)
		return RotateLeft(root)
	}
	return root
}
