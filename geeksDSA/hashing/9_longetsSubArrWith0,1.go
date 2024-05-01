package hashing

// Here we are replacing the zeros with -1 so we can use the longest subrray problem
func FindLargestSubArrayWithZerosAndOnes(arr []int) int {

	for i := 0; i < len(arr); i++ {
		if arr[i] == 0 {
			arr[i] = -1
		}
	}

	return LongestSubArrayWithGivenSum(arr, 0)
}
