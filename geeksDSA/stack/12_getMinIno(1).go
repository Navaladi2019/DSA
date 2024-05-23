package stack

type genMinStack struct {
	stack    ArrStack[int]
	minStack ArrStack[int]
}

func (s *genMinStack) Push(val int) {

	if minVal, ok := s.minStack.Peek(); ok {
		if minVal >= val {
			s.minStack.Push(val)
		}
	}
	s.stack.Push(val)
}

func (s *genMinStack) Pop() (val int, ok bool) {

	if val, ok = s.stack.Pop(); ok {
		if minVal, ok := s.minStack.Peek(); ok {
			if minVal == val {
				s.minStack.Pop()
			}
		}
	}

	return
}
