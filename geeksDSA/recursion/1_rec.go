package recursion

import (
	"fmt"
)

func test1(i int) {
	if i == 0 {
		return
	}
	fmt.Println(i)
	test1(i)
	fmt.Println(i)
}

func GetNto1(n int, resul *[]int) {

	if resul == nil {
		resul = new([]int)
	}
	if n == 0 {
		return
	}
	GetNto1(n-1, resul)
	*resul = append(*resul, n)
}

func finsSumofNaturalNumbers(n int) int {

	if n == 0 {
		return 0
	}

	return n + finsSumofNaturalNumbers(n-1)

}

func findisStringPalindromeUsingRecursion(str string, index int) bool {
	if index == len(str) || len(str)/2 == index {
		return true
	}

	if []rune(str)[index] != []rune(str)[len(str)-1-index] {
		return false
	}

	return findisStringPalindromeUsingRecursion(str, index+1)
}

func findSumOfDigoitsUsingRecursion(n int, sum int) int {

	lastvalue := n % 10

	if n == 10 {
		return sum + n
	}

	return findSumOfDigoitsUsingRecursion(n/10, lastvalue+sum)
}
