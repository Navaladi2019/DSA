package hashing

func LongestSubArrayWithGivenSum(arr []int, sum int) int {

	length := 0
	dict := make(map[int]int)

	presum := 0

	for i := 0; i < len(arr); i++ {
		presum += arr[i]

		if presum == sum {
			length = i + 1
		}

		if valindex, ok := dict[presum-sum]; ok {
			indexx := i - valindex
			length = max(indexx, length)
		}

		if _, ok := dict[presum]; !ok {
			dict[presum] = i
		}

	}

	return length

}
