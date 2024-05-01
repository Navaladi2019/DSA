package hashing

// here the concept is if we find the sum duplicate in dict then in somewhere the subarray sum becomes zero
func HasSubarraySumZero(arr []int) bool {

	dict := make(map[int]struct{})

	sum := 0

	for i := 0; i < len(arr); i++ {
		sum += arr[i]

		if sum == 0 || arr[i] == 0 {
			return true
		}
		if _, ok := dict[sum]; ok {
			return true
		} else {
			dict[sum] = struct{}{}
		}
	}

	return false
}

func HasSubarraySum(arr []int, sum int) bool {

	dict := make(map[int]struct{})

	presum := 0

	for i := 0; i < len(arr); i++ {
		presum += arr[i]

		if presum == sum {
			return true
		}
		if _, ok := dict[presum-sum]; ok {
			return true
		} else {
			dict[presum] = struct{}{}
		}
	}

	return false
}
