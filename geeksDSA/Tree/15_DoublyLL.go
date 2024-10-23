package tree

var prev *Node[int]

func BTToDLL(n *Node[int]) *Node[int] {
	if n == nil {
		return n
	}
	head := BTToDLL(n.left)
	if prev == nil {
		prev = n
	} else {
		n.left = prev
		prev.right = n
	}
	prev = n
	BTToDLL(n.right)
	return head
}

var lastnode *Node[int]

func BTTOLL(n *Node[int]) *Node[int] {

	tempright := n.right
	if n == nil {
		return nil
	}

	head := BTTOLL(n.left)

	if lastnode == nil {
		lastnode = n
	} else {
		n.left = head
		lastnode.right = n
	}
	BTTOLL(tempright)

	return head
}
