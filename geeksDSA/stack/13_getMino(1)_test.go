package stack

import (
	"fmt"
	"testing"
)

func Test_GetMino_1(t *testing.T) {

	s := GetMinStatck{}

	s.Push(10)
	s.Push(11)
	s.Push(9)
	fmt.Println(s.Peek())
	fmt.Println(s.min)
	s.Push(12)
	s.Push(7)
	s.Push(13)
	fmt.Println(s.Pop())
	fmt.Println(s.Pop())

	fmt.Println(s.Pop())
	fmt.Println(s.Pop())
	fmt.Println(s.Pop())

}
