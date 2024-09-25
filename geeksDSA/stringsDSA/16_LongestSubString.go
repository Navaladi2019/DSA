package stringsDSA

func LongestDistinctSubstring_Naive(str string) int {

	res := 0

	txtArr := make([]int, 256)

	j := 0

	freshLoopStarts := 0
	for i := 0; i < len(str); i++ {
		txtArr[int(str[i])]++
		if txtArr[int(str[i])] == 1 {
			j++
			res = max(res, j)
		} else {

			for freshLoopStarts < i {

				txtArr[int(str[freshLoopStarts])]--
				if str[freshLoopStarts] == str[i] {
					j = i - freshLoopStarts
					freshLoopStarts++
					break
				}
				freshLoopStarts++
			}

		}
	}
	return res
}

// abaacdbab
func LongestSubStringEfficient(str string) int {

	res := 0

	txtArr := make([]int, 256)

	for i := 0; i < len(txtArr); i++ {
		txtArr[i] = -1
	}

	lastDuplicateIndex := -1
	for i := 0; i < len(str); i++ {

		// if txtArr[str[i]] > -1 && lastDuplicateIndex < txtArr[str[i]] {
		// 	lastDuplicateIndex = txtArr[str[i]]
		// }

		//above Condition Can Be Written AS Below if the value does not exist then its -1
		// or assign lastduplicate

		lastDuplicateIndex = max(lastDuplicateIndex, txtArr[str[i]])

		val := i + 1 - (lastDuplicateIndex + 1)
		res = max(res, val)
		txtArr[str[i]] = i
	}

	return res

}

func LongestSubstringPractice(str string) (res int) {

	txtArr := make([]int, 256)

	for i := 0; i < len(txtArr); i++ {
		txtArr[i] = -1
	}

	duplicateIndex := -1
	for i := 0; i < len(str); i++ {

		duplicateIndex = max(duplicateIndex, txtArr[str[i]])
		res = max(i-duplicateIndex, res)
		txtArr[str[i]] = i
	}

	return
}
