package hashing

func LongestSubSequence(arr []int) int {

	dict := make(map[int]struct{})

	for i := 0; i < len(arr); i++ {
		dict[arr[i]] = struct{}{}
	}

	res := 1
	for i := 0; i < len(arr); i++ {
		curr := 1
		ele := arr[i]
		for {

			if _, ok := dict[ele+1]; ok {
				ele++
				curr++
			} else {
				break
			}
		}

		res = max(res, curr)

	}

	return res
}
