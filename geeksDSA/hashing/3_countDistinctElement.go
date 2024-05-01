package hashing

// here instead of dict we can use set but in golang we do not have set
func Count_DistinctElement(arr []int) int {

	dict := make(map[int]struct{})

	for i := 0; i < len(arr); i++ {
		dict[arr[i]] = struct{}{}
	}

	return len(dict)

}
