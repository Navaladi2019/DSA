package sorting

func FindInterSectionOdTwoSortedArray(arr1 []int, arr2 []int) []int {

	result := make([]int, 0, 3)

	j := 0

	for i := 0; i < len(arr1); i++ {

		if i > 0 && arr1[i] == arr1[i-1] {
			continue
		}

		for j < len(arr2) {

			if arr1[i] == arr2[j] {
				result = append(result, arr1[i])
				break
			}

			if arr1[i] < arr2[j] {
				break
			}
			j++
		}
	}

	return result

}
