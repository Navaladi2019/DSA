package bst

func Floor(root *Node[int], val int) *int {

	if root == nil {
		return nil
	} else if root.Value == val {
		return &root.Value
	} else if root.Value > val {
		return Floor(root.Left, val)
	} else {
		res := Floor(root.Right, val)
		if res != nil && *res > root.Value {
			return res
		}
		return &root.Value
	}
}

func Floor_iterative(root *Node[int], val int) *int {

	var res *int

	curr := root

	for curr != nil {

		if curr.Value == val {
			return &val
		} else if curr.Value < val {

			if res == nil {
				res = &curr.Value
			} else {

				if *res < curr.Value {
					res = &curr.Value
				}
			}

			curr = curr.Right

		} else {
			curr = curr.Left
		}
	}

	return res
}

func Floor_iterative1(root *Node[int], val int) *int {
	var res *int
	curr := root
	for curr != nil {
		if curr.Value == val {
			return &curr.Value
		} else if curr.Value < val {
			res = &curr.Value
			curr = curr.Right
		} else {
			curr = curr.Left
		}
	}
	return res
}

func FloorRecursion(root *Node[int], val int) *Node[int] {

	if root == nil {
		return nil
	}

	if root.Value == val {
		return root
	} else if root.Value > val {
		return FloorRecursion(root.Left, val)
	} else {

		if root.Right == nil {
			return root
		}

		nextPossible := FloorRecursion(root.Right, val)

		if nextPossible != nil && nextPossible.Value > root.Value {
			return nextPossible
		}
		return root
	}
}
