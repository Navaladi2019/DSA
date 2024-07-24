package tree

var prev *Node[int]

func BTToDLL(n *Node[int]) *Node[int] {
	if n == nil {
		return n
	}
	head := BTToDLL(n.left)
	if head == nil {
		head = n
	} else {
		n.left = prev
		prev.right = n
	}
	prev = n
	BTToDLL(n.right)
	return head
}
