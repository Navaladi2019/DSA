package stack

import (
	"fmt"
	"testing"
)

func Test_ImplementTwoStack(t *testing.T) {

	ts := TwoStacks[int]{}

	ts.Init(10)

	ts.Push1(1)
	ts.Push1(1)
	ts.Push1(1)
	ts.Push1(1)
	ts.Push1(1)
	ts.Push1(1)
	ts.Push1(1)
	ts.Push1(1)
	ts.Push1(1)
	ts.Push1(1)

	fmt.Println(ts.Peek1())
	fmt.Println(ts.Peek2())
	fmt.Println(ts.Pop1())
	fmt.Println(ts.Pop2())

	fmt.Println(ts.Peek1())
	fmt.Println(ts.Peek2())
}
