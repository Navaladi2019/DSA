package recursion

import (
	"fmt"
	"unicode/utf8"
)

func GenerateSubSets(str string, subsets []string, currentIndex int) []string {

	if currentIndex == utf8.RuneCountInString(str) {
		return subsets
	}

	newSubsets := make([]string, 0, len(subsets)+1)
	newSubsets = append(newSubsets, string([]rune(str)[currentIndex]))

	for _, e := range subsets {

		newSubsets = append(newSubsets, e+string([]rune(str)[currentIndex]))
	}

	subsets = append(subsets, newSubsets...)

	return GenerateSubSets(str, subsets, currentIndex+1)

}

func GenerateSubsetsOptimized(str string, curr string, currentIndex int) {

	if currentIndex == utf8.RuneCountInString(str) {

		fmt.Print(curr)
		return
	}

	GenerateSubsetsOptimized(str, curr, currentIndex+1) // ->

	GenerateSubsetsOptimized(str, curr+string([]rune(str)[currentIndex]), currentIndex+1)
}

func GenerateSubSetsRevision(str string, curr string, n int) {

	if n == utf8.RuneCountInString(str) {

		fmt.Print(curr)
		return
	}

	GenerateSubSetsRevision(str, curr+string([]rune(str)[n]), n+1)
	GenerateSubSetsRevision(str, curr, n+1)

}
