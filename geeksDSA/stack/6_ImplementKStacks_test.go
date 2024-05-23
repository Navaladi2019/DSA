package stack

import (
	"fmt"
	"testing"
)

func Test_KStacksInArray(t *testing.T) {

	ks := kStacks[string]{}

	ks.Init(10, 3)

	ks.Push(1, "k1")
	ks.Push(2, "k2")
	ks.Push(3, "k3")
	ks.Push(1, "k12")
	ks.Push(2, "k22")

	fmt.Println(ks.Peek(1))
	fmt.Println(ks.Peek(2))
	fmt.Println(ks.Pop(1))
	ks.Push(2, "k23")
	fmt.Println(ks.Pop(1))
	fmt.Println(ks.Pop(2))
	fmt.Println(ks.Peek(3))
}
