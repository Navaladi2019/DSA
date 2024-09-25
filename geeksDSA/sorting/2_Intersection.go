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

func FindInterSectionOdTwoSortedArray_1(arr1 []int, arr2 []int) (res []int) {

	i, j := 0, 0

	for i < len(arr1) && j < len(arr2) {

		if i > 0 && arr1[i] == arr1[i-1] {
			i++
			continue
		}

		if arr2[j] == arr1[i] {
			res = append(res, arr1[i])
			i++
			j++
		}

		if arr1[i] < arr2[j] {
			i++

		}

		if i < len(arr1) && arr1[i] > arr2[j] {
			j++

		}

	}

	return

}
