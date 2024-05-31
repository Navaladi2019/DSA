package tree

import (
	"fmt"
	"testing"
)

func Test_LevelOrderTraversal_Efficient(t *testing.T) {
	LevelOrderTraversalEfficient_NewLine_2(getNodeIn1())
	fmt.Println("****")
	LevelOrderTraversalEfficient_NewLine_2(getNodeIn2())
	fmt.Println("****")
	LevelOrderTraversalEfficient_NewLine_2(getNodeIn3())
	fmt.Println("****")
}
func Test_LevelOrderTraversal(t *testing.T) {
	LevelOrderTraversal_Naive(getNodeIn1())
	fmt.Println("****")
	LevelOrderTraversal_Naive(getNodeIn2())
	fmt.Println("****")
	LevelOrderTraversal_Naive(getNodeIn3())
	fmt.Println("****")
}

func getNodeIn1() *Node[int] {
	r := &Node[int]{data: 10}
	r.left = &Node[int]{data: 20}
	r.left.left = &Node[int]{data: 8}
	r.left.right = &Node[int]{data: 7}
	r.left.right.left = &Node[int]{data: 9}
	r.left.right.right = &Node[int]{data: 15}
	r.right = &Node[int]{data: 30}
	r.right.right = &Node[int]{data: 6}
	return r
}

func getNodeIn2() *Node[int] {
	r := &Node[int]{data: 8}
	r.right = &Node[int]{data: 6}
	r.right.right = &Node[int]{data: 4}
	r.right.left = &Node[int]{data: 2}
	r.right.right.left = &Node[int]{data: 3}
	return r
}

func getNodeIn3() *Node[int] {
	r := &Node[int]{data: 3}
	r.left = &Node[int]{data: 4}
	r.left.left = &Node[int]{data: 5}
	return r
}
