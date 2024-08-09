package heap

import (
	"fmt"
	"testing"
)

func Test_KthClosestElement(t *testing.T) {

	FindKClosestElemet([]int{10, 15, 7, 3, 4}, 2, 8)

	fmt.Println("****")
	FindKClosestElemet([]int{100, 80, 10, 5, 70}, 3, 2)
	fmt.Println("****")
	FindKClosestElemet([]int{20, 40, 5, 100, 150}, 3, 100)
}
