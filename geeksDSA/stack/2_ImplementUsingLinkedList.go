package stack

type Node[T any] struct {
	value T
	next  *Node[T]
}

type LLStack[T any] struct {
	head *Node[T]
	size int
}

func (s *LLStack[T]) Cap() int {
	return s.size
}

func (s *LLStack[T]) Peek() (T, bool) {

	var val T
	ok := false
	pointer := s.head

	if pointer != nil {
		val = pointer.value
		ok = true
	}
	return val, ok
}

func (s *LLStack[T]) Push(val T) {

	n := &Node[T]{value: val}
	n.next = s.head
	s.head = n
	s.size++
}

func (s *LLStack[T]) Pop() (T, bool) {

	var val T
	ok := false
	pointer := s.head

	if pointer != nil {
		val = pointer.value
		ok = true
		s.head = s.head.next
		s.size--
	}
	return val, ok
}

func (s *LLStack[T]) isEmpty() bool {
	return s.size == 0
}
