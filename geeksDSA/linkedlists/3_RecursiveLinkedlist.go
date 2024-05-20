package linkedlists

import "fmt"

func RecursiveLinkedList(n *NodeSingle) {

	if n == nil {
		return
	}
	fmt.Println(n.data)
	RecursiveLinkedList(n.next)

}
