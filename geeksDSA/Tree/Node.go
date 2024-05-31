package tree

type Node[T any] struct {
	left  *Node[T]
	right *Node[T]
	data  T
}
