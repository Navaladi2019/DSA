package tree

import (
	"fmt"
	"testing"
)

func Test_PrintKNode(t *testing.T) {

	PrintKNode(getNode1(), 2)
	fmt.Println("****")
	PrintKNode(getNode2(), 3)
	fmt.Println("****")
	PrintKNode(getNode3(), 1)
	fmt.Println("****")

}

func getNode1() *Node[int] {
	r := &Node[int]{data: 10}
	r.left = &Node[int]{data: 20}
	r.left.left = &Node[int]{data: 40}
	r.left.right = &Node[int]{data: 50}
	r.right = &Node[int]{data: 30}
	r.right.right = &Node[int]{data: 70}
	return r
}

func getNode2() *Node[int] {
	r := &Node[int]{data: 10}
	r.left = &Node[int]{data: 6}

	r.right = &Node[int]{data: 8}
	r.right.right = &Node[int]{data: 7}
	r.right.right.left = &Node[int]{data: 11}
	r.right.right.right = &Node[int]{data: 12}
	return r
}

func getNode3() *Node[int] {
	r := &Node[int]{data: 10}
	r.left = &Node[int]{data: 20}
	r.left.left = &Node[int]{data: 30}
	r.left.left.left = &Node[int]{data: 40}
	return r
}
