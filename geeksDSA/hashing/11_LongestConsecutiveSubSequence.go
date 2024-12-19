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
		if _, ok := (dict[ele-1]); ok {
			// here we are continuing because we have already considered the previous element as subsequence
			continue
		}
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

// here we did 2n look up
