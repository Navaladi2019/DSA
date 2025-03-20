package bitmanipulation

import (
	"fmt"
	"testing"
)

func Test_SwapTwoNumbers(t *testing.T) {

	a, b := SwapTwoNumbers(5, 7)

	fmt.Println(a, b)
}

func Test_checkIfIthBitisSetorNot(t *testing.T) {

	fmt.Println(CheckIfIthBitIsSetOrNot(13, 1))
	fmt.Println(CheckIfIthBitIsSetOrNot(13, 2))
}
