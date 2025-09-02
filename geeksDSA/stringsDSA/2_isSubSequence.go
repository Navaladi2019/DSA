package stringsDSA

func ISSubSequence(str string, str2 string) bool {

	r1 := []rune(str)
	r2 := []rune(str2)

	j := 0

	for i := 0; i < len(r1); i++ {

		if r1[i] == r2[j] && j == len(r2)-1 {

			return true
		}

		if r1[i] == r2[j] {
			j++
		}
	}

	return false
}

func ISSubSequenceRecursion(str string, str2 string, n int, m int) bool {
	if m == 0 {
		return true
	}
	if n == 0 {
		return false
	}

	if str[n-1] == str2[m-1] {
		return ISSubSequenceRecursion(str, str2, n-1, m-1)
	} else {
		return ISSubSequenceRecursion(str, str2, n-1, m)
	}

}

func ISSubSequence_1(str string, str2 string) bool {

	j := 0
	for i := 0; i < len(str) && j < len(str2); i++ {

		if str[i] == str2[j] {
			j++
		}

	}
	if j == len(str2) {
		return true
	}

	return false
}

func ISSubSequence_2(s1 string, s2 string) bool {

	j := 0

	for i := 0; i < len(s1); i++ {

		if s1[i] == s2[j] {
			j++
		}

		if j == len(s2) {
			return true
		}

	}

	return false
}
