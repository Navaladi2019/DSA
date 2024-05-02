package stringsDSA

func IsAnagramPatternPresent(str string, patt string) int {

	strBytes := []byte(str)

	_ = strBytes

	textArr := make([]int, 264)
	MatchArr := make([]int, 264)

	for i := 0; i < len(patt); i++ {
		textArr[str[i]]++
		MatchArr[patt[i]]++
	}

	for i := len(patt); i < len(str); i++ {

		if IsArraySame(textArr, MatchArr) {
			return i - len(patt)
		}

		textArr[str[i-len(patt)]]--
		textArr[str[i]]++
	}

	return -1
}

func IsArraySame(arr1 []int, arr2 []int) bool {

	for i := 0; i < len(arr1); i++ {

		if arr1[i] != arr2[i] {
			return false
		}

	}

	return true
}
