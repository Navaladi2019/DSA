package dsamaths

import (
	"strconv"
	"unicode/utf8"
)

func countDigits_ByConvertingToString(i int) int {
	str := strconv.Itoa(i)
	return utf8.RuneCountInString(str)
}

func countDigits_Divideby10(i int) int {
	count := 0

	for i > 0 {
		i = i / 10
		count++
	}

	return count
}
