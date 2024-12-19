package hashing

func InsersectionOfTwoArrays(arr1 []int, arr2 []int) []int {
	result := make([]int, 0, 10)

	dict := make(map[int]struct{})

	for i := 0; i < len(arr2); i++ {
		dict[arr2[i]] = struct{}{}
	}

	for i := 0; i < len(arr1); i++ {

		if _, ok := dict[arr1[i]]; ok {
			result = append(result, arr1[i])
		}
	}
	return result
}

func UnionOfTwoUnsortedArrays(arr1 []int, arr2 []int) int {

	m := make(map[int]struct{})

	for i := 0; i < len(arr1); i++ {

		if _, ok := m[arr1[i]]; !ok {
			m[arr1[i]] = struct{}{}
		}
	}

	for j := 0; j < len(arr2); j++ {

		if _, ok := m[arr2[j]]; !ok {
			m[arr1[j]] = struct{}{}
		}
	}

	return len(m)
}
