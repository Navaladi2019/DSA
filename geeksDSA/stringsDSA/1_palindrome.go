package stringsDSA

func IsPalindrome(str string) bool {

	i := 0
	j := len(str) - 1

	for i < j {
		if str[i] != str[j] {
			return false
		}
		i++
		j--
	}

	return true
}
