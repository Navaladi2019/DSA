package bst

type Node[T any] struct {
	Value  T
	height int
	Left   *Node[T]
	Right  *Node[T]
}
