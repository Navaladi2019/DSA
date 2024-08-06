// - Self balancing trees keeps the height in check to o logn so operations on BST does not increase to o(n)

// - we can attain self balancing trees while constructing the BST like

// 1,2,3,4,5,6,7,8,9,10 by converting this order to taking middle

package bst

import "testing"

func Test_ConstructSelfBalancedBinarttree(t *testing.T) {

	tv := ConstructSelfBalancingBST([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})

	t.Log(tv)
}
