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

func IsRotated_1(str string, original string) bool {

	if len(str) != len(original) {
		return false
	}

	// first find the starting index of rotated array in original array

	// starting inedx of rotated array
	j := 0

	for i := 0; i < len(original); i++ {

		if original[i] == str[j] {
			j++
		} else if str[0] == original[i] {
			j = 1
		} else {
			j = 0
		}
	}

	// if j is equal to length index then
	// there is no need to travel again as the full strig is matched with original
	if j == len(str) {
		return true
	}

	// if j is zero then it did not find the first character in the original array itself
	if j == 0 {
		return false
	}

	for i := 0; i < len(original) && j < len(str); i++ {

		if str[j] == original[i] {
			j++
		} else {
			return false
		}
	}

	return true
}
