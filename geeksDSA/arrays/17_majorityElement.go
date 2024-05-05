package arrays

func FindMajorityElement(arr []int) int {

	count := 1
	lastEle := arr[0]

	for i := 1; i < len(arr); i++ {

		if arr[i] != lastEle {
			count--
		} else {
			count++
		}

		if count == 0 {
			lastEle = arr[i]
			count = 1
		}
	}

	finalCount := 0

	for i := 0; i < len(arr); i++ {
		if arr[i] == lastEle {
			finalCount++

			if finalCount > len(arr)/2 {
				return i
			}
		}
	}

	return -1
}

func FindMajorityElement_1(arr []int) int {

	majEle := make(map[int]int)

	for i := 0; i < len(arr); i++ {

		if _, ok := majEle[arr[i]]; ok {
			majEle[arr[i]]++

			if majEle[arr[i]] > len(arr)/2 {
				return i
			}
		} else {
			majEle[arr[i]] = 1
		}
	}

	return -1
}
