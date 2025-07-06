package arrays

func findMaxConsecutiveOne(arr []int) int {
	res := 0
	currentConsecutiveOne := 0

	for i := 0; i < len(arr); i++ {

		if arr[i] == 0 {
			currentConsecutiveOne = 0
		} else {

			currentConsecutiveOne++
			res = max(currentConsecutiveOne, res)
		}

	}

	return res
}

func FindMaxConsucitiveones(arr []int) int {

	res := 0

	curr := 0

	for _, v := range arr {
		if v == 1 {
			curr++
			res = max(res, curr)
		} else {
			curr = 0
		}
	}

	return res
}

func Test_valueSematics() {

}
