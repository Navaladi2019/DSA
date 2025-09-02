package linkedlists

import "testing"

func Test_IsPalindrome(t *testing.T) {

	l := SinglyLinkedList{}
	l.Insert(1, 2, 3, 2, 1)

	got := IsPalindrome_Efficient(l)

	if got != true {
		t.Error("Has error in palindrome")
	}

	l = SinglyLinkedList{}
	l.Insert(1, 2, 1)

	got = IsPalindrome_Efficient(l)

	if got != true {
		t.Error("Has error in palindrome")
	}
	l = SinglyLinkedList{}
	l.Insert(1, 2, 2, 1)

	got = IsPalindrome_Efficient(l)

	if got != true {
		t.Error("Has error in palindrome")
	}

	l = SinglyLinkedList{}
	l.Insert(1, 2, 3, 4, 5)

	got = IsPalindrome_Efficient(l)

	if got != false {
		t.Error("Has error in palindrome")
	}
}
