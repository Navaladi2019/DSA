package recursion

import "testing"

func Test_testGetNto1(t *testing.T) {

	got := new([]int)

	GetNto1(5, got)

	for index, value := range *got {

		if index+1 != value {
			t.Errorf("failed for value index %d => %d", index, value)
		}
	}
}

func Test_FindNaturalNumbers(t *testing.T) {

	got := finsSumofNaturalNumbers(5)

	if got != 15 {
		t.Errorf("ERROR")
	}
}

func Test_IsPalindrome(t *testing.T) {

	got := findisStringPalindromeUsingRecursion("abba", 0)

	if got != true {
		t.Errorf("is palindrome")
	}

	got = findisStringPalindromeUsingRecursion("a", 0)

	if got != true {
		t.Errorf("is palindrome")
	}

	got = findisStringPalindromeUsingRecursion("abaiaba", 0)

	if got != true {
		t.Errorf("is palindrome")
	}

	got = findisStringPalindromeUsingRecursion("geeks", 0)

	if got != false {
		t.Errorf("is palindrome")
	}
}

func Test_SumOfDigits(t *testing.T) {

	got := findSumOfDigoitsUsingRecursion(253, 0)

	if got != 10 {
		t.Errorf("not sum")
	}

	got = findSumOfDigoitsUsingRecursion(9987, 0)

	if got != 33 {
		t.Errorf("not sum")
	}

	got = findSumOfDigoitsUsingRecursion(10, 0)

	if got != 1 {
		t.Errorf("is palindrome")
	}

}
