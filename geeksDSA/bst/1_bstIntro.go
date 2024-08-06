package bst

type TestBst struct {
	root *Node[int]
}

func (t *TestBst) Insert(val int) {

	if t.root == nil {
		t.root = &Node[int]{Value: val}
		return
	}

	curr := t.root

	for curr != nil {

		if val < curr.Value {

			if curr.Left == nil {
				curr.Left = &Node[int]{Value: val}
				break
			}
			curr = curr.Left

		} else {

			if curr.Right == nil {
				curr.Right = &Node[int]{Value: val}
				break
			}
			curr = curr.Right

		}
	}

}

func (t *TestBst) Search(val int) bool {
	curr := t.root
	for curr != nil {
		if curr.Value == val {
			return true
		}
		if val < curr.Value {
			curr = curr.Left
		} else {
			curr = curr.Right
		}
	}

	return false
}
