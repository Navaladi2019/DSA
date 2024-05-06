package arrays

func MaxAppearingElement(arr1 []int, arr2 []int) int {

	length := max(len(arr1), len(arr2))

	totArr := make([]int, 101)

	for i := 0; i < length; i++ {

		if len(arr1) > i {
			totArr[arr1[i]]++
		}

		if len(arr2) > i {
			totArr[arr2[i]]++
		}
	}

	maxappearing := 0
	for i := 1; i < 100; i++ {

		if totArr[i] > totArr[maxappearing] {
			maxappearing = i
		}
	}

	return maxappearing
}

func MaxApperaningElementINRanges(arr1 []int, arr2 []int) int {

	totArr := make([]int, 101)

	for i := 0; i < len(arr1); i++ {
		totArr[arr1[i]]++
		totArr[arr2[i]+1]--
	}

	maxindex := 0
	for i := 1; i < len(totArr); i++ {
		totArr[i] = totArr[i] + totArr[i-1]

		if totArr[i] > totArr[maxindex] {
			maxindex = i
		}
	}

	return maxindex
}
