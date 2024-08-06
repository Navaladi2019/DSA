package bst

import (
	"fmt"
	"math"
)

func Ceil_Arr(arr []int) {

	avl := AVL[int]{}

	avl.Insert(arr[0])
	fmt.Print(-1)

	for i := 1; i < len(arr); i++ {

		res := math.MaxInt
		curr := avl.root
		for curr != nil {
			if curr.Value < arr[i] {
				curr = curr.Right
			} else {
				res = min(curr.Value, res)
				curr = curr.Left
			}
		}
		avl.Insert(arr[i])

		if res == math.MaxInt {
			fmt.Print(-1, " ")
		} else {
			fmt.Print(res, " ")
		}

	}
}
