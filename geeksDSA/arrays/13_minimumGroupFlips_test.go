package arrays

import (
	"fmt"
	"testing"
)

func Test_MinimumConsecutiveFlips(t *testing.T) {

	arr := [][]int{
		{1, 1, 0, 0, 0, 1},
		{1, 0, 0, 0, 1, 0, 0, 1, 1, 1, 1},
		{0, 1},
	}

	for _, v := range arr {
		fmt.Println(MinimumConsecutiveFlips_Efficient(v))
	}

}
