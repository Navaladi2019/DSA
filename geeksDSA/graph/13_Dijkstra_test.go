package graph

import (
	"testing"
)

func Test_DijkstraAlgo(t *testing.T) {

	arr := make([][]int, 9)
	arr[0] = []int{0, 4, 8, 0, 0, 0, 0, 0, 0}   //A
	arr[1] = []int{4, 0, 11, 8, 0, 0, 0, 0, 0}  //B
	arr[2] = []int{8, 11, 0, 0, 7, 1, 0, 0, 0}  //C
	arr[3] = []int{0, 8, 0, 0, 2, 0, 7, 4, 0}   //D
	arr[4] = []int{0, 0, 7, 2, 0, 6, 0, 0, 0}   //E
	arr[5] = []int{0, 0, 1, 0, 6, 0, 0, 2, 0}   //F
	arr[6] = []int{0, 0, 0, 7, 0, 0, 0, 14, 9}  //G
	arr[7] = []int{0, 0, 0, 4, 0, 2, 14, 0, 10} //H
	arr[8] = []int{0, 0, 0, 0, 0, 0, 9, 10, 0}  //H
	DijkstraAlgo1(arr, 0)
}
