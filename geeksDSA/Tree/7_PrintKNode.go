package tree

import "fmt"

func PrintKNode(n *Node[int], k int) {
	if n == nil {
		return
	}
	if k == 0 {
		fmt.Println(n.data)
		return
	}
	PrintKNode(n.left, k-1)
	PrintKNode(n.right, k-1)
}
