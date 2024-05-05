package recursion

import "testing"

func Test_Premutations(t *testing.T) {

	sli := make([]string, 0, 20)

	PrintAllPermutations_Efficient([]string{"a", "b", "c"}, 0, &sli)

	t.Log(sli)
}
