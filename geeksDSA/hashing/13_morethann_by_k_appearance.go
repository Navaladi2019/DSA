package hashing

func MoreThan_n_bu_k_appearance(arr []int, k int) []int {
	res := make([]int, 0, 0)

	dict := make(map[int]int)
	for i := 0; i < len(arr); i++ {
		if _, ok := dict[arr[i]]; ok {
			dict[arr[i]]++
		} else {
			dict[arr[i]] = 1
		}
	}

	for key, val := range dict {
		if val > len(arr)/k {
			res = append(res, key)
		}
	}

	return res
}
