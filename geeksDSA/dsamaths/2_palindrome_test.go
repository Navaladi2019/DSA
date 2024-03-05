package dsamaths

import (
	"fmt"
	"testing"
)

type NumbersArrayPalindromeTestCase struct {
	numbers []int
	want    bool
}

type NumbersPalindromeTestCase struct {
	number int
	want   bool
}

var numberCases []NumbersPalindromeTestCase = []NumbersPalindromeTestCase{

	{878, true},
	{556, false},
	{9, true},
}

var numbersArrayPalindromeTestCase []NumbersArrayPalindromeTestCase = []NumbersArrayPalindromeTestCase{
	{[]int{1, 2, 3, 2, 1}, true},
	{[]int{1, 2, 3, 2, 0}, false},
	{[]int{1, 2, 1}, true},
	{[]int{2, 2}, true},
	{[]int{2, 0}, false},
	{[]int{2}, true},
}

func Test_PalidromeonNumbers(t *testing.T) {

	for _, c := range numberCases {
		t.Run(fmt.Sprintf("%d", c.number), func(t *testing.T) {
			got := istheNumberPalindrome(c.number)
			if got != c.want {
				t.Errorf("Expected %d as a palindrome but got the result as %t", c.number, c.want)
			}
		})
	}

}

func Test_PalindromeNumberArrays(t *testing.T) {

	for _, c := range numbersArrayPalindromeTestCase {
		t.Run(fmt.Sprintf("%d", c.numbers), func(t *testing.T) {
			got := isArrayNumbersPalindrome(c.numbers)
			if got != c.want {
				t.Errorf("Expected %t as a palindrome but got the result as %t", !c.want, c.want)
			}

		})
	}
}
