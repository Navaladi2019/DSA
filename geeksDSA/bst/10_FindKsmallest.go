package bst

import "fmt"

var KthSmallestCount int

func FindKthSmallest(n *Node[int], val int) {

	if n.Left != nil {
		FindKthSmallest(n.Left, val)
	}

	KthSmallestCount++

	if KthSmallestCount == val {
		fmt.Println(n.Value)
		return
	}
	if n.Right != nil {
		FindKthSmallest(n.Right, val)
	}
}

type AugmentedBST[T Numbers] struct {
	left      *AugmentedBST[T]
	right     *AugmentedBST[T]
	LeftCount int
}
