package stack

type kStacks[T any] struct {
	arr []T
	/*Number of stacks*/
	k       []int
	history []int
	freetop int
}

func (s *kStacks[T]) Init(cap int, k int) {

	s.arr = make([]T, cap)
	s.k = make([]int, k)
	s.history = make([]int, cap)

	for i := 0; i < len(s.k); i++ {

		s.k[i] = -1
	}

	for i := 0; i < len(s.history)-1; i++ {
		s.history[i] = i + 1
	}
	s.history[len(s.history)-1] = -1

	s.freetop = 0
}

func (s *kStacks[T]) Push(k int, val T) bool {

	if s.freetop == -1 {
		return false
	}
	freetop := s.freetop
	s.freetop = s.history[freetop]
	s.arr[freetop] = val
	s.history[freetop] = s.k[k-1]
	s.k[k-1] = freetop

	return true
}

func (s *kStacks[T]) Pop(k int) (val T, ok bool) {

	i := s.k[k-1]

	if i == -1 {
		return
	}

	ok = true
	val = s.arr[i]
	s.k[k-1] = s.history[i]
	s.history[i] = s.freetop
	s.freetop = i
	return
}

func (s *kStacks[T]) Peek(k int) (val T, ok bool) {

	i := s.k[k-1]

	if i == -1 {
		return
	}
	val = s.arr[i]
	ok = true
	return
}

func (s *kStacks[T]) isEmpty(k int) bool {
	return s.k[k] == -1
}

func (s *kStacks[T]) isFull() bool {
	return s.freetop == -1
}
