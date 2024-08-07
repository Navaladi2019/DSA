package bst

import "testing"

func Test_CheckSumWIthTwoPairs(t *testing.T) {

	root := &Node[int]{Value: 10}
	root.Left = &Node[int]{Value: 5}
	root.Right = &Node[int]{Value: 20}
	root.Right.Left = &Node[int]{Value: 16}
	root.Right.Right = &Node[int]{Value: 40}

	m := make(map[int]struct{})

	got := CheckSumWIthTwoPairs(root, 21, m)

	if got != true {
		t.Error("Check sum failed")
	}
}
