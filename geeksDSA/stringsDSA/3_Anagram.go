package stringsDSA

func isAnagram_Dict(str1 string, str2 string) bool {

	if len(str1) != len(str2) {
		return false
	}
	r1 := []rune(str1)
	r2 := []rune(str2)

	dict := make(map[rune]int)

	for i := 0; i < len(r1); i++ {

		if _, ok := dict[r1[i]]; ok {
			dict[r1[i]]++
		} else {
			dict[r1[i]] = 1
		}
	}

	for j := 0; j < len(r2); j++ {

		if val, ok := dict[r2[j]]; ok {

			if val < 1 {
				return false
			}

			dict[r2[j]]--

			if dict[r2[j]] == 0 {
				delete(dict, r2[j])
			}
		} else {
			return false
		}
	}

	if len(dict) > 0 {
		return false
	}

	return true
}
