package stack

type GetMinStatck struct {
	ArrStack[int]
	min int
}

// / here we are assuming that we will insert only positive values
func (s *GetMinStatck) Push(val int) {
	if s.IsEmpty() {

		s.min = val
	} else {
		if s.min > val {
			tempval := val
			val = val - s.min
			s.min = tempval
		}
	}
	s.ArrStack.Push(val)
}

func (s *GetMinStatck) Peek() (val int, ok bool) {

	val, ok = s.ArrStack.Peek()
	if val < 0 {
		val = s.min + (-val)
	}
	return
}

func (s *GetMinStatck) Pop() (val int, ok bool) {

	val, ok = s.ArrStack.Pop()
	if val < 0 {
		tempVal := val
		val = s.min
		s.min = s.min + (-tempVal)
	} else if s.min == val {
		s.min = 0
	}
	return
}
