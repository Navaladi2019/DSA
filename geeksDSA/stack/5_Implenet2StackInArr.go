package stack

type TwoStacks[T any] struct {
	arr   []T
	size1 int
	size2 int
}

func (tt *TwoStacks[T]) Init(cap int) {

	tt.arr = make([]T, cap)
}

func (tt *TwoStacks[T]) Cap() int {

	return cap(tt.arr)
}

func (tt *TwoStacks[T]) Push1(val T) bool {

	if tt.size1+tt.size2 == len(tt.arr) {

		return false
	}

	tt.arr[tt.size1] = val
	tt.size1++
	return true
}

func (tt *TwoStacks[T]) Push2(val T) bool {

	if tt.size1+tt.size2 == len(tt.arr) {

		return false
	}

	tt.arr[len(tt.arr)-tt.size2-1] = val
	tt.size2++

	return true
}

func (tt *TwoStacks[T]) Pop1() (val T, ok bool) {

	if tt.size1 == 0 {

		return val, ok
	}

	val = tt.arr[tt.size1-1]
	tt.size1--
	return val, true
}

func (tt *TwoStacks[T]) Pop2() (val T, ok bool) {

	if tt.size2 == 0 {

		return val, ok
	}

	val = tt.arr[len(tt.arr)-tt.size2]
	tt.size2--
	return val, true
}

func (tt *TwoStacks[T]) Peek1() (val T, ok bool) {

	if tt.size1 == 0 {

		return val, ok
	}

	val = tt.arr[tt.size1-1]
	return val, true
}

func (tt *TwoStacks[T]) Peek2() (val T, ok bool) {

	if tt.size2 == 0 {

		return val, ok
	}

	val = tt.arr[len(tt.arr)-tt.size2]
	return val, true
}
