package tree

import (
	"fmt"
	"testing"
)

func Test_InOrderTraversal2(t *testing.T) {
	n := GetNodeInOrderTraversal()
	TraversalInOrder(&n)
}
func Test_InOrderTraversalIterative(t *testing.T) {
	n := GetNodeInOrderTraversal()
	InOrderTraversalIterative(&n)
}

func Test_PreOrderTraversalIterative(t *testing.T) {
	n := GetNodeInOrderTraversal()
	PreOrderTraversalIterative(&n)
	fmt.Println("Space Efficient")
	PreOrderTraversalIterative_SpaceEfficient(&n)
}
func Test_PostOrderTraversalIterative(t *testing.T) {
	n := GetNodeInOrderTraversal()
	PostOrderTraversalIterative(&n)
}
