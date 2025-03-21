package recursion

import (
	"fmt"
	"testing"
)

func Test_Premutations(t *testing.T) {

	Permutation([]rune{'A', 'B', 'C'}, 0)

	sli := make([]string, 0, 20)

	PrintAllPermutations_Efficient([]string{"a", "b", "c"}, 0, &sli)

	fmt.Println(sli)
}
