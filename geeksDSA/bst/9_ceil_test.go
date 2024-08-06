package bst

import (
	"fmt"
	"testing"
)

func Test_Ceil_Arr(t *testing.T) {

	arr := []int{2, 8, 30, 15, 25, 12}
	Ceil_Arr(arr)

	fmt.Println("****")
	arr = []int{30, 20, 10}
	Ceil_Arr(arr)

	fmt.Println("****")
	arr = []int{10, 20, 30}
	Ceil_Arr(arr)

	fmt.Println("****")
	arr = []int{6, 18, 4, 5}
	Ceil_Arr(arr)

}
