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
