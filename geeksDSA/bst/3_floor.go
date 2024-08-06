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
