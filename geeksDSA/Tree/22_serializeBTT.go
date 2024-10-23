package tree

var arr []int = []int{}

func SerializeBinaryTree(n *Node[int]) {

	if n == nil {
		arr = append(arr, -1)
		return
	}
	arr = append(arr, n.data)
	SerializeBinaryTree(n.left)
	SerializeBinaryTree(n.right)
}

var serializeIndex int

func DeSerializeBinaryTree() *Node[int] {

	if serializeIndex > len(arr)-1 || arr[serializeIndex] == -1 {
		serializeIndex++
		return nil
	}

	node := &Node[int]{data: arr[serializeIndex]}
	serializeIndex++
	node.left = DeSerializeBinaryTree()
	node.right = DeSerializeBinaryTree()

	return node
}
