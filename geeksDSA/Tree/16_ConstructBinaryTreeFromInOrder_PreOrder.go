package tree

var prein int

func ConstructTree(inOr []int, pre []int, inS int, inE int) *Node[int] {

	if inS > inE {
		return nil
	}
	curr := &Node[int]{data: pre[prein]}
	prein++
	rootindex := inS
	for i := inS; i <= inE; i++ {
		if inOr[i] == curr.data {
			rootindex = i
			break
		}
	}

	curr.left = ConstructTree(inOr, pre, inS, rootindex-1)
	curr.right = ConstructTree(inOr, pre, rootindex+1, inE)

	return curr
}
