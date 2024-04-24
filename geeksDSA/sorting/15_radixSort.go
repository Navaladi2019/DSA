package sorting

// radix sort uses counting sort in the hood

func RadixSort(arr []int) {

	maxEle := FindMaxElementInArray(arr)

	for exp := 1; maxEle/exp > 0; exp = exp * 10 {
		CountingSortExp(arr, exp)
	}

}

func CountingSortExp(arr []int, exp int) {

	countArr := make([]int, 10)
	Output := make([]int, len(arr))

	for i := 0; i < len(arr); i++ {

		digit := (arr[i] / exp) % 10

		countArr[digit]++
	}

	for i := 1; i < len(countArr); i++ {

		countArr[i] = countArr[i] + countArr[i-1]
	}

	for i := len(arr) - 1; i >= 0; i-- {

		couuntIndex := (arr[i] / exp) % 10
		index := countArr[couuntIndex] - 1

		Output[index] = arr[i]
		countArr[couuntIndex]--

	}

	for i := 0; i < len(Output); i++ {

		arr[i] = Output[i]
	}

}

func FindMaxElementInArray(arr []int) int {
	maxEle := arr[0]

	for i := 1; i < len(arr); i++ {
		maxEle = max(maxEle, arr[i])
	}

	return maxEle
}
