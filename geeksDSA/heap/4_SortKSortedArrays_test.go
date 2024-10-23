package heap

import (
	"fmt"
	"testing"
)

func Test_SortKSorted(t *testing.T) {

	arr := SortKSorted_Efficient([]int{9, 8, 7, 19, 18}, 2)
	fmt.Println(arr)

	arr = SortKSorted_Efficient([]int{10, 9, 7, 8, 4, 70, 50, 60}, 4)
	fmt.Println(arr)
}
