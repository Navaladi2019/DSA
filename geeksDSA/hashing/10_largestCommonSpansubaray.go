package hashing

func LongestCommonSpanWithSameSumBinaryArrr(arr1 []int, arr2 []int) int {

	arr := make([]int, 0, len(arr1))

	for i := 0; i < len(arr1); i++ {
		arr = append(arr, arr1[i]-arr2[i])
	}

	return LongestSubArrayWithGivenSum(arr, 0)
}
