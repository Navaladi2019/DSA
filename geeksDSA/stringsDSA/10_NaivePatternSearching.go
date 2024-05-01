package stringsDSA

func NativePatternSearching(str string, patt string) (bool, []int) {

	haspattern := false
	patternIndexdx := make([]int, 0, 10)

	lastwindow := str[0:len(patternIndexdx)]

	if lastwindow == patt {
		haspattern = true
		patternIndexdx = append(patternIndexdx, 0)
	}

	for i := 1; i < len(str); i++ {

		lastwindow = lastwindow[1:] + string(str[i])

		if lastwindow == patt {
			haspattern = true
			patternIndexdx = append(patternIndexdx, i)
		}

	}

	return haspattern, patternIndexdx
}

func NativePatternSearching_other(str string, patt string) (bool, []int) {

	haspattern := false
	patternIndexdx := make([]int, 0, 10)

	for i := 0; i < len(str)-len(patt); i++ {

		j := 0
		for j < len(patt) {

			if str[i] != patt[j] {
				break
			}
			j++
		}

		if j == len(patt) {
			haspattern = true
			patternIndexdx = append(patternIndexdx, i)
		}

	}

	return haspattern, patternIndexdx
}

func NativePatternSearchingFor_DistinctPattern(str string, pat string) (bool, []int) {
	haspattern := false
	patternIndexdx := make([]int, 0, 10)

	i := 0

	for i <= len(str)-len(pat) {

		j := 0
		for j < len(pat) {

			if str[i+j] != pat[j] {

				break
			}
			j++
		}
		if j == len(pat) {
			haspattern = true
			patternIndexdx = append(patternIndexdx, i)

		}

		if j == 0 {
			i++
		} else {
			i = i + j
		}

	}

	return haspattern, patternIndexdx
}
