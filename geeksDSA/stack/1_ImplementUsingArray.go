package stack

// array implementation of stack is value type and not reference type
// so when you pass the value to other and push it it would not be available
type ArrStack[T any] struct {
	data []T
}

func (s *ArrStack[T]) Cap() int {
	return cap(s.data)
}

func (s *ArrStack[T]) Peek() (T, bool) {
	return s.extracLastdata()
}

func (s *ArrStack[T]) extracLastdata() (T, bool) {
	var zerodata T
	res := false
	if len(s.data) >= 1 {
		return s.data[len(s.data)-1], true
	}
	return zerodata, res
}

func (s *ArrStack[T]) Push(val T) {
	s.data = append(s.data, val)
}

func (s *ArrStack[T]) Pop() (T, bool) {

	val, ok := s.extracLastdata()
	if ok {
		s.data = s.data[:len(s.data)-1]
	}

	return val, ok
}

func (s *ArrStack[T]) IsEmpty() bool {
	return len(s.data) == 0
}
