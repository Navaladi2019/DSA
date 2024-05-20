package stack

import (
	"github.com/Navaladi2019/GoRefresher/linkedlists"
)

type LLStack[t comparable] struct {
	data linkedlists.SinglyLinkedList
}

func (s *LLStack[T]) Cap() int {
	return cap(s.data)
}

func (s *LLStack[T]) Peek() (T, bool) {
	return s.extracLastdata()
}

func (s *LLStack[T]) extracLastdata() (T, bool) {
	var zerodata T
	res := false
	if len(s.data) >= 1 {
		return s.data[len(s.data)-1], true
	}
	return zerodata, res
}

func (s *LLStack[T]) Pop() (T, bool) {

	val, ok := s.extracLastdata()
	if ok {
		s.data = s.data[0 : len(s.data)-2]
	}

	return val, ok
}

func (s *LLStack[T]) isEmpty() bool {
	return len(s.data) == 0
}
