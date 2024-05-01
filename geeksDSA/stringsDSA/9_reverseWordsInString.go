package stringsDSA

import "strings"

func ReverseWords_Naive(str string) string {

	strarr := strings.Split(str, " ")

	builder := strings.Builder{}
	for i := len(strarr) - 1; i >= 0; i-- {
		builder.WriteString(strarr[i])
	}

	return builder.String()
}

func ReverseWords_Efficient(str string) string {

	r := []rune(str)

	start := 0

	for end := 0; end < len(r); end++ {
		if r[end] == ' ' {
			reverseRune(r, start, end-1)
			start = end + 1
		}
	}

	reverseRune(r, start, len(r)-1)
	reverseRune(r, 0, len(r)-1)

	return string(r)

}

func reverseRune(r []rune, start int, end int) {

	for start != end {
		r[start], r[end] = r[end], r[start]
	}

}
