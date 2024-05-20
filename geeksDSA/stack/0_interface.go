package stack

type StackInterface[T comparable] interface {
	Peek() (T, bool)

	Pop() (T, bool)

	isEmpty() bool

	Size() int
}
