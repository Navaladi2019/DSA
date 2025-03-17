package recursion

import (
	"fmt"
)

func PrintAllPermutations(str string, prefix string, res *[]string) {

	if len(str) == 0 {

		*res = append(*res, prefix)
		return
	}

	for j := 0; j < len(str); j++ {

		tempprefix := prefix + string(str[j])

		tempstr := ""

		for i := 0; i < len(str); i++ {
			if i != j {
				tempstr = tempstr + string(str[i])
			}
		}

		if str == "ABC" && j == 2 {

			fmt.Printf("ok")
		}

		PrintAllPermutations(tempstr, tempprefix, res)
	}
}

func PrintAllPermutations_Efficient(str []string, i int, res *[]string) {

	if i == len(str) {

		strt := ""

		for i := 0; i < len(str); i++ {
			strt = strt + str[i]
		}
		*res = append(*res, strt)
		return
	}

	for j := i; j < len(str); j++ {
		str[i], str[j] = str[j], str[i]
		PrintAllPermutations_Efficient(str, i+1, res)
		str[i], str[j] = str[j], str[i]
	}
}

func Permutation(str []rune, i int) {

	if i == len(str) {
		fmt.Println(string(str))
		return
	}
	for j := i; j < len(str); j++ {
		str[i], str[j] = str[j], str[i]
		Permutation(str, i+1)
		str[i], str[j] = str[j], str[i]
	}
}
