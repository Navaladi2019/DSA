package hashing

func CountDistinctElementINEveryWindow(arr []int, window int) []int {

	res := make([]int, 0, 10)

	dict := make(map[int]int)
	for i := 0; i < window; i++ {
		if _, ok := dict[arr[i]]; ok {
			dict[arr[i]]++
		} else {
			dict[arr[i]] = 1
		}

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
