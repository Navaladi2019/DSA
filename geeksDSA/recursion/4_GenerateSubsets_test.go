package recursion

import "testing"

func Test_Subsets(t *testing.T) {

	got := GenerateSubSets("ABCD", make([]string, 0), 0)

	GenerateSubsetsOptimized("ABC", "", 0)

	t.Log(got)
}
