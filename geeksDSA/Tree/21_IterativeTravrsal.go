package tree

import (
	"fmt"

	"github.com/Navaladi2019/GoRefresher/geeksDSA/stack"
)

func InOrderTraversalIterative(root *Node[int]) {

	sta := stack.ArrStack[*Node[int]]{}

	curr := root

	for curr != nil || !sta.IsEmpty() {

		for curr != nil {
			sta.Push(curr)
			curr = curr.left
		}

		curr, _ = sta.Pop()
		fmt.Println(curr.data)
		curr = curr.right

	}

}

func PreOrderTraversalIterative(root *Node[int]) {

	sta := stack.ArrStack[*Node[int]]{}
	curr := root

	for curr != nil || !sta.IsEmpty() {
		if curr != nil {
			fmt.Println(curr.data)
			sta.Push(curr)
			curr = curr.left
		} else {
			if tempcur, ok := sta.Pop(); ok {
				curr = tempcur.right
			}
		}
	}
}
func PreOrderTraversalIterative_SpaceEfficient(root *Node[int]) {

	sta := stack.ArrStack[*Node[int]]{}

	var curr *Node[int] = root
	for curr != nil || !sta.IsEmpty() {
		if curr == nil {
			curr, _ = sta.Pop()
		}

		fmt.Println(curr.data)

		if curr.right != nil {
			sta.Push(curr.right)
		}

		curr = curr.left

	}

}
func PostOrderTraversalIterative(root *Node[int]) {

	sta := stack.ArrStack[*Node[int]]{}

	curr := root

	for curr != nil || !sta.IsEmpty() {

		if curr == nil {

			tempCurr, _ := sta.Pop()
			if tempCurr.right != nil {
				sta.Push(&Node[int]{data: tempCurr.data})
				curr = tempCurr.right
			} else {
				fmt.Println(tempCurr.data)
				continue
			}
		}

		if curr.left == nil && curr.right == nil {
			fmt.Println(curr.data)
			curr = nil
			continue
		}

		if curr.left != nil {
			sta.Push(curr)
			curr = curr.left
		}

	}

}
