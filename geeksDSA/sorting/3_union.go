package sorting

func finduionOfTwoArrays(arr1 []int, arr2 []int) []int {
	result := make([]int, 0, 20)

	lastpushed := min(arr1[0], arr2[0])

	i := 0
	j := 0

	if arr1[0] < arr2[0] {
		i = 1
	} else {
		j = 1
	}

	result = append(result, lastpushed)
	for i < len(arr1) && j < len(arr2) {

		if arr1[i] == result[len(result)-1] {
			i++
			continue
		}

		if arr2[j] == result[len(result)-1] {
			j++
			continue
		}

		if arr1[i] == arr2[j] {
			result = append(result, arr1[i])
			i++
			j++
		} else if arr1[i] < arr2[j] {
			result = append(result, arr1[i])
			i++
		} else if arr2[j] < arr1[i] {
			result = append(result, arr2[j])
			j++
		}

	}

	if i >= len(arr1) {
		for j < len(arr2) {
			if arr2[j] != result[len(result)-1] {
				result = append(result, arr2[j])
			}
			j++
		}

	}

	if j >= len(arr2) {
		for i < len(arr1) {
			if arr1[i] != result[len(result)-1] {
				result = append(result, arr1[i])

			}
			i++
		}

	}

	return result
}

func FinduionOfTwoArrays_1(arr1 []int, arr2 []int) (res []int) {

	i, j := 0, 0

	if arr1[0] < arr2[0] {
		res = append(res, arr1[0])
		i++
	} else {
		res = append(res, arr2[0])
		j++
	}

	for i < len(arr1) && j < len(arr2) {

		if arr1[i] == res[len(res)-1] {
			i++
			continue
		}
		if arr2[j] == res[len(res)-1] {
			j++
			continue
		}

		if arr1[i] == arr2[j] {
			res = append(res, arr1[i])
			i++
			j++
		} else if arr1[i] < arr2[j] {
			res = append(res, arr1[i])
			i++
		} else {
			res = append(res, arr2[j])
			j++
		}
	}

	for i < len(arr1) {
		if arr1[i] == res[len(res)-1] {
			i++
			continue
		}
	}

	for j < len(arr2) {
		if arr2[j] == res[len(res)-1] {
			j++
			continue
		}
	}

	return
}
