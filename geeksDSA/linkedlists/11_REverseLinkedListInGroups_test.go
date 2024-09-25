package linkedlists

import (
	"fmt"
	"testing"
)

/*
Reverses SingleyLinkedList in groups

[10,20,30,40,50,60] -> [30,20,10,60,50,40] for k =3

[10,20,30,40,50] -> [30,20,10,50,40] for k =3

[10,20,30] -> [30,20,10] for k =4
*/
func Test_reverseLinkedListInGroups(t *testing.T) {

	l := SinglyLinkedList{}

	l.Insert(10, 20, 30, 40, 50, 60, 70, 80, 90, 91, 92)

	RecursiveLinkedList(l.head)

	fmt.Println("****")

	l = ReverseLinkedListInGroups_1(l, 3)

	RecursiveLinkedList(l.head)

}
