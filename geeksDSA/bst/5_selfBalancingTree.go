// - Self balancing trees keeps the height in check to o logn so operations on BST does not increase to o(n)

// - we can attain self balancing trees while constructing the BST like

// 1,2,3,4,5,6,7,8,9,10 by converting this order to taking middle

package bst

import "slices"

//1,2,3,4,5,6,7,8,9,10
func ConstructSelfBalancingBST(arr []int) *Node[int] {

	if len(arr) == 0 {
		return nil
	}

	slices.Sort(arr)
	mid := len(arr) / 2
	node := &Node[int]{Value: arr[mid]}
	node.Left = ConstructSelfBalancingBST(arr[:mid])
	node.Right = ConstructSelfBalancingBST(arr[mid+1:])
	return node

}
