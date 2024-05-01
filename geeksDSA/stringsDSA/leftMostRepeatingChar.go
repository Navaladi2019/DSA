package stringsDSA

import "math"

func LeftMostRepeatingChar(str string) int {

	res := -1

	dict := make(map[rune]int)
	r := []rune(str)

	for i := 0; i < len(r); i++ {
		if _, ok := dict[r[i]]; ok {
			dict[r[i]]++
		} else {
			dict[r[i]] = 1
		}
	}

	for i := 0; i < len(r); i++ {
		val := dict[r[i]]
		if val > 1 {
			return i
		}
	}

	return res
}

// Here we consider string is ASCII chars only
func LeftMostRepeatingChar_Efficient(str string) int {

	fIndex := make([]int, 256)

	fill(fIndex, -1)

	res := math.MaxInt32

	for i := 0; i < len(str); i++ {

		if fIndex[str[i]] == -1 {
			fIndex[str[i]] = i
		} else {
			res = min(res, fIndex[str[i]])
		}
	}

	if res != math.MaxInt32 {
		return res
	}
	return -1
}

func fill[k any](arr []k, val k) {
	for i := 0; i < len(arr); i++ {
		arr[i] = val
	}
}
