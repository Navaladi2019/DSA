package tree

import "fmt"

func TraversalPreOrder(n *Node[int]) {
	if n == nil {
		return
	}
	fmt.Println(n.data)
	TraversalPreOrder(n.left)

	TraversalPreOrder(n.right)
}
