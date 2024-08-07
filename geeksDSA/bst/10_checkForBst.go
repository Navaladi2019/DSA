package bst

import (
	"fmt"
	"math"
)

func CheckForBST(root *Node[int], low int, high int) bool {
	if root == nil {
		return true
	}
	if root.Value < low || root.Value > high {
		return false
	}
	left := CheckForBST(root.Left, low, root.Value)
	right := CheckForBST(root.Right, root.Value, high)
	return left && right
}

func IsBst(root *Node[int]) {
	res := CheckForBST(root, math.MinInt, math.MaxInt)
	fmt.Println(res)
}
