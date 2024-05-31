package tree

import "fmt"

func TraversalInOrder(n *Node[int]) {
	if n == nil {
		return
	}
	TraversalInOrder(n.left)
	fmt.Println(n.data)
	TraversalInOrder(n.right)
}
