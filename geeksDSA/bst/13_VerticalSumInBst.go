package bst

import "fmt"

var maxin, minin int

func FindverticalSum(root *Node[int], currentPosition int, m map[int]int) {

	if root == nil {
		return
	}

	m[currentPosition] += root.Value

	maxin = max(maxin, currentPosition)
	minin = min(minin, currentPosition)

	FindverticalSum(root.Left, currentPosition-1, m)
	FindverticalSum(root.Right, currentPosition+1, m)

}

func FindverticalSumParent(root *Node[int]) {
	m := make(map[int]int)
	FindverticalSum(root, 0, m)
	for minin <= maxin {
		fmt.Println(m[minin])
		minin++
	}
}
