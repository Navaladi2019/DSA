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

// here instead of dict we can use char array of 256

func IsAnagram_SmallCase(str1 string, str2 string) bool {

	arr := make([]int, 26)

	for i := 0; i < len(str1); i++ {
		arr[str1[i]-'a']++
	}

	for i := 0; i < len(str2); i++ {
		arr[str2[i]-'a']--
	}

	for _, v := range arr {
		if v != 0 {
			return false
		}
	}

	return true
}
