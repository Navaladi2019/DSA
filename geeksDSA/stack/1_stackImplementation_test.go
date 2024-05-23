package stack

import (
	"fmt"
	"testing"
)

func Test_ArrStack(t *testing.T) {

	arrSta := ArrStack[int]{}

	arrSta.Push(1)
	arrSta.Push(2)

	fmt.Println(arrSta.data)
	AnotherMethod(&arrSta)

	fmt.Println(arrSta.data)
}

func AnotherMethod(s *ArrStack[int]) {

	s.Push(3)
}

func Test_LLStack(t *testing.T) {

	arrSta := LLStack[int]{}

	arrSta.Push(1)
	arrSta.Push(2)

	for e := arrSta.head; e != nil; e = e.next {
		fmt.Print(e.value)
		fmt.Print(" ,")
	}

	fmt.Println("***")
	AnotherMethod_LLStack(&arrSta)

	fmt.Println("***")

	for e := arrSta.head; e != nil; e = e.next {
		fmt.Print(e.value)
		fmt.Print(" ,")
	}
}

func AnotherMethod_LLStack(s *LLStack[int]) {

	s.Push(3)
	s.Push(4)

	for e := s.head; e != nil; e = e.next {
		fmt.Print(e.value)
		fmt.Print(" ,")
	}
}
