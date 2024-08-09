package bst

import "fmt"

// for the next problem top view of binary tree we have to use level order traversal
// and not inOrder traversal, if i use inorder traversal then it will give preference to left and not the top element
// so we have to use level order traversal and map with horizontal distance

var maxinv, mininv int

func VerticalTraversal(root *Node[int], order int, m map[int][]int) {
	if root == nil {
		return
	}

	if val, ok := m[order]; ok {
		m[order] = append(val, root.Value)
	} else {
		m[order] = []int{root.Value}
	}
	maxinv = max(maxinv, order)
	mininv = min(mininv, order)
	VerticalTraversal(root.Left, order-1, m)
	VerticalTraversal(root.Right, order+1, m)
}

func FindVerticalTraversal(root *Node[int]) {
	m := make(map[int][]int)
	VerticalTraversal(root, 0, m)
	for mininv <= maxinv {
		fmt.Println(m[mininv])
		mininv++
	}
}
