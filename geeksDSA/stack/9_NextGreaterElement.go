package stack

import "fmt"

func NextGreaterElement(arr []int) {

	s := ArrStack[int]{}

	for i := len(arr) - 1; i >= 0; i-- {

		for !s.isEmpty() {
			val, _ := s.Peek()
			if arr[val] < arr[i] {
				s.Pop()
			} else {
				break
			}

		}

		if s.isEmpty() {
			fmt.Println(-1)
		} else {
			val, _ := s.Peek()
			fmt.Println(arr[val])
		}
		s.Push(i)
	}

}
