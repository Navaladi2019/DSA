package stringsDSA

import "strings"

func IsRotated(str string, original string) bool {

	if len(str) != len(original) {
		return false
	}

	j := 0
	for i := 0; i < len(str); i++ {
		if str[i] == original[j] {
			j++
		} else if str[i] == original[0] {
			j = 1
		} else {
			j = 0
		}
	}

	for i := 0; i <= len(original)-j; i++ {
		if str[i] == original[j] {
			j++
		} else {
			return false
		}
	}

	return true
}

func IsRotatedEasyApproach(str1 string, str2 string) bool {

	if len(str1) != len(str2) {
		return false
	}

	if strings.Index(str1+str1, str2) > -1 {
		return true
	}

	return false
}
