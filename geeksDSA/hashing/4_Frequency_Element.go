package hashing

func Frequency_Element(arr []int) map[int]int {

	dict := make(map[int]int)

	for i := 0; i < len(arr); i++ {
		dict[arr[i]] = dict[arr[i]] + 1
	}

	return dict

}
