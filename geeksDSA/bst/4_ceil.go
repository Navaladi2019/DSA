package bst

func Ceil(root *Node[int], val int) *int {

	if root == nil {
		return nil
	} else if root.Value == val {
		return &root.Value
	} else if root.Value > val {
		res := Ceil(root.Left, val)
		if res != nil && *res < root.Value {
			return res
		}
		return &root.Value
	} else {
		return Ceil(root.Right, val)
	}
}

func Ceil_1(root *Node[int], val int) *int {

	var res *int
	curr := root
	for curr != nil {

		if curr.Value == val {
			res = &curr.Value
			break
		} else if curr.Value < val {
			curr = curr.Right
		} else {
			if res == nil || *(res) > curr.Value {
				res = &curr.Value
			}
			curr = curr.Left
		}

	}
	return res
}
