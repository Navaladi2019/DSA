package backtracking

import (
	"fmt"
	"strings"
)

func PrintALlPermutationDoesNotContainAB_Naive(str []string, i int) {

	if i == len(str) {

		if strings.Contains(strings.Join(str, ""), "AB") {
			fmt.Println(str)
		}

	}
	for j := i; j < len(str); j++ {
		str[i], str[j] = str[j], str[i]
		PrintALlPermutationDoesNotContainAB_Naive(str, j+1)
		str[i], str[j] = str[j], str[i]

	}
}

func PrintALlPermutationDoesNotContainAB_Efficient(str []string, i int) {

	if i == len(str) {
		fmt.Println(str)
	}
	for j := i; j < len(str); j++ {

		if IsSafeAB(str, i, j) {
			str[i], str[j] = str[j], str[i]
			PrintALlPermutationDoesNotContainAB_Efficient(str, i+1)
			str[i], str[j] = str[j], str[i]
		}
	}
}

func IsSafeAB(str []string, i int, j int) bool {

	PrevCharIndex := i - 1

	if PrevCharIndex >= 0 && str[PrevCharIndex] == "A" && str[j] == "B" {
		return false
	}

	return true
}
