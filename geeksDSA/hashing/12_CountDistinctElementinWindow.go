package hashing

func CountDistinctElementINEveryWindow(arr []int, window int) []int {

	res := make([]int, 0, 10)

	dict := make(map[int]int)
	for i := 0; i < window; i++ {
		//if _, ok := dict[arr[i]]; ok {
		dict[arr[i]]++
		// } else {
		// 	dict[arr[i]] = 1
		// }

	}

	res = append(res, len(dict))
	for i := 1; i <= len(arr)-window; i++ {

		dict[arr[i-1]]--
		if val := dict[arr[i-1]]; val == 0 {
			delete(dict, arr[i-1])
		}

		if _, ok := dict[arr[i+window-1]]; ok {
			dict[arr[i+window-1]]++
		} else {
			dict[arr[i+window-1]] = 1
		}
		res = append(res, len(dict))
	}

	return res
}

func CountDistinctElementInEveryWindowNaive(arr []int, window int) []int {

	res := make([]int, 0, 10)

	for i := 0; i <= len(arr)-window; i++ {

		count := 0
		for j := 0; j < window; j++ {

			isExists := false

			for k := j + 1; k < window; k++ {

				if arr[k+i] == arr[j+i] {
					isExists = true
					break
				}
			}

			if isExists == false {
				count++
			}

		}

		res = append(res, count)
	}

	return res
}
