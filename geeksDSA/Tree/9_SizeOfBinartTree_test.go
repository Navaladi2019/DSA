package tree

import (
	"fmt"
	"testing"
)

func Test_SizeOfBinaryTree(t *testing.T) {

	fmt.Println(SizeOfBinaryTree_Recursive(getNodeIn1()))
	fmt.Println(SizeOfBinaryTree_Recursive(getNodeIn2()))
	fmt.Println(SizeOfBinaryTree_Recursive(getNodeIn3()))
}
