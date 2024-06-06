package tree

import "fmt"

func TraversalPostOrder(n *Node[int]) {
	if n == nil {
		return
	}

	TraversalPostOrder(n.left)
	TraversalPostOrder(n.right)
	fmt.Println(n.data)

}
