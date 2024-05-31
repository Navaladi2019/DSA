package tree

import "fmt"

func PrintLeftNode(n *Node[int], isLeft bool) {

	if n == nil {
		return
	}

	if isLeft {
		fmt.Println(n.data)
	}

	PrintLeftNode(n.left, true)
	PrintLeftNode(n.right, false)
}

var maxlevel int

func PrintLeftMostNode(n *Node[int], level int) {

	if n == nil {
		return
	}

	if level < maxlevel {
		fmt.Println(n.data)
	}

	PrintLeftNode(n.left, true)
	PrintLeftNode(n.right, false)
}
