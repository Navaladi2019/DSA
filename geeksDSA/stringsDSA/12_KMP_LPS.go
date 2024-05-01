package stringsDSA

import "slices"

// this is o^2 which is not desirable but there is a solution with o(n)
func FindLPS_Naive(str string) []int {

	runes := []rune(str)

	prefix := make([]string, 0, len(str)+1)

	suffix := make([]string, 0, len(str)+1)

	prefix = append(prefix, "")

	suffix = append(suffix, "")
	suffix = append(suffix, string(runes[0]))

	LPS := make([]int, 0, len(runes))
	LPS = append(LPS, 0)
	length := 0
	for i := 1; i < len(runes); i++ {

		length = 0
		prefix = append(prefix, prefix[len(prefix)-1]+string(runes[i-1]))

		for j := len(suffix) - 1; j > 0; j-- {

			suffix[j] = suffix[j] + string(runes[i])

		}
		suffix = slices.Concat(suffix[:1], []string{string(runes[i])}, suffix[1:])

		for j := 1; j < len(suffix) && j < len(prefix); j++ {
			if prefix[j] == suffix[j] {
				length = max(len(prefix[j]), length)
			}
		}

		LPS = append(LPS, length)

	}

	return LPS

}

func FindLPS(str string) []int {
	LPS := make([]int, len(str))

	LPS[0] = 0
	LastPS := 0
	i := 1

	for i < len(str) {

		if str[LastPS] == str[i] {
			LastPS++
			LPS[i] = LastPS
			i++

		} else if LastPS == 0 {
			LPS[i] = 0
			i++
		} else {
			LastPS--
		}

	}

	return LPS
}

func FindPatternKMP(str string, patt string) int {

	i := 0
	j := 0

	LPS := FindLPS(patt)
	for i < len(str) {

		if str[i] == patt[j] {
			i++
			j++
		} else if j > 0 {
			j = LPS[j-1]
		} else {
			i++
			j = 0
		}

		if j == len(patt) {
			return i - len(patt)
		}
	}

	return -1
}
