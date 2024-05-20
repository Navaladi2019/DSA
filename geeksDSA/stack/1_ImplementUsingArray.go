package stack

type ArrStack[T comparable] struct {
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

func (s *ArrStack[T]) Pop() (T, bool) {

	val, ok := s.extracLastdata()
	if ok {
		s.data = s.data[0 : len(s.data)-2]
	}

	return val, ok
}

func (s *ArrStack[T]) isEmpty() bool {
	return len(s.data) == 0
}
