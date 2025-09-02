package hashing

/* here the concept is if we find the prefix sum duplicate in dict then
in somewhere the subarray sum becomes zero
*/
func HasSubarraySumZero(arr []int) bool {

	dict := make(map[int]struct{})

	presum := 0

	for i := 0; i < len(arr); i++ {
		presum += arr[i]

		if presum == 0 || arr[i] == 0 {
			return true
		}
		if _, ok := dict[presum]; ok {
			return true
		} else {
			dict[presum] = struct{}{}
		}
	}

	return false
}

// [5,8,6,13,3,-1]
// [5,13,19,32,35,34] stored in dictionary but within the loop
// 0 -> 0-16 => -16 does not exist i++
// 1 -> 5-16 => -11 does not exist
// 2 -> 13-16 => -3 does not exist
// 3 -> 19 -16 => 3 does not exist
// 4 -> 32-16  => 16 does not exist
// 5 -> 35-16 => 19 [it exist] so the target sum exist
func HasSubarraySum(arr []int, targetSum int) bool {

	dict := make(map[int]struct{})

	presum := 0

	for i := 0; i < len(arr); i++ {
		presum += arr[i]

		if presum == targetSum {
			return true
		}
		if _, ok := dict[presum-targetSum]; ok {
			return true
		} else {
			dict[presum] = struct{}{}
		}
	}

	return false
}
